package main

import (
	"reflect"
	"testing"
)

func TestRollsIsDeterministic(t *testing.T) {
	first := rolls(42, 12)
	second := rolls(42, 12)

	if len(first) != 12 {
		t.Fatalf("rolls(42, 12) returned %d values, want 12", len(first))
	}
	if !reflect.DeepEqual(first, second) {
		t.Errorf("same seed produced different sequences:\n first  = %v\n second = %v", first, second)
	}
}

func TestRollsRangeAndVariety(t *testing.T) {
	got := rolls(2026, 40)
	if len(got) != 40 {
		t.Fatalf("rolls(2026, 40) returned %d values, want 40", len(got))
	}

	seen := make(map[int]bool)
	for i, v := range got {
		if v < 0 || v >= 100 {
			t.Errorf("rolls(2026, 40)[%d] = %d, want a value in [0, 100)", i, v)
		}
		seen[v] = true
	}
	if len(seen) < 2 {
		t.Errorf("rolls(2026, 40) = %v, want more than one distinct value", got)
	}
}

func TestRollsEmpty(t *testing.T) {
	for _, n := range []int{0, -3} {
		got := rolls(1, n)
		if got == nil {
			t.Errorf("rolls(1, %d) = nil, want a non-nil empty slice", n)
		}
		if len(got) != 0 {
			t.Errorf("rolls(1, %d) has length %d, want 0", n, len(got))
		}
	}
}

func TestRollsDifferentSeeds(t *testing.T) {
	a := rolls(1, 16)
	b := rolls(2, 16)
	if len(a) != len(b) {
		t.Fatalf("sequences have different lengths: %d and %d", len(a), len(b))
	}
	if reflect.DeepEqual(a, b) {
		t.Errorf("seeds 1 and 2 produced the same sequence %v", a)
	}
}
