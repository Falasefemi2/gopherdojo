package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			name: "repeated words",
			text: "the cat and the hat",
			want: map[string]int{"the": 2, "cat": 1, "and": 1, "hat": 1},
		},
		{
			name: "case and whitespace",
			text: "The  CAT\tthe\ncat",
			want: map[string]int{"the": 2, "cat": 2},
		},
		{
			name: "single word",
			text: "hello",
			want: map[string]int{"hello": 1},
		},
		{
			name: "empty text",
			text: "",
			want: map[string]int{},
		},
		{
			name: "whitespace only",
			text: " \t\n ",
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordCount(tt.text)
			if got == nil {
				t.Fatalf("wordCount(%q) = nil, want non-nil map", tt.text)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordCount(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}

func TestGetOr(t *testing.T) {
	m := map[string]int{"present": 4, "zero": 0}

	tests := []struct {
		name     string
		m        map[string]int
		key      string
		fallback int
		want     int
	}{
		{"present key", m, "present", 99, 4},
		{"stored zero", m, "zero", 99, 0},
		{"missing key", m, "absent", 99, 99},
		{"nil map", nil, "absent", 7, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOr(tt.m, tt.key, tt.fallback); got != tt.want {
				t.Errorf("getOr(%v, %q, %d) = %d, want %d", tt.m, tt.key, tt.fallback, got, tt.want)
			}
		})
	}
}

func TestRemoveKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	removeKey(m, "b")
	want := map[string]int{"a": 1, "c": 3}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("removeKey(m, %q) left %v, want %v", "b", m, want)
	}

	removeKey(m, "missing")
	if !reflect.DeepEqual(m, want) {
		t.Errorf("removeKey on a missing key changed the map: got %v, want %v", m, want)
	}
}
