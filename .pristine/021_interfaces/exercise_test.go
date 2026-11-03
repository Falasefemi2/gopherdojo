package main

import (
	"math"
	"testing"
)

var (
	_ Shape = Rectangle{}
	_ Shape = Circle{}
)

func TestTotalArea(t *testing.T) {
	tests := []struct {
		name   string
		shapes []Shape
		want   float64
	}{
		{"nil slice", nil, 0},
		{"empty slice", []Shape{}, 0},
		{"rectangle", []Shape{Rectangle{W: 2, H: 3}}, 6},
		{"mixed basket", []Shape{Rectangle{W: 2, H: 3}, Circle{R: 1}}, 6 + math.Pi},
		{"zero area shape", []Shape{Rectangle{W: 0, H: 5}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalArea(tt.shapes)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("totalArea(%d shapes) = %v, want %v", len(tt.shapes), got, tt.want)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	rect := Rectangle{W: 2, H: 3}
	bigRect := Rectangle{W: 4, H: 5}
	circle := Circle{R: 2}

	tests := []struct {
		name   string
		shapes []Shape
		want   Shape
		wantOK bool
	}{
		{"empty", nil, nil, false},
		{"single shape", []Shape{rect}, rect, true},
		{"circle wins", []Shape{rect, circle}, circle, true},
		{"rectangle wins", []Shape{circle, bigRect}, bigRect, true},
		{"tie keeps the first", []Shape{rect, Rectangle{W: 3, H: 2}}, rect, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := maxArea(tt.shapes)
			if ok != tt.wantOK {
				t.Fatalf("maxArea(%d shapes) ok = %v, want %v", len(tt.shapes), ok, tt.wantOK)
			}
			if !tt.wantOK {
				if got != nil {
					t.Errorf("maxArea(empty) = %v, want nil", got)
				}
				return
			}
			if got != tt.want {
				t.Errorf("maxArea(%d shapes) = %v (area %v), want %v (area %v)",
					len(tt.shapes), got, got.Area(), tt.want, tt.want.Area())
			}
		})
	}
}

func TestImplicitSatisfaction(t *testing.T) {
	shapes := []Shape{Rectangle{W: 3, H: 4}, Circle{R: 1}}
	if got, want := shapes[0].Name(), "rectangle"; got != want {
		t.Errorf("Rectangle.Name() = %q, want %q", got, want)
	}
	if got, want := shapes[1].Name(), "circle"; got != want {
		t.Errorf("Circle.Name() = %q, want %q", got, want)
	}
	if got, want := shapes[0].Area(), 12.0; got != want {
		t.Errorf("Rectangle{3,4}.Area() = %v, want %v", got, want)
	}
}
