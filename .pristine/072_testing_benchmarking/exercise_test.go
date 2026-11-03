package main

import "testing"

func TestSumSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"negative", -3, 0},
		{"one", 1, 0},
		{"four", 4, 14},
		{"ten", 10, 285},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSquares(tt.n); got != tt.want {
				t.Errorf("SumSquares(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestSumSquaresGrows(t *testing.T) {
	a, b := SumSquares(5), SumSquares(6)
	if b <= a {
		t.Errorf("SumSquares(6) = %d must exceed SumSquares(5) = %d", b, a)
	}
}

func BenchmarkSumSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSquares(100)
	}
}
