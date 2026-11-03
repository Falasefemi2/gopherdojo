package main

import (
	"reflect"
	"testing"
)

func TestTally(t *testing.T) {
	tests := []struct {
		name   string
		deltas []int
		want   []int
	}{
		{"nil input", nil, []int{}},
		{"empty input", []int{}, []int{}},
		{"single delta", []int{5}, []int{5}},
		{"positive deltas", []int{1, 2, 3}, []int{1, 3, 6}},
		{"mixed deltas", []int{10, -3, 5}, []int{10, 7, 12}},
		{"zero delta", []int{7, 0, -7}, []int{7, 7, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tally(tt.deltas)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tally(%v) = %v, want %v", tt.deltas, got, tt.want)
			}
			if got == nil {
				t.Errorf("tally(%v) returned a nil slice, want a non-nil slice", tt.deltas)
			}
		})
	}
}

func TestTallyManyRequests(t *testing.T) {
	deltas := make([]int, 100)
	for i := range deltas {
		deltas[i] = 2
	}

	got := tally(deltas)
	if len(got) != len(deltas) {
		t.Fatalf("tally returned %d totals, want %d", len(got), len(deltas))
	}
	for i, want := range got {
		if want != (i+1)*2 {
			t.Fatalf("tally total %d = %d, want %d", i, want, (i+1)*2)
		}
	}
}
