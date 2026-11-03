package main

import (
	"reflect"
	"testing"
	"time"
)

func collectInts(t *testing.T, ch <-chan int) []int {
	t.Helper()

	got := make([]int, 0)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-time.After(2 * time.Second):
			t.Fatal("timeout: channel neither delivered a value nor closed")
		}
	}
}

func TestProduce(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"zero", 0, []int{}},
		{"one", 1, []int{1}},
		{"three", 3, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int, tt.n)
			produce(tt.n, ch)
			if got := collectInts(t, ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("produce(%d) sent %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestConsume(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", nil, []int{}},
		{"single value", []int{9}, []int{9}},
		{"several values", []int{4, 5, 6}, []int{4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int, len(tt.in))
			for _, v := range tt.in {
				ch <- v
			}
			close(ch)

			got := consume(ch)
			if got == nil {
				t.Fatalf("consume(%v) = nil, want an empty non-nil slice", tt.in)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("consume(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestSumProduced(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"five", 5, 15},
		{"ten", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumProduced(tt.n); got != tt.want {
				t.Errorf("sumProduced(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
