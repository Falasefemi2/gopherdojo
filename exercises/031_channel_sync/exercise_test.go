package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitFor(t *testing.T) {
	tests := []struct {
		name string
		work func() int
		want int
	}{
		{"positive result", func() int { return 42 }, 42},
		{"zero result", func() int { return 0 }, 0},
		{"negative result", func() int { return -7 }, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waitFor(tt.work); got != tt.want {
				t.Errorf("waitFor(work) = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestWaitForRunsWorkOnce(t *testing.T) {
	var calls atomic.Int32

	if got := waitFor(func() int { calls.Add(1); return 1 }); got != 1 {
		t.Fatalf("waitFor(work) = %d, want 1", got)
	}
	if got := calls.Load(); got != 1 {
		t.Errorf("work ran %d times, want exactly 1", got)
	}
}

func TestSignalWhenDone(t *testing.T) {
	var ran atomic.Bool
	done := make(chan struct{})

	signalWhenDone(func() { ran.Store(true) }, done)

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout: done channel was not closed")
	}

	if !ran.Load() {
		t.Error("done was closed before work ran")
	}

	for i := 0; i < 3; i++ {
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			t.Fatalf("receive %d from the closed done channel blocked", i)
		}
	}
}
