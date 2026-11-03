package main

import (
	"testing"
)

func TestEncodeStd(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"hello", "hello", "aGVsbG8="},
		{"empty", "", ""},
		{"plus and slash", "\xfb\xff", "+/8="},
		{"single byte", "a", "YQ=="},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeStd(tt.in); got != tt.want {
				t.Errorf("encodeStd(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestEncodeURL(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"hello", "hello", "aGVsbG8="},
		{"empty", "", ""},
		{"url safe alphabet", "\xfb\xff", "-_8="},
		{"no padding chars needed", "abc", "YWJj"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeURL(tt.in); got != tt.want {
				t.Errorf("encodeURL(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestDecodeRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"text", "hello, gopher"},
		{"empty", ""},
		{"binary", "\x00\x01\xfe\xff"},
		{"unicode", "café ☕"},
	}

	for _, tt := range tests {
		t.Run("std "+tt.name, func(t *testing.T) {
			got, err := decodeStd(encodeStd(tt.in))
			if err != nil {
				t.Fatalf("decodeStd(encodeStd(%q)) returned error: %v", tt.in, err)
			}
			if got != tt.in {
				t.Errorf("std roundtrip = %q, want %q", got, tt.in)
			}
		})

		t.Run("url "+tt.name, func(t *testing.T) {
			got, err := decodeURL(encodeURL(tt.in))
			if err != nil {
				t.Fatalf("decodeURL(encodeURL(%q)) returned error: %v", tt.in, err)
			}
			if got != tt.in {
				t.Errorf("url roundtrip = %q, want %q", got, tt.in)
			}
		})
	}
}

func TestDecodeKnown(t *testing.T) {
	if got, err := decodeStd("aGVsbG8="); err != nil || got != "hello" {
		t.Errorf("decodeStd(aGVsbG8=) = (%q, %v), want (\"hello\", nil)", got, err)
	}
	if got, err := decodeURL("-_8="); err != nil || got != "\xfb\xff" {
		t.Errorf("decodeURL(-_8=) = (%q, %v), want (binary, nil)", got, err)
	}
}

func TestDecodeErrors(t *testing.T) {
	tests := []struct {
		name   string
		decode func(string) (string, error)
		in     string
	}{
		{"std illegal character", decodeStd, "!!!!"},
		{"std bad length", decodeStd, "a"},
		{"std url alphabet rejected", decodeStd, "-_8="},
		{"url std alphabet rejected", decodeURL, "+/8="},
		{"url bad length", decodeURL, "abcde"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.decode(tt.in)
			if err == nil {
				t.Errorf("decode(%q) error = nil, want an error", tt.in)
			}
			if got != "" {
				t.Errorf("decode(%q) = %q, want the empty string on error", tt.in, got)
			}
		})
	}
}
