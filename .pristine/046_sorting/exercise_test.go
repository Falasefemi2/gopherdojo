package main

import (
	"slices"
	"testing"
)

func TestSortedInts(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"nil input", nil, []int{}},
		{"empty input", []int{}, []int{}},
		{"single value", []int{7}, []int{7}},
		{"unsorted", []int{3, 1, 4, 1, 5}, []int{1, 1, 3, 4, 5}},
		{"already sorted", []int{-2, 0, 5}, []int{-2, 0, 5}},
		{"negatives and duplicates", []int{0, -3, 2, -3}, []int{-3, -3, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := append([]int(nil), tt.in...)
			got := sortedInts(tt.in)
			if !slices.Equal(got, tt.want) {
				t.Errorf("sortedInts(%v) = %v, want %v", tt.in, got, tt.want)
			}
			if got == nil {
				t.Errorf("sortedInts(%v) returned nil, want a non-nil slice", tt.in)
			}
			if !slices.Equal(tt.in, before) {
				t.Errorf("sortedInts modified its input: got %v, want %v", tt.in, before)
			}
			if len(got) > 0 {
				got[0] = 9999
				if !slices.Equal(tt.in, before) {
					t.Errorf("sortedInts result aliases its input")
				}
			}
		})
	}
}

func TestSortedStrings(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"nil input", nil, []string{}},
		{"empty input", []string{}, []string{}},
		{"single value", []string{"only"}, []string{"only"}},
		{"unsorted", []string{"pear", "apple", "fig"}, []string{"apple", "fig", "pear"}},
		{"duplicates", []string{"b", "a", "b"}, []string{"a", "b", "b"}},
		{"case sensitive", []string{"banana", "Apple"}, []string{"Apple", "banana"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := append([]string(nil), tt.in...)
			got := sortedStrings(tt.in)
			if !slices.Equal(got, tt.want) {
				t.Errorf("sortedStrings(%v) = %v, want %v", tt.in, got, tt.want)
			}
			if got == nil {
				t.Errorf("sortedStrings(%v) returned nil, want a non-nil slice", tt.in)
			}
			if !slices.Equal(tt.in, before) {
				t.Errorf("sortedStrings modified its input: got %v, want %v", tt.in, before)
			}
			if len(got) > 0 {
				got[0] = "zzz"
				if !slices.Equal(tt.in, before) {
					t.Errorf("sortedStrings result aliases its input")
				}
			}
		})
	}
}
