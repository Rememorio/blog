package main

import "testing"

func TestSliceHeaderCopySharesBackingArray(t *testing.T) {
	original := make([]string, 2, 4)
	original[0] = "alpha"
	copyOfHeader := original

	copyOfHeader[0] = "changed"
	if original[0] != "changed" {
		t.Fatalf("original[0] = %q, want shared mutation", original[0])
	}

	copyOfHeader = append(copyOfHeader, "gamma")
	if len(original) != 2 {
		t.Fatalf("len(original) = %d, want 2", len(original))
	}
	if original[:3][2] != "gamma" {
		t.Fatalf("original backing array did not receive appended value")
	}
}

func TestAppendMaySplitBackingArray(t *testing.T) {
	original := []int{1}
	alias := original
	grown := append(original, 2)

	grown[0] = 99
	if alias[0] != 1 {
		t.Fatalf("alias[0] = %d, want old backing array value 1", alias[0])
	}
	if grown[0] != 99 {
		t.Fatalf("grown[0] = %d, want 99", grown[0])
	}
}

func TestChannelSendCopiesElementValue(t *testing.T) {
	type envelope struct {
		status int
		body   []byte
	}

	results := make(chan envelope, 1)
	sent := envelope{status: 201, body: []byte("go")}
	results <- sent

	sent.status = 500
	sent.body[0] = 'n'
	received := <-results

	if received.status != 201 {
		t.Fatalf("received.status = %d, want copied value 201", received.status)
	}
	if string(received.body) != "no" {
		t.Fatalf("received.body = %q, want mutation through shared slice backing array", received.body)
	}
}

func TestMapAssignmentSharesState(t *testing.T) {
	first := map[string]int{"ok": 1}
	second := first
	second["ok"] = 2

	if first["ok"] != 2 {
		t.Fatalf("first[ok] = %d, want shared map update 2", first["ok"])
	}
}

func TestInterfaceHoldingTypedNilIsNotNil(t *testing.T) {
	var pointer *fetchResult
	var value any = pointer

	if value == nil {
		t.Fatal("interface containing (*fetchResult)(nil) unexpectedly equals nil")
	}
	if typed, ok := value.(*fetchResult); !ok || typed != nil {
		t.Fatalf("type assertion = (%v, %v), want (nil, true)", typed, ok)
	}
}

func TestRangeVariablesArePerIteration(t *testing.T) {
	values := []string{"alpha", "beta", "gamma"}
	pointers := make([]*string, 0, len(values))
	for _, value := range values {
		pointers = append(pointers, &value)
	}

	for i, pointer := range pointers {
		if *pointer != values[i] {
			t.Fatalf("pointers[%d] = %q, want %q", i, *pointer, values[i])
		}
		if i > 0 && pointer == pointers[i-1] {
			t.Fatalf("range iterations %d and %d reused one variable", i-1, i)
		}
	}
}
