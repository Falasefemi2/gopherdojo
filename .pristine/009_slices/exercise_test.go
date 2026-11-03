package main

import (
	"reflect"
	"testing"
)

func TestAppendAll(t *testing.T) {
	tests := []struct {
		name  string
		s     []int
		extra []int
		want  []int
	}{
		{"both non-empty", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{"nil plus values", nil, []int{1}, []int{1}},
		{"values plus nil", []int{1}, nil, []int{1}},
		{"both empty", []int{}, []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appendAll(tt.s, tt.extra)
			if got == nil {
				t.Fatalf("appendAll(%v, %v) = nil, want non-nil slice", tt.s, tt.extra)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendAll(%v, %v) = %v, want %v", tt.s, tt.extra, got, tt.want)
			}
		})
	}
}

func TestAppendAllDoesNotAlias(t *testing.T) {
	base := make([]int, 2, 8)
	base[0], base[1] = 1, 2

	got := appendAll(base, []int{3})
	got[0] = 99

	if base[0] != 1 {
		t.Errorf("appendAll shares its backing array with the input: base[0] = %d, want 1", base[0])
	}
}

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"odd length", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"even length", []int{1, 2}, []int{2, 1}},
		{"single", []int{7}, []int{7}},
		{"empty", []int{}, []int{}},
		{"nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseInPlace(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Errorf("reverseInPlace(%v) left %v, want %v", tt.in, tt.in, tt.want)
			}
		})
	}
}

func TestReverseInPlaceAliases(t *testing.T) {
	backing := []int{1, 2, 3, 4}
	view := backing[:3]

	reverseInPlace(view)

	if backing[0] != 3 {
		t.Errorf("reverseInPlace must mutate the caller's elements: backing[0] = %d, want 3", backing[0])
	}
}

func TestEvens(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"some evens", []int{2, 4, 5, 6}, []int{2, 4, 6}},
		{"no evens", []int{1, 3, 5}, []int{}},
		{"nil input", nil, []int{}},
		{"zero and negatives", []int{0, -2, 3}, []int{0, -2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := evens(tt.in)
			if got == nil {
				t.Fatalf("evens(%v) = nil, want non-nil slice", tt.in)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evens(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
