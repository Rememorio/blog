package main

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func TestCancelCauseKeepsTheFirstCauseOnEachBranch(t *testing.T) {
	parentCause := errors.New("client disconnected")
	childCause := errors.New("worker rejected result")

	t.Run("parent wins when it cancels first", func(t *testing.T) {
		parent, cancelParent := context.WithCancelCause(context.Background())
		child, cancelChild := context.WithCancelCause(parent)

		cancelParent(parentCause)
		cancelChild(childCause)

		if got := context.Cause(parent); !errors.Is(got, parentCause) {
			t.Fatalf("parent cause = %v, want %v", got, parentCause)
		}
		if got := context.Cause(child); !errors.Is(got, parentCause) {
			t.Fatalf("child cause = %v, want inherited %v", got, parentCause)
		}
	})

	t.Run("child keeps its earlier local cause", func(t *testing.T) {
		parent, cancelParent := context.WithCancelCause(context.Background())
		child, cancelChild := context.WithCancelCause(parent)

		cancelChild(childCause)
		cancelParent(parentCause)

		if got := context.Cause(parent); !errors.Is(got, parentCause) {
			t.Fatalf("parent cause = %v, want %v", got, parentCause)
		}
		if got := context.Cause(child); !errors.Is(got, childCause) {
			t.Fatalf("child cause = %v, want local %v", got, childCause)
		}
	})
}

func TestCancelSignalsButDoesNotJoinWorker(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cleanupRelease := make(chan struct{})
		workerDone := make(chan struct{})

		go func() {
			defer close(workerDone)
			<-ctx.Done()
			<-cleanupRelease
		}()

		cancel()
		synctest.Wait()
		select {
		case <-workerDone:
			t.Fatal("cancel unexpectedly waited for worker cleanup")
		default:
		}

		close(cleanupRelease)
		synctest.Wait()
		select {
		case <-workerDone:
		default:
			t.Fatal("worker did not finish after cleanup was released")
		}
	})
}

func TestTimeoutCauseAndWithoutCancelHaveDifferentLifetimes(t *testing.T) {
	timeoutCause := errors.New("fetch budget exhausted")

	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeoutCause(context.Background(), time.Second, timeoutCause)
		defer cancel()

		<-ctx.Done()
		if !errors.Is(ctx.Err(), context.DeadlineExceeded) {
			t.Fatalf("Err = %v, want DeadlineExceeded", ctx.Err())
		}
		if got := context.Cause(ctx); !errors.Is(got, timeoutCause) {
			t.Fatalf("Cause = %v, want %v", got, timeoutCause)
		}
	})

	type requestIDKey struct{}
	parent, cancelParent := context.WithCancel(context.WithValue(context.Background(), requestIDKey{}, "req-42"))
	detached := context.WithoutCancel(parent)
	cancelParent()

	if detached.Done() != nil || detached.Err() != nil || context.Cause(detached) != nil {
		t.Fatal("WithoutCancel unexpectedly retained cancellation or deadline state")
	}
	if got := detached.Value(requestIDKey{}); got != "req-42" {
		t.Fatalf("request ID = %v, want req-42", got)
	}
}

func TestAfterFuncStopDoesNotWaitForAStartedCallback(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		callbackStarted := make(chan struct{})
		callbackRelease := make(chan struct{})
		callbackDone := make(chan struct{})

		stop := context.AfterFunc(ctx, func() {
			close(callbackStarted)
			<-callbackRelease
			close(callbackDone)
		})
		cancel()
		<-callbackStarted

		if stop() {
			t.Fatal("stop reported success after the callback started")
		}
		select {
		case <-callbackDone:
			t.Fatal("stop unexpectedly waited for callback completion")
		default:
		}

		close(callbackRelease)
		synctest.Wait()
		select {
		case <-callbackDone:
		default:
			t.Fatal("callback did not finish after release")
		}
	})
}

type closeSignalListener struct {
	net.Listener
	closed chan struct{}
	once   sync.Once
}

func (l *closeSignalListener) Close() error {
	l.once.Do(func() { close(l.closed) })
	return l.Listener.Close()
}

func TestServerShutdownWaitsForAnActiveHandler(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	signaledListener := &closeSignalListener{Listener: listener, closed: make(chan struct{})}

	handlerStarted := make(chan struct{})
	handlerRelease := make(chan struct{})
	server := &http.Server{Handler: http.HandlerFunc(func(response http.ResponseWriter, _ *http.Request) {
		close(handlerStarted)
		<-handlerRelease
		_, _ = io.WriteString(response, "done")
	})}

	serveDone := make(chan error, 1)
	go func() { serveDone <- server.Serve(signaledListener) }()

	requestDone := make(chan error, 1)
	go func() {
		response, requestErr := http.Get("http://" + listener.Addr().String())
		if requestErr == nil {
			_, requestErr = io.Copy(io.Discard, response.Body)
			response.Body.Close()
		}
		requestDone <- requestErr
	}()
	<-handlerStarted

	shutdownDone := make(chan error, 1)
	go func() { shutdownDone <- server.Shutdown(context.Background()) }()
	<-signaledListener.closed

	select {
	case shutdownErr := <-shutdownDone:
		t.Fatalf("Shutdown returned while the handler was active: %v", shutdownErr)
	default:
	}

	close(handlerRelease)
	select {
	case shutdownErr := <-shutdownDone:
		if shutdownErr != nil {
			t.Fatalf("Shutdown = %v, want nil", shutdownErr)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Shutdown did not finish after the handler returned")
	}

	if requestErr := <-requestDone; requestErr != nil {
		t.Fatalf("request = %v, want nil", requestErr)
	}
	if serveErr := <-serveDone; !errors.Is(serveErr, http.ErrServerClosed) {
		t.Fatalf("Serve = %v, want ErrServerClosed", serveErr)
	}
}
