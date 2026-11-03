package main

import (
	"slices"
	"testing"
)

func TestSumConcurrent(t *testing.T) {
	values := []int{1, 2, 3, 4}

	tests := []struct {
		name    string
		values  []int
		workers int
		want    int
	}{
		{"empty", nil, 4, 0},
		{"single value", []int{7}, 3, 7},
		{"one worker", values, 1, 10},
		{"two workers", values, 2, 10},
		{"three workers", values, 3, 10},
		{"more workers than values", []int{1, 2, 3}, 100, 6},
		{"zero workers means one", []int{1, 2, 3}, 0, 6},
		{"negative workers means one", []int{1, 2, 3}, -2, 6},
		{"mixed signs", []int{-5, 2, -3}, 2, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumConcurrent(tt.values, tt.workers); got != tt.want {
				t.Errorf("SumConcurrent(%v, %d) = %d, want %d", tt.values, tt.workers, got, tt.want)
			}
		})
	}
}

func TestSumConcurrentLarge(t *testing.T) {
	values := make([]int, 1000)
	want := 0
	for i := range values {
		values[i] = i + 1
		want += i + 1
	}

	if got := SumConcurrent(values, 8); got != want {
		t.Errorf("SumConcurrent(1..1000, 8) = %d, want %d", got, want)
	}
}

func TestMapConcurrent(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		f      func(int) int
		want   []int
	}{
		{
			name:   "squares",
			values: []int{1, 2, 3, 4},
			f:      func(n int) int { return n * n },
			want:   []int{1, 4, 9, 16},
		},
		{
			name:   "order is preserved",
			values: []int{5, 4, 3, 2, 1},
			f:      func(n int) int { return n * 10 },
			want:   []int{50, 40, 30, 20, 10},
		},
		{
			name:   "single value",
			values: []int{7},
			f:      func(n int) int { return -n },
			want:   []int{-7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapConcurrent(tt.values, tt.f)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MapConcurrent(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}

	t.Run("empty is non-nil", func(t *testing.T) {
		got := MapConcurrent(nil, func(n int) int { return n })
		if got == nil {
			t.Fatal("MapConcurrent(nil) = nil, want an empty, non-nil slice")
		}
		if len(got) != 0 {
			t.Errorf("MapConcurrent(nil) has length %d, want 0", len(got))
		}
	})
}
