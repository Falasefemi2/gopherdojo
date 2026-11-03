package main

import (
	"os"
	"syscall"
	"testing"
)

type fakeSignal struct{ name string }

func (f fakeSignal) Signal() {}

func (f fakeSignal) String() string { return f.name }

func TestSignalLabel(t *testing.T) {
	tests := []struct {
		name string
		sig  os.Signal
		want string
	}{
		{"interrupt", os.Interrupt, "interrupt"},
		{"terminate", syscall.SIGTERM, "terminate"},
		{"kill", syscall.SIGKILL, "kill"},
		{"hangup", syscall.SIGHUP, "hup"},
		{"nil", nil, "unknown"},
		{"other", fakeSignal{"custom"}, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := signalLabel(tt.sig); got != tt.want {
				t.Errorf("signalLabel(%v) = %q, want %q", tt.sig, got, tt.want)
			}
		})
	}
}

func TestWaitForSignal(t *testing.T) {
	ch := make(chan os.Signal, 1)
	ch <- os.Interrupt

	if got := waitForSignal(ch); got != os.Interrupt {
		t.Errorf("waitForSignal() = %v, want %v", got, os.Interrupt)
	}
}

func TestWaitForSignalClosedChannel(t *testing.T) {
	ch := make(chan os.Signal)
	close(ch)

	if got := waitForSignal(ch); got != nil {
		t.Errorf("waitForSignal(closed) = %v, want nil", got)
	}
}
