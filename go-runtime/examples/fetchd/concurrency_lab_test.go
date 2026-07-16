package main

import (
	"context"
	"sync"
	"testing"
	"testing/synctest"
)

func TestBufferedChannelAppliesBackpressure(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		results := make(chan int, 1)
		secondSendCompleted := false

		go func() {
			results <- 1
			results <- 2
			secondSendCompleted = true
		}()

		synctest.Wait()
		if secondSendCompleted {
			t.Fatal("second send completed while the one-slot buffer was full")
		}
		if got := len(results); got != 1 {
			t.Fatalf("len(results) = %d, want 1", got)
		}

		if got := <-results; got != 1 {
			t.Fatalf("first receive = %d, want 1", got)
		}
		synctest.Wait()
		if !secondSendCompleted {
			t.Fatal("second send remained blocked after the receiver freed a slot")
		}
		if got := <-results; got != 2 {
			t.Fatalf("second receive = %d, want 2", got)
		}
	})
}

func TestSelectCanExitOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan fetchResult)
	cancel()

	select {
	case <-results:
		t.Fatal("received a result that was never sent")
	case <-ctx.Done():
		if ctx.Err() != context.Canceled {
			t.Fatalf("ctx.Err() = %v, want context.Canceled", ctx.Err())
		}
	}
}

func TestMutexProtectsSharedSummary(t *testing.T) {
	const (
		workers    = 32
		increments = 100
	)

	counts := map[string]int{"completed": 0}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for range increments {
				mu.Lock()
				counts["completed"]++
				mu.Unlock()
			}
		})
	}
	wg.Wait()

	if got, want := counts["completed"], workers*increments; got != want {
		t.Fatalf("completed = %d, want %d", got, want)
	}
}

func BenchmarkCoordinationContracts(b *testing.B) {
	b.Run("channel-handoff", func(b *testing.B) {
		values := make(chan int)
		done := make(chan struct{})
		go func() {
			for range values {
			}
			close(done)
		}()

		b.ResetTimer()
		for i := range b.N {
			values <- i
		}
		b.StopTimer()
		close(values)
		<-done
	})

	b.Run("mutex-counter", func(b *testing.B) {
		var mu sync.Mutex
		counter := 0
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		})
		if counter != b.N {
			b.Fatalf("counter = %d, want %d", counter, b.N)
		}
	})
}
