package main

import (
	"slices"
	"testing"
)

func TestUnwind(t *testing.T) {
	tests := []struct {
		name   string
		labels []string
		want   []string
	}{
		{"nil input", nil, []string{}},
		{"empty input", []string{}, []string{}},
		{"single label", []string{"only"}, []string{"only"}},
		{"three labels", []string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{"duplicates", []string{"x", "x", "y"}, []string{"y", "x", "x"}},
		{
			"resource cleanup",
			[]string{"open", "read", "close"},
			[]string{"close", "read", "open"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := append([]string(nil), tt.labels...)
			got := unwind(tt.labels)
			if !slices.Equal(got, tt.want) {
				t.Errorf("unwind(%v) = %v, want %v", tt.labels, got, tt.want)
			}
			if got == nil {
				t.Errorf("unwind(%v) returned nil, want a non-nil slice", tt.labels)
			}
			if !slices.Equal(tt.labels, before) {
				t.Errorf("unwind modified its input: got %v, want %v", tt.labels, before)
			}
		})
	}
}
