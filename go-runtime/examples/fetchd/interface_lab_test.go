package main

import (
	"reflect"
	"testing"
)

type resultDescriber interface {
	describe() string
}

func (r fetchResult) describe() string {
	return r.URL
}

type fetchProblem struct {
	target string
}

func (p *fetchProblem) Error() string {
	if p == nil {
		return "fetch problem"
	}
	return "fetch problem: " + p.target
}

type resultBatch []fetchResult

func filterBatch[S ~[]E, E any](values S, keep func(E) bool) S {
	filtered := make(S, 0, len(values))
	for _, value := range values {
		if keep(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func setStatusByName(target any, status int) bool {
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return false
	}

	field := value.Elem().FieldByName("Status")
	if !field.IsValid() || !field.CanSet() || field.Kind() != reflect.Int {
		return false
	}
	field.SetInt(int64(status))
	return true
}

func TestInterfaceRetainsDynamicTypeAndValue(t *testing.T) {
	want := fetchResult{URL: "https://go.dev", Status: 200}
	var value resultDescriber = want

	if got := value.describe(); got != want.URL {
		t.Fatalf("describe() = %q, want %q", got, want.URL)
	}
	if got := reflect.TypeOf(value); got != reflect.TypeFor[fetchResult]() {
		t.Fatalf("dynamic type = %v, want %v", got, reflect.TypeFor[fetchResult]())
	}
	if got, ok := reflect.TypeAssert[fetchResult](reflect.ValueOf(value)); !ok || got != want {
		t.Fatalf("TypeAssert = (%+v, %v), want (%+v, true)", got, ok, want)
	}
}

func TestTypedNilInterfaceRetainsDynamicType(t *testing.T) {
	var empty any
	if empty != nil || reflect.TypeOf(empty) != nil || reflect.ValueOf(empty).IsValid() {
		t.Fatal("nil interface unexpectedly has a dynamic type or valid reflected value")
	}

	var problem *fetchProblem
	var err error = problem
	if err == nil {
		t.Fatal("error containing (*fetchProblem)(nil) unexpectedly equals nil")
	}
	if got := reflect.TypeOf(err); got != reflect.TypeFor[*fetchProblem]() {
		t.Fatalf("dynamic type = %v, want %v", got, reflect.TypeFor[*fetchProblem]())
	}
	if got, ok := reflect.TypeAssert[*fetchProblem](reflect.ValueOf(err)); !ok || got != nil {
		t.Fatalf("TypeAssert = (%v, %v), want (nil, true)", got, ok)
	}
}

func TestGenericFilterPreservesNamedSliceType(t *testing.T) {
	input := resultBatch{
		{URL: "https://go.dev", Status: 200},
		{URL: "https://example.invalid", Status: 503},
	}

	filtered := filterBatch(input, func(result fetchResult) bool {
		return result.Status < 500
	})

	var preserved resultBatch = filtered
	if len(preserved) != 1 || preserved[0].URL != "https://go.dev" {
		t.Fatalf("filtered = %+v, want the successful result", preserved)
	}
	if got := reflect.TypeOf(filtered); got != reflect.TypeFor[resultBatch]() {
		t.Fatalf("result type = %v, want %v", got, reflect.TypeFor[resultBatch]())
	}
}

func TestReflectionNeedsAddressableExportedField(t *testing.T) {
	result := fetchResult{Status: 200}
	copyField := reflect.ValueOf(result).FieldByName("Status")
	if copyField.CanSet() {
		t.Fatal("field from a non-addressable struct copy unexpectedly reports CanSet")
	}
	if setStatusByName(result, 204) {
		t.Fatal("reflection mutated a non-pointer input")
	}
	if !setStatusByName(&result, 204) {
		t.Fatal("reflection could not mutate an addressable exported field")
	}
	if result.Status != 204 {
		t.Fatalf("result.Status = %d, want 204", result.Status)
	}
}

type byteSizer interface {
	byteSize() int64
}

func (r fetchResult) byteSize() int64 {
	return r.Bytes
}

func sumGeneric[T byteSizer](values []T) int64 {
	var total int64
	for _, value := range values {
		total += value.byteSize()
	}
	return total
}

func sumInterface(values []byteSizer) int64 {
	var total int64
	for _, value := range values {
		total += value.byteSize()
	}
	return total
}

func sumReflect(values any) int64 {
	reflected := reflect.ValueOf(values)
	var total int64
	for index := range reflected.Len() {
		total += reflected.Index(index).FieldByName("Bytes").Int()
	}
	return total
}

var abstractionBoundarySink int64

func BenchmarkAbstractionBoundaries(b *testing.B) {
	concrete := make([]fetchResult, 64)
	interfaces := make([]byteSizer, len(concrete))
	var want int64
	for index := range concrete {
		concrete[index].Bytes = int64(index + 1)
		interfaces[index] = concrete[index]
		want += concrete[index].Bytes
	}

	b.Run("generic-constraint", func(b *testing.B) {
		for b.Loop() {
			abstractionBoundarySink = sumGeneric(concrete)
		}
	})
	b.Run("interface-values", func(b *testing.B) {
		for b.Loop() {
			abstractionBoundarySink = sumInterface(interfaces)
		}
	})
	b.Run("reflection", func(b *testing.B) {
		for b.Loop() {
			abstractionBoundarySink = sumReflect(concrete)
		}
	})

	if abstractionBoundarySink != want {
		b.Fatalf("sum = %d, want %d", abstractionBoundarySink, want)
	}
}
