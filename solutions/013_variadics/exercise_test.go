package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"no arguments", nil, 0},
		{"single", []int{5}, 5},
		{"ascending", []int{1, 2, 3}, 6},
		{"with negatives", []int{-1, 2, -3}, -2},
		{"empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.nums...); got != tt.want {
				t.Errorf("sum(%v...) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMaxOf(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		want   int
		wantOK bool
	}{
		{"empty", nil, 0, false},
		{"single", []int{3}, 3, true},
		{"last is best", []int{1, 9, 4, 9}, 9, true},
		{"all negative", []int{-5, -2, -9}, -2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := maxOf(tt.nums...)
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("maxOf(%v...) = (%d, %t), want (%d, %t)",
					tt.nums, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestJoinInts(t *testing.T) {
	tests := []struct {
		name string
		sep  string
		nums []int
		want string
	}{
		{"dash", "-", []int{1, 2, 3}, "1-2-3"},
		{"comma and space", ", ", []int{4, 5}, "4, 5"},
		{"no numbers", ",", nil, ""},
		{"single number", "-", []int{7}, "7"},
		{"negatives", "|", []int{-1, 0, 1}, "-1|0|1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinInts(tt.sep, tt.nums...); got != tt.want {
				t.Errorf("joinInts(%q, %v...) = %q, want %q", tt.sep, tt.nums, got, tt.want)
			}
		})
	}
}

func TestJoinIntsSpreadMatchesElements(t *testing.T) {
	nums := []int{4, 5, 6}

	spread := joinInts("+", nums...)
	listed := joinInts("+", 4, 5, 6)

	if spread != listed {
		t.Errorf("joinInts spread = %q, listed = %q; want the same result", spread, listed)
	}
}
