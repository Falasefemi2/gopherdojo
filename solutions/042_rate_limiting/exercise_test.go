package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRateLimited(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  []int
	}{
		{"five items", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"single item", []int{42}, []int{42}},
		{"repeated values", []int{7, 7, 7}, []int{7, 7, 7}},
		{"empty items", nil, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := make(chan []int, 1)
			go func() { reply <- rateLimited(tt.items, time.Millisecond) }()

			select {
			case got := <-reply:
				if got == nil {
					t.Fatalf("rateLimited(%v, 1ms) = nil, want an empty non-nil slice", tt.items)
				}
				if len(got) != len(tt.want) {
					t.Fatalf("rateLimited(%v, 1ms) returned %d items, want %d", tt.items, len(got), len(tt.want))
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("rateLimited(%v, 1ms) = %v, want %v", tt.items, got, tt.want)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("timeout: rateLimited(%v, 1ms) did not return", tt.items)
			}
		})
	}
}
