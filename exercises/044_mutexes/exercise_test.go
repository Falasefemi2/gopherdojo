package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		keys []string
		want map[string]int
	}{
		{"nil input", nil, map[string]int{}},
		{"empty input", []string{}, map[string]int{}},
		{"single key", []string{"go"}, map[string]int{"go": 1}},
		{"duplicates", []string{"go", "go", "go"}, map[string]int{"go": 3}},
		{
			"mixed keys",
			[]string{"a", "b", "a", "c", "a", "b"},
			map[string]int{"a": 3, "b": 2, "c": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.keys)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWords(%v) = %v, want %v", tt.keys, got, tt.want)
			}
			if got == nil {
				t.Errorf("countWords(%v) returned a nil map, want a non-nil map", tt.keys)
			}
		})
	}
}

func TestCountWordsExactTotal(t *testing.T) {
	keys := make([]string, 0, 200)
	for i := 0; i < 200; i++ {
		keys = append(keys, "hit")
	}

	got := countWords(keys)
	if got["hit"] != 200 {
		t.Fatalf("countWords with 200 identical keys = %d hits, want 200", got["hit"])
	}
	if len(got) != 1 {
		t.Fatalf("countWords returned %d distinct keys, want 1", len(got))
	}
}
