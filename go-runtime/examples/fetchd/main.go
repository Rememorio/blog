package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultRequestTimeout = 1500 * time.Millisecond
	maxResponseBytes      = 1 << 20
)

type fetchResult struct {
	URL      string `json:"url"`
	Status   int    `json:"status,omitempty"`
	Bytes    int64  `json:"bytes,omitempty"`
	Duration string `json:"duration"`
	Error    string `json:"error,omitempty"`
}

type fetchHandler struct {
	client     *http.Client
	timeout    time.Duration
	maxTargets int
}

func (h fetchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "GET requests only", http.StatusMethodNotAllowed)
		return
	}

	targets := r.URL.Query()["url"]
	if len(targets) == 0 || len(targets) > h.maxTargets {
		http.Error(w, "provide between 1 and 8 url parameters", http.StatusBadRequest)
		return
	}
	for _, target := range targets {
		parsed, err := url.ParseRequestURI(target)
		if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
			http.Error(w, "every url must be an absolute HTTP(S) URL", http.StatusBadRequest)
			return
		}
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	results := make(chan fetchResult, len(targets))
	for _, target := range targets {
		go h.fetch(ctx, target, results)
	}

	batch := make([]fetchResult, 0, len(targets))
	for range targets {
		batch = append(batch, <-results)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(batch); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func (h fetchHandler) fetch(ctx context.Context, target string, results chan<- fetchResult) {
	started := time.Now()
	result := fetchResult{URL: target}
	defer func() {
		result.Duration = time.Since(started).Round(time.Millisecond).String()
		results <- result
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		result.Error = err.Error()
		return
	}

	resp, err := h.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			result.Error = "request deadline exceeded"
		} else {
			result.Error = err.Error()
		}
		return
	}
	defer resp.Body.Close()

	result.Status = resp.StatusCode
	result.Bytes, err = io.Copy(io.Discard, io.LimitReader(resp.Body, maxResponseBytes))
	if err != nil {
		result.Error = err.Error()
	}
}

func main() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 16
	transport.MaxIdleConnsPerHost = 8
	transport.IdleConnTimeout = 90 * time.Second
	transport.ResponseHeaderTimeout = time.Second

	handler := fetchHandler{
		client:     &http.Client{Transport: transport},
		timeout:    defaultRequestTimeout,
		maxTargets: 8,
	}
	server := &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("fetchd listening on %s", server.Addr)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
