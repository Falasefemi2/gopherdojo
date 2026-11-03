package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func receiveInt(t *testing.T, ch <-chan int) (int, bool) {
	t.Helper()

	select {
	case v, ok := <-ch:
		return v, ok
	case <-time.After(2 * time.Second):
		t.Fatal("timeout: expected a value on the channel")
		return 0, false
	}
}

func TestTrySendBuffered(t *testing.T) {
	ch := make(chan int, 1)

	if !trySend(ch, 5) {
		t.Fatal("trySend on an empty buffered channel = false, want true")
	}
	if trySend(ch, 6) {
		t.Fatal("trySend on a full buffered channel = true, want false")
	}

	v, ok := receiveInt(t, ch)
	if !ok || v != 5 {
		t.Errorf("buffered value = (%d, %v), want (5, true)", v, ok)
	}
}

func TestTrySendUnbufferedWithoutReceiver(t *testing.T) {
	ch := make(chan int)
	var sent atomic.Bool

	done := make(chan struct{})
	go func() {
		sent.Store(trySend(ch, 1))
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout: trySend blocked instead of falling through to default")
	}

	if sent.Load() {
		t.Error("trySend with no receiver = true, want false")
	}
}

func TestTryReceive(t *testing.T) {
	tests := []struct {
		name   string
		ch     func() <-chan int
		wantV  int
		wantOK bool
	}{
		{
			name: "buffered value",
			ch: func() <-chan int {
				ch := make(chan int, 2)
				ch <- 9
				return ch
			},
			wantV:  9,
			wantOK: true,
		},
		{
			name:   "empty channel",
			ch:     func() <-chan int { return make(chan int, 1) },
			wantV:  0,
			wantOK: false,
		},
		{
			name: "closed channel",
			ch: func() <-chan int {
				ch := make(chan int, 1)
				close(ch)
				return ch
			},
			wantV:  0,
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotOK := tryReceive(tt.ch())
			if gotV != tt.wantV || gotOK != tt.wantOK {
				t.Errorf("tryReceive = (%d, %v), want (%d, %v)", gotV, gotOK, tt.wantV, tt.wantOK)
			}
		})
	}
}

func TestTryReceiveSequential(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	for _, want := range []int{1, 2} {
		got, ok := tryReceive(ch)
		if !ok || got != want {
			t.Fatalf("tryReceive = (%d, %v), want (%d, true)", got, ok, want)
		}
	}

	if got, ok := tryReceive(ch); ok {
		t.Errorf("tryReceive on an emptied channel = (%d, true), want (0, false)", got)
	}
}
