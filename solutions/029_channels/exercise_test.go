package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"positive", 7, 7},
		{"zero", 0, 0},
		{"negative", -12, -12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundTrip(tt.in); got != tt.want {
				t.Errorf("roundTrip(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestRelayAll(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"nil input", nil, []int{}},
		{"empty input", []int{}, []int{}},
		{"single value", []int{5}, []int{5}},
		{"several values keep order", []int{5, 3, 9, 0}, []int{5, 3, 9, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := make(chan []int, 1)
			go func() { reply <- relayAll(tt.in) }()

			select {
			case got := <-reply:
				if got == nil {
					t.Fatalf("relayAll(%v) = nil, want an empty non-nil slice", tt.in)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("relayAll(%v) = %v, want %v", tt.in, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("timeout: relayAll(%v) did not return", tt.in)
			}
		})
	}
}
