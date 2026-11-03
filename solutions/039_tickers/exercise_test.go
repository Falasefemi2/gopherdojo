package main

import (
	"testing"
	"time"
)

func TestCollectTicks(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantLen int
	}{
		{"zero ticks", 0, 0},
		{"one tick", 1, 1},
		{"five ticks", 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := make(chan []time.Time, 1)
			go func() { reply <- collectTicks(tt.n, time.Millisecond) }()

			select {
			case got := <-reply:
				if got == nil {
					t.Fatalf("collectTicks(%d, 1ms) = nil, want an empty non-nil slice", tt.n)
				}
				if len(got) != tt.wantLen {
					t.Fatalf("collectTicks(%d, 1ms) returned %d ticks, want %d", tt.n, len(got), tt.wantLen)
				}
				for i := 1; i < len(got); i++ {
					if got[i].Before(got[i-1]) {
						t.Errorf("tick %d at %v is before tick %d at %v", i, got[i], i-1, got[i-1])
					}
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("timeout: collectTicks(%d, 1ms) did not return", tt.n)
			}
		})
	}
}

func TestCollectTicksExactCount(t *testing.T) {
	reply := make(chan []time.Time, 1)
	go func() { reply <- collectTicks(3, time.Millisecond) }()

	select {
	case got := <-reply:
		if len(got) != 3 {
			t.Fatalf("collectTicks(3, 1ms) returned %d ticks, want exactly 3", len(got))
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout: collectTicks(3, 1ms) did not return")
	}
}
