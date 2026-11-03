package main

import (
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	want := "Hello from the dojo!\n"
	if got := greeting(); got != want {
		t.Errorf("greeting() = %q, want %q", got, want)
	}
}

func TestGreetingShape(t *testing.T) {
	got := greeting()
	if got == "" {
		t.Fatal("greeting() is empty; return the embedded string")
	}
	if !strings.HasSuffix(got, "\n") {
		t.Errorf("greeting() = %q, want a trailing newline", got)
	}
	if lines := strings.Split(strings.TrimSuffix(got, "\n"), "\n"); len(lines) != 1 {
		t.Errorf("greeting() spans %d lines, want 1", len(lines))
	}
}

func TestEmbeddedVariableIsLoaded(t *testing.T) {
	want := "Hello from the dojo!\n"
	if greetingFile != want {
		t.Errorf("embedded greetingFile = %q, want %q", greetingFile, want)
	}
}
