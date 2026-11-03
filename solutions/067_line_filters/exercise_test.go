package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestFilterLines(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		substr string
		want   string
	}{
		{"single match", "alpha\nbeta\ngamma\n", "et", "BETA\n"},
		{"several matches keep order", "one\ntwo\nthree\nfour\n", "o", "ONE\nTWO\nFOUR\n"},
		{"case-sensitive miss", "alpha\n", "ALPHA", ""},
		{"no matches", "a\nb\n", "z", ""},
		{"empty input", "", "x", ""},
		{"last line without newline", "alpha\nbeta", "beta", "BETA\n"},
		{"empty substring keeps all", "a\nb\n", "", "A\nB\n"},
		{"unicode uppercased", "café\n", "fé", "CAFÉ\n"},
		{"crlf line endings", "alpha\r\nbeta\r\n", "alpha", "ALPHA\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			err := filterLines(strings.NewReader(tt.input), &out, tt.substr)
			if err != nil {
				t.Fatalf("filterLines(%q, %q) returned error: %v", tt.input, tt.substr, err)
			}
			if got := out.String(); got != tt.want {
				t.Errorf("filterLines(%q, %q) wrote %q, want %q", tt.input, tt.substr, got, tt.want)
			}
		})
	}
}

type failingWriter struct{}

func (failingWriter) Write(p []byte) (int, error) {
	return 0, errors.New("write failed")
}

type failingReader struct{}

func (failingReader) Read(p []byte) (int, error) {
	return 0, errors.New("read failed")
}

func TestFilterLinesWriterError(t *testing.T) {
	err := filterLines(strings.NewReader("alpha\nbeta\n"), failingWriter{}, "a")
	if err == nil {
		t.Error("filterLines with a failing writer returned nil, want an error")
	}
}

func TestFilterLinesReaderError(t *testing.T) {
	var out bytes.Buffer
	err := filterLines(failingReader{}, &out, "a")
	if err == nil {
		t.Error("filterLines with a failing reader returned nil, want an error")
	}
}
