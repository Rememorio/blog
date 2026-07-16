package main

import (
	"runtime"
	"runtime/metrics"
	"sync"
	"testing"
)

func TestSchedulerBurst(t *testing.T) {
	const workers = 64

	ready := make(chan struct{}, workers)
	release := make(chan struct{})
	var finished sync.WaitGroup
	finished.Add(workers)

	for range workers {
		go func() {
			defer finished.Done()
			ready <- struct{}{}
			<-release
			runtime.Gosched()
		}()
	}

	for range workers {
		<-ready
	}
	close(release)
	finished.Wait()
}

func TestSchedulerMetricsAreAvailable(t *testing.T) {
	samples := []metrics.Sample{
		{Name: "/sched/gomaxprocs:threads"},
		{Name: "/sched/goroutines:goroutines"},
		{Name: "/sched/goroutines/runnable:goroutines"},
		{Name: "/sched/latencies:seconds"},
	}
	metrics.Read(samples)

	for _, sample := range samples[:3] {
		if sample.Value.Kind() != metrics.KindUint64 {
			t.Fatalf("metric %s has kind %v, want uint64", sample.Name, sample.Value.Kind())
		}
	}
	histogram := samples[3].Value.Float64Histogram()
	if len(histogram.Buckets) == 0 || len(histogram.Counts) == 0 {
		t.Fatal("scheduler latency histogram has no buckets")
	}
}

func BenchmarkGoroutineHandoff(b *testing.B) {
	ping := make(chan struct{})
	pong := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ping:
				pong <- struct{}{}
			case <-done:
				return
			}
		}
	}()

	b.ResetTimer()
	for range b.N {
		ping <- struct{}{}
		<-pong
	}
	b.StopTimer()
	close(done)
}
