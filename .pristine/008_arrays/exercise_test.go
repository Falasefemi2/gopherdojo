package main

import "testing"

func TestSum5(t *testing.T) {
	tests := []struct {
		name string
		in   [5]int
		want int
	}{
		{"ascending", [5]int{1, 2, 3, 4, 5}, 15},
		{"all zero", [5]int{}, 0},
		{"mixed signs", [5]int{-1, -2, 3, 4, 5}, 9},
		{"repeated", [5]int{7, 7, 7, 7, 7}, 35},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum5(tt.in); got != tt.want {
				t.Errorf("sum5(%v) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestDouble5(t *testing.T) {
	tests := []struct {
		name string
		in   [5]int
		want [5]int
	}{
		{"ascending", [5]int{1, 2, 3, 4, 5}, [5]int{2, 4, 6, 8, 10}},
		{"all zero", [5]int{}, [5]int{}},
		{"mixed signs", [5]int{-1, 0, 1, 2, 3}, [5]int{-2, 0, 2, 4, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := tt.in
			got := double5(tt.in)
			if got != tt.want {
				t.Errorf("double5(%v) = %v, want %v", tt.in, got, tt.want)
			}
			if tt.in != original {
				t.Errorf("double5 mutated its input: %v, want %v", tt.in, original)
			}
		})
	}
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		in   [5]int
		want [5]int
	}{
		{"ascending", [5]int{1, 2, 3, 4, 5}, [5]int{5, 1, 2, 3, 4}},
		{"all zero", [5]int{}, [5]int{}},
		{"repeated", [5]int{1, 1, 2, 2, 3}, [5]int{3, 1, 1, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := tt.in
			got := rotateRight(tt.in)
			if got != tt.want {
				t.Errorf("rotateRight(%v) = %v, want %v", tt.in, got, tt.want)
			}
			if tt.in != original {
				t.Errorf("rotateRight mutated its input: %v, want %v", tt.in, original)
			}
		})
	}
}
