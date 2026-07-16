package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestFetchHandlerCollectsResponses(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("hello"))
	}))
	t.Cleanup(upstream.Close)

	handler := fetchHandler{
		client:     upstream.Client(),
		timeout:    time.Second,
		maxTargets: 8,
	}
	recorder := httptest.NewRecorder()
	target := "/fetch?url=" + url.QueryEscape(upstream.URL) + "&url=" + url.QueryEscape(upstream.URL)
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, target, nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	var results []fetchResult
	if err := json.Unmarshal(recorder.Body.Bytes(), &results); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("results = %d, want 2", len(results))
	}
	for _, result := range results {
		if result.Status != http.StatusCreated || result.Bytes != 5 || result.Error != "" {
			t.Fatalf("unexpected result: %#v", result)
		}
	}
}

func TestFetchHandlerPropagatesDeadline(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(250 * time.Millisecond):
			w.WriteHeader(http.StatusNoContent)
		case <-r.Context().Done():
		}
	}))
	t.Cleanup(upstream.Close)

	handler := fetchHandler{
		client:     upstream.Client(),
		timeout:    25 * time.Millisecond,
		maxTargets: 8,
	}
	recorder := httptest.NewRecorder()
	target := "/fetch?url=" + url.QueryEscape(upstream.URL)
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, target, nil))

	if !strings.Contains(recorder.Body.String(), "request deadline exceeded") {
		t.Fatalf("response = %q, want deadline error", recorder.Body.String())
	}
}

func TestFetchHandlerRejectsInvalidTarget(t *testing.T) {
	handler := fetchHandler{
		client:     http.DefaultClient,
		timeout:    time.Second,
		maxTargets: 8,
	}
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/fetch?url=mailto:test@example.com", nil))

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
}
