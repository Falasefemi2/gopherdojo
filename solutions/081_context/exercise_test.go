package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFetchSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	got, err := fetch(ctx, time.Millisecond)
	if err != nil {
		t.Fatalf("fetch() unexpected error: %v", err)
	}
	if got != "payload" {
		t.Errorf("fetch() = %q, want %q", got, "payload")
	}
}

func TestFetchDeadlineWins(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	got, err := fetch(ctx, 200*time.Millisecond)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("fetch() error = %v, want context.DeadlineExceeded", err)
	}
	if got != "" {
		t.Errorf("fetch() = %q on timeout, want empty", got)
	}
}

func TestFetchCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	got, err := fetch(ctx, 50*time.Millisecond)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("fetch() error = %v, want context.Canceled", err)
	}
	if got != "" {
		t.Errorf("fetch() = %q when cancelled, want empty", got)
	}
}
