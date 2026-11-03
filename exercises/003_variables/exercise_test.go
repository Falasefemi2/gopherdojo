package main

import "testing"

func TestSwap(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
	}{
		{"positive", 3, 7},
		{"equal", 0, 0},
		{"negative and positive", -1, 5},
		{"positive and negative", 5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := swap(tt.a, tt.b)
			if gotA != tt.b || gotB != tt.a {
				t.Errorf("swap(%d, %d) = (%d, %d), want (%d, %d)",
					tt.a, tt.b, gotA, gotB, tt.b, tt.a)
			}
		})
	}
}

func TestRotate3(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c int
		wantA   int
		wantB   int
		wantC   int
	}{
		{"ascending", 1, 2, 3, 2, 3, 1},
		{"all zero", 0, 0, 0, 0, 0, 0},
		{"mixed signs", -1, 0, 1, 0, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB, gotC := rotate3(tt.a, tt.b, tt.c)
			if gotA != tt.wantA || gotB != tt.wantB || gotC != tt.wantC {
				t.Errorf("rotate3(%d, %d, %d) = (%d, %d, %d), want (%d, %d, %d)",
					tt.a, tt.b, tt.c, gotA, gotB, gotC, tt.wantA, tt.wantB, tt.wantC)
			}
		})
	}
}

func TestSwapReport(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want string
	}{
		{"positive", 3, 7, "before: 3, 7; after: 7, 3"},
		{"zero and negative", -2, 0, "before: -2, 0; after: 0, -2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapReport(tt.a, tt.b); got != tt.want {
				t.Errorf("swapReport(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
