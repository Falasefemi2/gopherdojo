package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestExitCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"nil", nil, 0},
		{"usage", errUsage, 2},
		{"wrapped usage", fmt.Errorf("invocation failed: %w", errUsage), 2},
		{"other", errors.New("boom"), 1},
		{"typed other", fmt.Errorf("read config: %w", os.ErrNotExist), 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exitCode(tt.err); got != tt.want {
				t.Errorf("exitCode(%v) = %d, want %d", tt.err, got, tt.want)
			}
		})
	}
}
