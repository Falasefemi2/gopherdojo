package main

import (
	"testing"
	"time"
)

func TestSelectReady(t *testing.T) {
	tests := []struct {
		name string
		val  int
	}{
		{"positive value", 7},
		{"zero value", 0},
		{"negative value", -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ready := make(chan int, 1)
			ready <- tt.val
			idle := make(chan int) // never written

			type result struct {
				val int
				ok  bool
			}
			reply := make(chan result, 1)
			go func() {
				v, ok := selectReady(ready, idle)
				reply <- result{val: v, ok: ok}
			}()

			select {
			case got := <-reply:
				if !got.ok {
					t.Fatalf("selectReady chose the idle channel, want the ready channel")
				}
				if got.val != tt.val {
					t.Errorf("selectReady = %d, want %d", got.val, tt.val)
				}
			case <-time.After(2 * time.Second):
				t.Fatal("timeout: selectReady blocked instead of taking the ready case")
			}
		})
	}
}

func TestSelectReadyAlreadyBuffered(t *testing.T) {
	ready := make(chan int, 1)
	ready <- 11
	idle := make(chan int)

	done := make(chan struct{})
	var got int
	var ok bool
	go func() {
		got, ok = selectReady(ready, idle)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout: a buffered ready channel did not win the select")
	}

	if !ok || got != 11 {
		t.Errorf("selectReady = (%d, %v), want (11, true)", got, ok)
	}
}
