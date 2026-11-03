package main

import (
	"testing"
	"time"
)

func TestAwaitValue(t *testing.T) {
	tests := []struct {
		name   string
		ch     func() <-chan int
		d      time.Duration
		wantV  int
		wantOK bool
	}{
		{
			name: "value arrives in time",
			ch: func() <-chan int {
				ch := make(chan int, 1)
				ch <- 7
				return ch
			},
			d:      50 * time.Millisecond,
			wantV:  7,
			wantOK: true,
		},
		{
			name: "zero value arrives in time",
			ch: func() <-chan int {
				ch := make(chan int, 1)
				ch <- 0
				return ch
			},
			d:      50 * time.Millisecond,
			wantV:  0,
			wantOK: true,
		},
		{
			name:   "nobody sends",
			ch:     func() <-chan int { return make(chan int) },
			d:      10 * time.Millisecond,
			wantV:  0,
			wantOK: false,
		},
		{
			name: "closed channel",
			ch: func() <-chan int {
				ch := make(chan int)
				close(ch)
				return ch
			},
			d:      50 * time.Millisecond,
			wantV:  0,
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotOK := awaitValue(tt.ch(), tt.d)
			if gotV != tt.wantV || gotOK != tt.wantOK {
				t.Errorf("awaitValue(ch, %v) = (%d, %v), want (%d, %v)",
					tt.d, gotV, gotOK, tt.wantV, tt.wantOK)
			}
		})
	}
}

func TestAwaitValueTimeoutIsBounded(t *testing.T) {
	start := time.Now()
	_, ok := awaitValue(make(chan int), 10*time.Millisecond)
	if ok {
		t.Fatal("awaitValue reported a value from an empty channel")
	}
	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Fatalf("awaitValue took %v, want roughly 10ms", elapsed)
	}
}
