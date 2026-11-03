package main

import (
	"slices"
	"testing"
)

func TestByLengthThenName(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{"nil input", nil, []string{}},
		{"empty input", []string{}, []string{}},
		{"single word", []string{"solo"}, []string{"solo"}},
		{
			"length then name",
			[]string{"pear", "fig", "apple", "kiwi", "fig"},
			[]string{"fig", "fig", "kiwi", "pear", "apple"},
		},
		{
			"equal lengths sort alphabetically",
			[]string{"bb", "aa", "cc"},
			[]string{"aa", "bb", "cc"},
		},
		{
			"already ordered",
			[]string{"a", "bb", "ccc"},
			[]string{"a", "bb", "ccc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := append([]string(nil), tt.words...)
			got := byLengthThenName(tt.words)
			if !slices.Equal(got, tt.want) {
				t.Errorf("byLengthThenName(%v) = %v, want %v", tt.words, got, tt.want)
			}
			if got == nil {
				t.Errorf("byLengthThenName(%v) returned nil, want a non-nil slice", tt.words)
			}
			if !slices.Equal(tt.words, before) {
				t.Errorf("byLengthThenName modified its input: got %v, want %v", tt.words, before)
			}
			if len(got) > 0 {
				got[0] = "zzzzzzzz"
				if !slices.Equal(tt.words, before) {
					t.Errorf("byLengthThenName result aliases its input")
				}
			}
		})
	}
}

func TestByLengthThenNameTieBreak(t *testing.T) {
	got := byLengthThenName([]string{"b", "a", "c", "aa", "ab"})
	want := []string{"a", "b", "c", "aa", "ab"}
	if !slices.Equal(got, want) {
		t.Fatalf("tie-break order = %v, want %v", got, want)
	}
}
