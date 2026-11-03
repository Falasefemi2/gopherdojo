package main

import (
	"reflect"
	"testing"
)

func TestRangeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"positive", []int{1, 2, 3}, 6},
		{"mixed signs", []int{-4, 10, -6}, 0},
		{"single value", []int{42}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSum(tt.nums); got != tt.want {
				t.Errorf("rangeSum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestFirstIndex(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		target string
		want   int
	}{
		{"first element", []string{"go", "range"}, "go", 0},
		{"middle element", []string{"go", "range", "runes"}, "range", 1},
		{"missing", []string{"go", "range"}, "map", -1},
		{"empty slice", nil, "go", -1},
		{"duplicates return first", []string{"x", "y", "x"}, "x", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstIndex(tt.words, tt.target); got != tt.want {
				t.Errorf("firstIndex(%v, %q) = %d, want %d", tt.words, tt.target, got, tt.want)
			}
		})
	}
}

func TestVowelCount(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"lowercase", "hello", 2},
		{"uppercase", "AEIOU", 5},
		{"no vowels", "rhythm", 0},
		{"punctuation ignored", "Go, range!", 3},
		{"non-ASCII is not a vowel", "café", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelCount(tt.s); got != tt.want {
				t.Errorf("vowelCount(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestSortedKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want []string
	}{
		{"sorted", map[string]int{"b": 2, "a": 1, "c": 3}, []string{"a", "b", "c"}},
		{"single", map[string]int{"only": 1}, []string{"only"}},
		{"empty", map[string]int{}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedKeys(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedKeys(%v) = %v, want %v", tt.m, got, tt.want)
			}
			if got == nil {
				t.Errorf("sortedKeys(%v) must return a non-nil slice", tt.m)
			}
		})
	}
}
