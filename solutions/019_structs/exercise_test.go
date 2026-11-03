package main

import "testing"

func TestNewPoint(t *testing.T) {
	tests := []struct {
		name string
		x, y int
		want Point
	}{
		{"origin", 0, 0, Point{}},
		{"positive", 3, 4, Point{X: 3, Y: 4}},
		{"negative", -2, 5, Point{X: -2, Y: 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPoint(tt.x, tt.y); got != tt.want {
				t.Errorf("newPoint(%d, %d) = %+v, want %+v", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		name   string
		p      Point
		dx, dy int
		want   Point
	}{
		{"move right and down", Point{X: 1, Y: 2}, 3, 4, Point{X: 4, Y: 6}},
		{"move by zero", Point{X: 5, Y: -1}, 0, 0, Point{X: 5, Y: -1}},
		{"move negative", Point{X: 0, Y: 0}, -2, 3, Point{X: -2, Y: 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := tt.p
			got := translate(tt.p, tt.dx, tt.dy)
			if got != tt.want {
				t.Errorf("translate(%+v, %d, %d) = %+v, want %+v", tt.p, tt.dx, tt.dy, got, tt.want)
			}
			if tt.p != original {
				t.Errorf("translate mutated its input: got %+v, want %+v", tt.p, original)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name string
		r    Rect
		want int
	}{
		{"unit square", Rect{Min: Point{}, Max: Point{X: 1, Y: 1}}, 1},
		{"wide", Rect{Min: Point{}, Max: Point{X: 4, Y: 2}}, 8},
		{"offset", Rect{Min: Point{X: 1, Y: 1}, Max: Point{X: 4, Y: 5}}, 12},
		{"zero size", Rect{Min: Point{X: 2, Y: 2}, Max: Point{X: 2, Y: 2}}, 0},
		{"inverted", Rect{Min: Point{X: 2, Y: 2}, Max: Point{X: 0, Y: 0}}, 0},
		{"negative height", Rect{Min: Point{X: 0, Y: 5}, Max: Point{X: 3, Y: 2}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := area(tt.r); got != tt.want {
				t.Errorf("area(%+v) = %d, want %d", tt.r, got, tt.want)
			}
		})
	}
}
