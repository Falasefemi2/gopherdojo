package main

import "testing"

func TestCountHits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{"zero workers", 0, 0},
		{"one worker", 1, 1},
		{"ten workers", 10, 10},
		{"one hundred workers", 100, 100},
		{"one thousand workers", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHits(tt.n); got != tt.want {
				t.Errorf("countHits(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestCountHitsRepeatedCallsAreIndependent(t *testing.T) {
	for i := 0; i < 3; i++ {
		if got := countHits(100); got != 100 {
			t.Fatalf("call %d: countHits(100) = %d, want 100", i+1, got)
		}
	}
}
