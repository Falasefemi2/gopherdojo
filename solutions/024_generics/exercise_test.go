package main

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		got  float64
		want float64
	}{
		{"zero value for empty ints", float64(Sum([]int{})), 0},
		{"ints", float64(Sum([]int{1, 2, 3})), 6},
		{"negative ints", float64(Sum([]int{-2, 5, -3})), 0},
		{"int64", float64(Sum([]int64{10, 20, 12})), 42},
		{"float64", Sum([]float64{1.5, 2.25, 0.25}), 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if math.Abs(tt.got-tt.want) > 1e-9 {
				t.Errorf("Sum() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name   string
		run    func() (int, bool)
		want   int
		wantOK bool
	}{
		{"empty", func() (int, bool) { return Max([]int{}) }, 0, false},
		{"single", func() (int, bool) { return Max([]int{7}) }, 7, true},
		{"unsorted", func() (int, bool) { return Max([]int{3, 9, 4}) }, 9, true},
		{"all negative", func() (int, bool) { return Max([]int{-5, -2, -9}) }, -2, true},
		{"int64", func() (int, bool) {
			v, ok := Max([]int64{4, 11, 6})
			return int(v), ok
		}, 11, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.run()
			if ok != tt.wantOK {
				t.Fatalf("Max() ok = %v, want %v", ok, tt.wantOK)
			}
			if got != tt.want {
				t.Errorf("Max() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxFloat(t *testing.T) {
	got, ok := Max([]float64{1.5, -2, 4.25})
	if !ok {
		t.Fatal("Max([]float64{1.5, -2, 4.25}) ok = false, want true")
	}
	if want := 4.25; math.Abs(got-want) > 1e-9 {
		t.Errorf("Max(...) = %v, want %v", got, want)
	}

	if got, ok := Max([]float64{}); ok {
		t.Errorf("Max(empty) = %v, true, want zero value and false", got)
	}
}

func TestMap(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		got := Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n * 2) })
		want := []string{"2", "4", "6"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map(ints) = %v, want %v", got, want)
		}
	})

	t.Run("string to int", func(t *testing.T) {
		got := Map([]string{"a", "bb", "ccc"}, func(s string) int { return len(s) })
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map(strings) = %v, want %v", got, want)
		}
	})

	t.Run("empty is non-nil", func(t *testing.T) {
		got := Map([]int{}, func(n int) int { return n })
		if got == nil {
			t.Fatal("Map(empty) = nil, want an empty, non-nil slice")
		}
		if len(got) != 0 {
			t.Errorf("Map(empty) has length %d, want 0", len(got))
		}
	})
}
