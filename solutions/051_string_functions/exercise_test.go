package main

import (
	"slices"
	"testing"
)

func TestMatchText(t *testing.T) {
	tests := []struct {
		name string
		text string
		sub  string
		want TextMatch
	}{
		{"single occurrence", "gopher", "go", TextMatch{true, 1, 0}},
		{"repeated occurrence", "go gopher go", "go", TextMatch{true, 3, 0}},
		{"non-overlapping count", "banana", "ana", TextMatch{true, 1, 1}},
		{"repeated pair", "banana", "na", TextMatch{true, 2, 2}},
		{"absent", "banana", "zz", TextMatch{false, 0, -1}},
		{"empty text", "", "x", TextMatch{false, 0, -1}},
		{"empty substring", "ab", "", TextMatch{true, 3, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchText(tt.text, tt.sub); got != tt.want {
				t.Errorf("matchText(%q, %q) = %+v, want %+v", tt.text, tt.sub, got, tt.want)
			}
		})
	}
}

func TestRepeatJoin(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		times    int
		words    []string
		sep      string
		repeated string
		joined   string
		parts    []string
	}{
		{
			name: "repeat and join", word: "ab", times: 2,
			words: []string{"x", "y"}, sep: "-",
			repeated: "abab", joined: "x-y", parts: []string{"x", "y"},
		},
		{
			name: "zero repeats", word: "ab", times: 0,
			words: []string{"a"}, sep: "|",
			repeated: "", joined: "a", parts: []string{"a"},
		},
		{
			name: "empty word list", word: "z", times: 1,
			words: []string{}, sep: "|",
			repeated: "z", joined: "", parts: []string{""},
		},
		{
			name: "single word", word: "go", times: 3,
			words: []string{"solo"}, sep: ",",
			repeated: "gogogo", joined: "solo", parts: []string{"solo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repeated, joined, parts := repeatJoin(tt.word, tt.times, tt.words, tt.sep)
			if repeated != tt.repeated {
				t.Errorf("repeated = %q, want %q", repeated, tt.repeated)
			}
			if joined != tt.joined {
				t.Errorf("joined = %q, want %q", joined, tt.joined)
			}
			if !slices.Equal(parts, tt.parts) {
				t.Errorf("parts = %q, want %q", parts, tt.parts)
			}
		})
	}
}

func TestNormalizeTitle(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		prefix     string
		normalized string
		hasPrefix  bool
	}{
		{
			name: "trims and uppercases", title: "  go server  ", prefix: "GO-",
			normalized: "GO-SERVER", hasPrefix: true,
		},
		{
			name: "prefix must be at the start", title: "release notes", prefix: "NOTES",
			normalized: "RELEASE-NOTES", hasPrefix: false,
		},
		{
			name: "each space becomes a dash", title: "payments  api", prefix: "PAYMENTS",
			normalized: "PAYMENTS--API", hasPrefix: true,
		},
		{
			name: "empty after trim", title: "   ", prefix: "",
			normalized: "", hasPrefix: true,
		},
		{
			name: "unicode uppercase", title: "café menu", prefix: "CAFÉ-",
			normalized: "CAFÉ-MENU", hasPrefix: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized, hasPrefix := normalizeTitle(tt.title, tt.prefix)
			if normalized != tt.normalized {
				t.Errorf("normalizeTitle(%q, %q) normalized = %q, want %q",
					tt.title, tt.prefix, normalized, tt.normalized)
			}
			if hasPrefix != tt.hasPrefix {
				t.Errorf("normalizeTitle(%q, %q) hasPrefix = %v, want %v",
					tt.title, tt.prefix, hasPrefix, tt.hasPrefix)
			}
		})
	}
}
