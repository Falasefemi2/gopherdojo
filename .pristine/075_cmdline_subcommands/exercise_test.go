package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDispatch(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantOut string
		wantErr bool
	}{
		{"greet", []string{"greet", "dojo"}, "Hello, dojo!\n", false},
		{"greet unicode", []string{"greet", "café"}, "Hello, café!\n", false},
		{"add one", []string{"add", "5"}, "sum: 5\n", false},
		{"add many", []string{"add", "1", "2", "3"}, "sum: 6\n", false},
		{"add negatives", []string{"add", "-4", "10"}, "sum: 6\n", false},
		{"no args", nil, "", true},
		{"unknown", []string{"list"}, "", true},
		{"greet missing name", []string{"greet"}, "", true},
		{"greet extra arg", []string{"greet", "a", "b"}, "", true},
		{"add missing numbers", []string{"add"}, "", true},
		{"add invalid number", []string{"add", "1", "x"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := dispatch(tt.args, &buf)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("dispatch(%q) error = nil, want an error", tt.args)
				}
			} else if err != nil {
				t.Fatalf("dispatch(%q) unexpected error: %v", tt.args, err)
			}

			if got := buf.String(); got != tt.wantOut {
				t.Errorf("dispatch(%q) wrote %q, want %q", tt.args, got, tt.wantOut)
			}
		})
	}
}

func TestDispatchUnknownCommandNamesIt(t *testing.T) {
	var buf bytes.Buffer
	err := dispatch([]string{"explode"}, &buf)
	if err == nil {
		t.Fatal("dispatch([explode]) error = nil, want an error")
	}
	if !strings.Contains(err.Error(), "explode") {
		t.Errorf("dispatch([explode]) error = %q, want it to name the command", err)
	}
}
