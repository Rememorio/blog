//go:build go1.27

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"testing"
	"time"
)

// This lab checks the implementation's bounded drain, not an API guarantee
// that closing any response body makes its connection reusable.
func TestEarlySmallBodyCloseCanReuseAfterDrain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "small response")
	}))
	t.Cleanup(server.Close)
	client := newHTTP1Client(t, 1)
	idle := make(chan error, 1)
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), &httptrace.ClientTrace{
		PutIdleConn: func(err error) { idle <- err },
	}))
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if err := resp.Body.Close(); err != nil {
		t.Fatal(err)
	}
	// Close may return before draining finishes. Observe the pool event
	// before issuing the request whose reuse is being measured.
	select {
	case err := <-idle:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("small response was not drained into the idle pool")
	}
	next := tracedGet(t, client, server.URL, true)
	if next.err != nil {
		t.Fatal(next.err)
	}
	if !next.reused {
		t.Fatal("request did not reuse the connection after the drain completed")
	}
}
