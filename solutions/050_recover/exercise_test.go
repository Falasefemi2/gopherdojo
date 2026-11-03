package main

import (
	"strings"
	"testing"
)

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{"exact division", 10, 2, 5, false},
		{"truncates toward zero", 7, 2, 3, false},
		{"negative rounds toward zero", -7, 2, -3, false},
		{"zero numerator", 0, 5, 0, false},
		{"divide by zero", 1, 0, 0, true},
		{"zero over zero", 0, 0, 0, true},
		{"negative divide by zero", -1, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safeDivide(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("safeDivide(%d, %d) returned nil error, want an error", tt.a, tt.b)
				}
				if !strings.Contains(err.Error(), "divide by zero") {
					t.Errorf("safeDivide(%d, %d) error = %q, want it to mention divide by zero", tt.a, tt.b, err)
				}
				if got != 0 {
					t.Errorf("safeDivide(%d, %d) result = %d, want 0 on error", tt.a, tt.b, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("safeDivide(%d, %d) returned unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("safeDivide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
