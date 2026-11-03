package main

import "testing"

func TestIncrement(t *testing.T) {
	tests := []struct {
		name  string
		start int
		want  int
	}{
		{"zero", 0, 1},
		{"positive", 41, 42},
		{"negative", -2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.start
			increment(&n)
			if n != tt.want {
				t.Errorf("increment(&n) with n=%d left n=%d, want %d", tt.start, n, tt.want)
			}
		})
	}

	t.Run("nil is a no-op", func(t *testing.T) {
		increment(nil)
	})
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("swap(&1, &2) produced a=%d, b=%d, want a=2, b=1", a, b)
	}

	x := 7
	swap(&x, &x)
	if x != 7 {
		t.Errorf("swap(&x, &x) changed x to %d, want 7", x)
	}

	t.Run("nil left is a no-op", func(t *testing.T) {
		y := 3
		swap(nil, &y)
		if y != 3 {
			t.Errorf("swap(nil, &y) changed y to %d, want 3", y)
		}
	})

	t.Run("nil right is a no-op", func(t *testing.T) {
		z := 4
		swap(&z, nil)
		if z != 4 {
			t.Errorf("swap(&z, nil) changed z to %d, want 4", z)
		}
	})
}

func TestSetIfPositive(t *testing.T) {
	tests := []struct {
		name   string
		start  int
		v      int
		useNil bool
		wantOK bool
		want   int
	}{
		{"positive stores", 7, 5, false, true, 5},
		{"zero is ignored", 7, 0, false, false, 7},
		{"negative is ignored", 7, -3, false, false, 7},
		{"nil pointer is ignored", 7, 5, true, false, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.start
			var p *int = &n
			if tt.useNil {
				p = nil
			}

			if got := setIfPositive(p, tt.v); got != tt.wantOK {
				t.Errorf("setIfPositive(%d) = %v, want %v", tt.v, got, tt.wantOK)
			}
			if n != tt.want {
				t.Errorf("after setIfPositive(%d), value = %d, want %d", tt.v, n, tt.want)
			}
		})
	}
}
