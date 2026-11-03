package main

import "testing"

func TestSubtotal(t *testing.T) {
	tests := []struct {
		name  string
		price int
		qty   int
		want  int
	}{
		{"typical", 250, 3, 750},
		{"zero quantity", 250, 0, 0},
		{"negative quantity", 250, -1, 0},
		{"negative price", -1, 3, 0},
		{"free item", 0, 5, 0},
		{"larger order", 99, 7, 693},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtotal(tt.price, tt.qty); got != tt.want {
				t.Errorf("subtotal(%d, %d) = %d, want %d", tt.price, tt.qty, got, tt.want)
			}
		})
	}
}

func TestTotalWithTax(t *testing.T) {
	tests := []struct {
		name  string
		price int
		qty   int
		want  int
	}{
		{"round number", 1000, 1, 1100},
		{"tax truncates", 999, 1, 1098},
		{"three items", 250, 3, 825},
		{"free item", 0, 5, 0},
		{"zero quantity", 250, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalWithTax(tt.price, tt.qty); got != tt.want {
				t.Errorf("totalWithTax(%d, %d) = %d, want %d", tt.price, tt.qty, got, tt.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	if got := apply(func(n int) int { return n * n }, 4); got != 16 {
		t.Errorf("apply(square, 4) = %d, want 16", got)
	}
	if got := apply(func(n int) int { return n + 1 }, -3); got != -2 {
		t.Errorf("apply(increment, -3) = %d, want -2", got)
	}

	calls := 0
	counted := func(n int) int {
		calls++
		return n * 2
	}
	if got := apply(counted, 5); got != 10 {
		t.Errorf("apply(counted, 5) = %d, want 10", got)
	}
	if calls != 1 {
		t.Errorf("apply called f %d times, want exactly 1", calls)
	}
}
