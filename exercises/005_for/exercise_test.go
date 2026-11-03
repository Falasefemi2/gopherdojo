package main

import (
	"reflect"
	"testing"
)

func TestSumTo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"five", 5, 15},
		{"ten", 10, 55},
		{"negative", -3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumTo(tt.n); got != tt.want {
				t.Errorf("sumTo(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestCountdown(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"three", 3, []int{3, 2, 1}},
		{"one", 1, []int{1}},
		{"zero", 0, []int{}},
		{"negative", -2, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countdown(tt.n)
			if got == nil {
				t.Fatalf("countdown(%d) = nil, want non-nil slice", tt.n)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countdown(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestNextPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"one", 1, 1},
		{"exact power", 2, 2},
		{"just above", 3, 4},
		{"exact sixteen", 16, 16},
		{"just above sixteen", 17, 32},
		{"zero", 0, 1},
		{"negative", -5, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPowerOfTwo(tt.n); got != tt.want {
				t.Errorf("nextPowerOfTwo(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
