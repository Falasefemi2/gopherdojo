package main

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
	}{
		{"five milliseconds", 5 * time.Millisecond},
		{"twenty milliseconds", 20 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timer := newTimer(tt.d)
			if timer == nil {
				t.Fatal("newTimer returned nil")
			}
			defer timer.Stop()

			select {
			case <-timer.C:
			case <-time.After(2 * time.Second):
				t.Fatalf("timer for %v did not fire", tt.d)
			}
		})
	}
}

func TestResetTimer(t *testing.T) {
	timer := newTimer(time.Hour)
	if timer == nil {
		t.Fatal("newTimer returned nil")
	}
	defer timer.Stop()

	got := resetTimer(timer, 5*time.Millisecond)
	if got != timer {
		t.Error("resetTimer returned a different timer, want the same one")
	}

	select {
	case <-got.C:
	case <-time.After(2 * time.Second):
		t.Fatal("reset timer did not fire within 2s")
	}
}

func TestStoppedTimerDoesNotFire(t *testing.T) {
	timer := newTimer(5 * time.Millisecond)
	if timer == nil {
		t.Fatal("newTimer returned nil")
	}
	if !timer.Stop() {
		// The timer may already have fired; drain it so the test stays honest.
		select {
		case <-timer.C:
		default:
		}
	}

	select {
	case <-timer.C:
		t.Fatal("stopped timer delivered a tick")
	case <-time.After(20 * time.Millisecond):
	}
}
