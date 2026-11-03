package main

import "testing"

func TestClassifyAge(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want string
	}{
		{"lower boundary", 0, "minor"},
		{"last minor", 17, "minor"},
		{"first adult", 18, "adult"},
		{"last adult", 64, "adult"},
		{"first senior", 65, "senior"},
		{"upper boundary", 130, "senior"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classifyAge(tt.age); got != tt.want {
				t.Errorf("classifyAge(%d) = %q, want %q", tt.age, got, tt.want)
			}
		})
	}
}

func TestClassifyAgePanics(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want string
	}{
		{"negative age", -1, "age must be >= 0"},
		{"very negative age", -100, "age must be >= 0"},
		{"too old", 131, "age must be <= 130"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := recoverPanic(func() {
				classifyAge(tt.age)
			})
			if got != tt.want {
				t.Errorf("classifyAge(%d) panic = %v, want %q", tt.age, got, tt.want)
			}
		})
	}
}

// recoverPanic runs fn and returns the recovered panic value, or nil when fn
// returns normally. The deferred recover keeps the panic inside the test.
func recoverPanic(fn func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}
