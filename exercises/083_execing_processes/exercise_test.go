package main

import (
	"runtime"
	"strings"
	"testing"
)

func TestGoos(t *testing.T) {
	got, err := goos()
	if err != nil {
		t.Fatalf("goos() unexpected error: %v", err)
	}
	if got != runtime.GOOS {
		t.Errorf("goos() = %q, want runtime.GOOS %q", got, runtime.GOOS)
	}
	if got != strings.TrimSpace(got) {
		t.Errorf("goos() = %q, want surrounding whitespace removed", got)
	}
}
