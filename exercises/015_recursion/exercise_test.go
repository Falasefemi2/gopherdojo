package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 1},
		{"one", 1, 1},
		{"five", 5, 120},
		{"ten", 10, 3628800},
		{"negative is invalid", -3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.n); got != tt.want {
				t.Errorf("factorial(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 1},
		{"seven", 7, 13},
		{"ten", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestSumDigits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"single digit", 7, 7},
		{"no carry", 1234, 10},
		{"negative uses absolute value", -99, 18},
		{"trailing zeros", 1000, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigits(tt.n); got != tt.want {
				t.Errorf("sumDigits(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
