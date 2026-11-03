package main

import (
	"strings"
	"testing"
)

func TestGoVersion(t *testing.T) {
	out, err := goVersion()
	if err != nil {
		t.Fatalf("goVersion() unexpected error: %v", err)
	}
	if out == "" {
		t.Fatal("goVersion() is empty, want the go version line")
	}
	if !strings.HasPrefix(out, "go version") {
		t.Errorf("goVersion() = %q, want it to start with %q", out, "go version")
	}
	if out != strings.TrimSpace(out) {
		t.Errorf("goVersion() = %q, want surrounding whitespace removed", out)
	}
	if strings.Contains(out, "\n") {
		t.Errorf("goVersion() = %q, want a single line", out)
	}
}
