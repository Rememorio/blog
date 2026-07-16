package main

import (
	"runtime/metrics"
	"testing"
)

var (
	memoryIntSink     int64
	memoryPointerSink *fetchResult
	memorySliceSink   []fetchResult
)

//go:noinline
func stackResultChecksum() int64 {
	result := fetchResult{URL: "https://go.dev", Status: 200, Bytes: 42}
	return result.Bytes + int64(len(result.URL))
}

//go:noinline
func escapingResult() *fetchResult {
	return &fetchResult{URL: "https://go.dev", Status: 200, Bytes: 42}
}

//go:noinline
func growResultSlice(count int) []fetchResult {
	var results []fetchResult
	for index := range count {
		results = append(results, fetchResult{Status: 200, Bytes: int64(index)})
	}
	return results
}

//go:noinline
func preallocatedResultSlice(count int) []fetchResult {
	results := make([]fetchResult, 0, count)
	for index := range count {
		results = append(results, fetchResult{Status: 200, Bytes: int64(index)})
	}
	return results
}

func TestAllocationShapeDistinguishesEscapeAndPreallocation(t *testing.T) {
	stackAllocs := testing.AllocsPerRun(1_000, func() {
		memoryIntSink = stackResultChecksum()
	})
	escapeAllocs := testing.AllocsPerRun(1_000, func() {
		memoryPointerSink = escapingResult()
	})
	growAllocs := testing.AllocsPerRun(1_000, func() {
		memorySliceSink = growResultSlice(64)
	})
	preallocatedAllocs := testing.AllocsPerRun(1_000, func() {
		memorySliceSink = preallocatedResultSlice(64)
	})

	if stackAllocs != 0 {
		t.Fatalf("stack-shaped checksum allocated %.1f times per run, want 0", stackAllocs)
	}
	if escapeAllocs < 1 {
		t.Fatalf("escaping pointer allocated %.1f times per run, want at least 1", escapeAllocs)
	}
	if growAllocs <= preallocatedAllocs {
		t.Fatalf("growing slice allocated %.1f times per run, want more than preallocated %.1f", growAllocs, preallocatedAllocs)
	}
}

func TestRuntimeMetricsExposeAllocationAndGCSignals(t *testing.T) {
	names := []string{
		"/gc/heap/allocs:bytes",
		"/gc/heap/allocs:objects",
		"/gc/heap/live:bytes",
		"/gc/cycles/total:gc-cycles",
		"/gc/limiter/last-enabled:gc-cycle",
	}
	samples := make([]metrics.Sample, len(names))
	for index, name := range names {
		samples[index].Name = name
	}
	metrics.Read(samples)

	for _, sample := range samples {
		if sample.Value.Kind() != metrics.KindUint64 {
			t.Errorf("metric %s kind = %v, want uint64", sample.Name, sample.Value.Kind())
		}
	}
}

func BenchmarkAllocationShapes(b *testing.B) {
	b.Run("stack-value", func(b *testing.B) {
		for b.Loop() {
			memoryIntSink = stackResultChecksum()
		}
	})
	b.Run("escaping-pointer", func(b *testing.B) {
		for b.Loop() {
			memoryPointerSink = escapingResult()
		}
	})
	b.Run("grow-slice", func(b *testing.B) {
		for b.Loop() {
			memorySliceSink = growResultSlice(64)
		}
	})
	b.Run("preallocated-slice", func(b *testing.B) {
		for b.Loop() {
			memorySliceSink = preallocatedResultSlice(64)
		}
	})
}
