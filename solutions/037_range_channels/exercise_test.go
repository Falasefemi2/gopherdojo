package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSquareChan(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"zero", 0, []int{}},
		{"one", 1, []int{1}},
		{"four", 4, []int{1, 4, 9, 16}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := squareChan(tt.n)

			reply := make(chan []int, 1)
			go func() {
				got := make([]int, 0, tt.n)
				for v := range ch {
					got = append(got, v)
				}
				reply <- got
			}()

			select {
			case got := <-reply:
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("squareChan(%d) yielded %v, want %v", tt.n, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("timeout: squareChan(%d) never closed its channel", tt.n)
			}
		})
	}
}

func TestSumChannel(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"empty stream", 0, 0},
		{"one value", 1, 1},
		{"four values", 4, 30},
		{"five values", 5, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := make(chan int, 1)
			go func() { reply <- sumChannel(squareChan(tt.n)) }()

			select {
			case got := <-reply:
				if got != tt.want {
					t.Errorf("sumChannel(squareChan(%d)) = %d, want %d", tt.n, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("timeout: sumChannel(squareChan(%d)) did not return", tt.n)
			}
		})
	}
}
