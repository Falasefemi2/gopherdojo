package main

import (
	"strings"
	"testing"
)

func TestSHA256Hex(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name: "hello",
			in:   "hello",
			want: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			name: "abc",
			in:   "abc",
			want: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		},
		{
			name: "trailing newline matters",
			in:   "hello\n",
			want: "5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03",
		},
		{
			name: "pangram",
			in:   "The quick brown fox jumps over the lazy dog",
			want: "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sha256Hex(tt.in); got != tt.want {
				t.Errorf("sha256Hex(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestSHA256HexShape(t *testing.T) {
	got := sha256Hex("café")
	if len(got) != 64 {
		t.Errorf("sha256Hex returned %d characters, want 64", len(got))
	}
	if strings.ToLower(got) != got {
		t.Errorf("sha256Hex(%q) = %q, want lowercase hex", "café", got)
	}
	for i, r := range got {
		if !strings.ContainsRune("0123456789abcdef", r) {
			t.Errorf("sha256Hex contains %q at position %d, want only hex digits", r, i)
		}
	}
}

func TestSHA256HexDeterministic(t *testing.T) {
	first := sha256Hex("payload")
	second := sha256Hex("payload")
	if first != second {
		t.Errorf("same input gave different digests: %q and %q", first, second)
	}
	if first == sha256Hex("payloaD") {
		t.Errorf("inputs differing only by case produced the same digest %q", first)
	}
}
