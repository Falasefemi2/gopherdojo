package main

import (
	"reflect"
	"testing"
)

func TestBufferStats(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		fill    int
		wantLen int
		wantCap int
	}{
		{"unbuffered channel", 0, 0, 0, 0},
		{"empty buffer", 3, 0, 0, 3},
		{"partly filled", 3, 2, 2, 3},
		{"completely full", 5, 5, 5, 5},
		{"single slot", 1, 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen, gotCap := bufferStats(tt.size, tt.fill)
			if gotLen != tt.wantLen || gotCap != tt.wantCap {
				t.Errorf("bufferStats(%d, %d) = (len %d, cap %d), want (len %d, cap %d)",
					tt.size, tt.fill, gotLen, gotCap, tt.wantLen, tt.wantCap)
			}
		})
	}
}

func TestBufferedRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"nil input", nil, []int{}},
		{"empty input", []int{}, []int{}},
		{"single value", []int{7}, []int{7}},
		{"several values", []int{1, 2, 3}, []int{1, 2, 3}},
		{"zero and negatives", []int{0, -1, 5}, []int{0, -1, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bufferedRoundTrip(tt.in)
			if got == nil {
				t.Fatalf("bufferedRoundTrip(%v) = nil, want an empty non-nil slice", tt.in)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bufferedRoundTrip(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
