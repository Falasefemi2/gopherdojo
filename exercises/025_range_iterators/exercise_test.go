package main

import (
	"iter"
	"slices"
	"testing"
)

func collect(seq iter.Seq[int]) []int {
	var out []int
	for v := range seq {
		out = append(out, v)
	}
	return out
}

func TestCountdown(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"three", 3, []int{3, 2, 1}},
		{"one", 1, []int{1}},
		{"zero yields nothing", 0, []int{}},
		{"negative yields nothing", -4, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := collect(Countdown(tt.n))
			if !slices.Equal(got, tt.want) {
				t.Errorf("range Countdown(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"eight", 8, []int{0, 1, 1, 2, 3, 5, 8, 13}},
		{"one", 1, []int{0}},
		{"two", 2, []int{0, 1}},
		{"zero yields nothing", 0, []int{}},
		{"negative yields nothing", -3, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := collect(Fibonacci(tt.n))
			if !slices.Equal(got, tt.want) {
				t.Errorf("range Fibonacci(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestEarlyBreakStopsTheSequence(t *testing.T) {
	var got []int
	for v := range Countdown(10) {
		got = append(got, v)
		if len(got) == 3 {
			break
		}
	}
	want := []int{10, 9, 8}
	if !slices.Equal(got, want) {
		t.Errorf("breaking after 3 values gave %v, want %v", got, want)
	}

	var fibs []int
	for v := range Fibonacci(100) {
		fibs = append(fibs, v)
		if len(fibs) == 4 {
			break
		}
	}
	wantFibs := []int{0, 1, 1, 2}
	if !slices.Equal(fibs, wantFibs) {
		t.Errorf("breaking Fibonacci after 4 values gave %v, want %v", fibs, wantFibs)
	}
}
