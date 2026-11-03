package main

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{"top A", 100, "A"},
		{"bottom A", 90, "A"},
		{"top B", 89, "B"},
		{"bottom B", 80, "B"},
		{"top C", 79, "C"},
		{"bottom C", 70, "C"},
		{"top D", 69, "D"},
		{"bottom D", 60, "D"},
		{"top F", 59, "F"},
		{"bottom F", 0, "F"},
		{"negative score", -1, "invalid"},
		{"above maximum", 101, "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grade(tt.score); got != tt.want {
				t.Errorf("grade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestHTTPClass(t *testing.T) {
	tests := []struct {
		name string
		code int
		want string
	}{
		{"200", 200, "success"},
		{"204", 204, "success"},
		{"301", 301, "redirect"},
		{"404", 404, "client error"},
		{"503", 503, "server error"},
		{"100", 100, "unknown"},
		{"199", 199, "unknown"},
		{"600", 600, "unknown"},
		{"zero", 0, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := httpClass(tt.code); got != tt.want {
				t.Errorf("httpClass(%d) = %q, want %q", tt.code, got, tt.want)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name string
		n    int
		lo   int
		hi   int
		want int
	}{
		{"inside", 5, 1, 10, 5},
		{"below", 0, 1, 10, 1},
		{"above", 11, 1, 10, 10},
		{"at low bound", 1, 1, 10, 1},
		{"at high bound", 10, 1, 10, 10},
		{"negative range", -5, -5, 5, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clamp(tt.n, tt.lo, tt.hi); got != tt.want {
				t.Errorf("clamp(%d, %d, %d) = %d, want %d", tt.n, tt.lo, tt.hi, got, tt.want)
			}
		})
	}
}
