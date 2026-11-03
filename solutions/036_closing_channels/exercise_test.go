package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCloseAfter(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"empty", nil},
		{"single value", []int{4}},
		{"several values", []int{10, 20, 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := closeAfter(tt.vals)

			got := make([]int, 0, len(tt.vals))
		drain:
			for {
				select {
				case v, ok := <-ch:
					if !ok {
						break drain
					}
					got = append(got, v)
				case <-time.After(2 * time.Second):
					t.Fatal("timeout: channel was not closed after its values")
				}
			}

			want := tt.vals
			if want == nil {
				want = []int{}
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("closeAfter(%v) yielded %v, want %v", tt.vals, got, want)
			}

			for i := 0; i < 2; i++ {
				select {
				case v, ok := <-ch:
					if ok {
						t.Fatalf("receive %d after close = (%d, true), want ok == false", i, v)
					}
					if v != 0 {
						t.Errorf("receive %d after close = %d, want the zero value", i, v)
					}
				case <-time.After(2 * time.Second):
					t.Fatal("timeout: receive from a closed channel blocked")
				}
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"empty", nil},
		{"single value", []int{7}},
		{"several values", []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int, len(tt.vals))
			for _, v := range tt.vals {
				ch <- v
			}
			close(ch)

			got := readAll(ch)
			want := tt.vals
			if want == nil {
				want = []int{}
			}
			if got == nil {
				t.Fatalf("readAll(%v) = nil, want an empty non-nil slice", tt.vals)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("readAll(%v) = %v, want %v", tt.vals, got, want)
			}
		})
	}
}
