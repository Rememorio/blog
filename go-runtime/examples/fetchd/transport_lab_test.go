package main

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type tracedResponse struct {
	conn   net.Conn
	reused bool
	proto  int
	err    error
}

func newHTTP1Client(t *testing.T, maxConns int) *http.Client {
	t.Helper()
	transport := &http.Transport{
		MaxConnsPerHost:     maxConns,
		MaxIdleConns:        maxConns,
		MaxIdleConnsPerHost: maxConns,
		IdleConnTimeout:     time.Minute,
	}
	t.Cleanup(transport.CloseIdleConnections)
	return &http.Client{Transport: transport}
}

func tracedGet(t *testing.T, client *http.Client, target string, readBody bool) tracedResponse {
	t.Helper()
	var result tracedResponse
	putIdle := make(chan error, 1)
	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			result.conn = info.Conn
			result.reused = info.Reused
		},
		PutIdleConn: func(err error) {
			putIdle <- err
		},
	}
	req, err := http.NewRequestWithContext(
		httptrace.WithClientTrace(context.Background(), trace),
		http.MethodGet,
		target,
		nil,
	)
	if err != nil {
		result.err = err
		return result
	}
	resp, err := client.Do(req)
	if err != nil {
		result.err = err
		return result
	}
	result.proto = resp.ProtoMajor
	if readBody {
		_, err = io.Copy(io.Discard, resp.Body)
	}
	closeErr := resp.Body.Close()
	if err == nil {
		err = closeErr
	}
	result.err = err
	if readBody && err == nil && result.proto == 1 {
		select {
		case putErr := <-putIdle:
			result.err = putErr
		case <-time.After(2 * time.Second):
			result.err = errors.New("connection was not returned to the idle pool")
		}
	}
	return result
}

func TestTransportReusesConnectionAfterBodyEOF(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "fetchd")
	}))
	t.Cleanup(server.Close)

	client := newHTTP1Client(t, 1)
	first := tracedGet(t, client, server.URL, true)
	second := tracedGet(t, client, server.URL, true)

	if first.err != nil || second.err != nil {
		t.Fatalf("requests failed: first=%v second=%v", first.err, second.err)
	}
	if first.proto != 1 || second.proto != 1 {
		t.Fatalf("protocols = HTTP/%d and HTTP/%d, want HTTP/1", first.proto, second.proto)
	}
	if first.reused {
		t.Fatal("first request unexpectedly reused a connection")
	}
	if !second.reused {
		t.Fatal("second request did not reuse the idle HTTP/1 connection")
	}
	if first.conn != second.conn {
		t.Fatal("GotConn reported different connections after reading the first body to EOF")
	}
}

func TestEarlyBodyClosePreventsHTTP1Reuse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "response body must reach EOF")
	}))
	t.Cleanup(server.Close)

	client := newHTTP1Client(t, 1)
	first := tracedGet(t, client, server.URL, false)
	second := tracedGet(t, client, server.URL, true)

	if first.err != nil || second.err != nil {
		t.Fatalf("requests failed: first=%v second=%v", first.err, second.err)
	}
	if second.reused {
		t.Fatal("HTTP/1 connection was reused after the first body closed before EOF")
	}
	if first.conn == second.conn {
		t.Fatal("GotConn reported the same connection after an early body close")
	}
}

func TestMaxConnsPerHostQueueObservesCancellation(t *testing.T) {
	var handled atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		handled.Add(1)
		_, _ = io.WriteString(w, "held response")
	}))
	t.Cleanup(server.Close)

	client := newHTTP1Client(t, 1)
	firstResp, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer firstResp.Body.Close()

	ctx, cancel := context.WithCancel(context.Background())
	getConn := make(chan struct{})
	var getConnOnce sync.Once
	trace := &httptrace.ClientTrace{
		GetConn: func(string) {
			getConnOnce.Do(func() { close(getConn) })
		},
	}
	req, err := http.NewRequestWithContext(
		httptrace.WithClientTrace(ctx, trace),
		http.MethodGet,
		server.URL,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	result := make(chan error, 1)
	go func() {
		resp, requestErr := client.Do(req)
		if resp != nil {
			resp.Body.Close()
		}
		result <- requestErr
	}()

	select {
	case <-getConn:
	case <-time.After(2 * time.Second):
		t.Fatal("second request never entered Transport.getConn")
	}
	cancel()

	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("queued request error = %v, want context canceled", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("queued request did not observe cancellation")
	}
	if got := handled.Load(); got != 1 {
		t.Fatalf("server handled %d requests while the only HTTP/1 connection was busy, want 1", got)
	}
}

func TestHTTP2MultiplexesConcurrentRequestsOnOneConnection(t *testing.T) {
	arrived := make(chan string, 2)
	release := make(chan struct{})
	var newConnections atomic.Int32

	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/warm" {
			_, _ = io.WriteString(w, "warm")
			return
		}
		arrived <- r.URL.Path
		<-release
		_, _ = io.WriteString(w, r.URL.Path)
	}))
	server.EnableHTTP2 = true
	server.Config.ConnState = func(_ net.Conn, state http.ConnState) {
		if state == http.StateNew {
			newConnections.Add(1)
		}
	}
	server.StartTLS()
	t.Cleanup(server.Close)

	client := server.Client()
	t.Cleanup(client.CloseIdleConnections)
	warm := tracedGet(t, client, server.URL+"/warm", true)
	if warm.err != nil {
		t.Fatal(warm.err)
	}
	if warm.proto != 2 {
		t.Fatalf("warm request used HTTP/%d, want HTTP/2", warm.proto)
	}

	results := make(chan tracedResponse, 2)
	for _, path := range []string{"/one", "/two"} {
		path := path
		go func() {
			results <- tracedGet(t, client, server.URL+path, true)
		}()
	}

	for range 2 {
		select {
		case <-arrived:
		case <-time.After(2 * time.Second):
			close(release)
			t.Fatal("two HTTP/2 streams did not reach the server concurrently")
		}
	}
	close(release)

	first := <-results
	second := <-results
	if first.err != nil || second.err != nil {
		t.Fatalf("HTTP/2 requests failed: first=%v second=%v", first.err, second.err)
	}
	if first.proto != 2 || second.proto != 2 {
		t.Fatalf("protocols = HTTP/%d and HTTP/%d, want HTTP/2", first.proto, second.proto)
	}
	if first.conn != warm.conn || second.conn != warm.conn {
		t.Fatal("concurrent HTTP/2 streams did not share the warmed connection")
	}
	if got := newConnections.Load(); got != 1 {
		t.Fatalf("TLS connections = %d, want one multiplexed HTTP/2 connection", got)
	}
}
