package main

import "testing"

func TestByteLen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"ascii", "Go", 2},
		{"accented", "héllo", 6},
		{"japanese", "日本語", 9},
		{"emoji", "a🙂", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteLen(tt.s); got != tt.want {
				t.Errorf("byteLen(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestRuneLen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"ascii", "Go", 2},
		{"accented", "héllo", 5},
		{"japanese", "日本語", 3},
		{"emoji", "a🙂", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeLen(tt.s); got != tt.want {
				t.Errorf("runeLen(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestReverseRunes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"ascii", "Go", "oG"},
		{"accented", "héllo", "olléh"},
		{"japanese", "日本語", "語本日"},
		{"emoji", "a🙂b", "b🙂a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseRunes(tt.s); got != tt.want {
				t.Errorf("reverseRunes(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
