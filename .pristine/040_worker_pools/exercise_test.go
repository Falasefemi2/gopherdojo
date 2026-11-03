package main

import (
	"testing"
	"time"
)

func TestPoolSquare(t *testing.T) {
	tests := []struct {
		name    string
		jobs    []int
		workers int
		want    int
	}{
		{"single worker", []int{1, 2, 3, 4}, 1, 30},
		{"two workers", []int{1, 2, 3, 4}, 2, 30},
		{"more workers than jobs", []int{2, 3}, 5, 13},
		{"zero workers means one", []int{1, 2}, 0, 5},
		{"negative workers means one", []int{3}, -2, 9},
		{"empty job list", nil, 3, 0},
		{"negative jobs", []int{-1, 2}, 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := make(chan int, 1)
			go func() { reply <- poolSquare(tt.jobs, tt.workers) }()

			select {
			case got := <-reply:
				if got != tt.want {
					t.Errorf("poolSquare(%v, %d) = %d, want %d", tt.jobs, tt.workers, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("timeout: poolSquare(%v, %d) did not return", tt.jobs, tt.workers)
			}
		})
	}
}
