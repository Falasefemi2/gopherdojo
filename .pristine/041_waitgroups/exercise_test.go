package main

import (
	"reflect"
	"testing"
)

func TestConcurrentTotal(t *testing.T) {
	tests := []struct {
		name string
		n    int
		step int
		want int64
	}{
		{"no goroutines", 0, 5, 0},
		{"negative count", -3, 2, 0},
		{"one goroutine", 1, 3, 3},
		{"many goroutines", 100, 2, 200},
		{"odd number", 7, 3, 21},
		{"negative step", 4, -2, -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := concurrentTotal(tt.n, tt.step)
			if got != tt.want {
				t.Errorf("concurrentTotal(%d, %d) = %d, want %d", tt.n, tt.step, got, tt.want)
			}
		})
	}
}

func TestWriteIndexes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"zero", 0, []int{}},
		{"one", 1, []int{0}},
		{"four", 4, []int{0, 10, 20, 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := writeIndexes(tt.n)
			if got == nil {
				t.Fatalf("writeIndexes(%d) = nil, want an empty non-nil slice", tt.n)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("writeIndexes(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
