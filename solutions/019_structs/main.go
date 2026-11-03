// 019: Structs. A struct groups named fields; assigning one copies every
// field, so helpers that "move" a struct return a new value.
//
// Scenario: a layout editor models points and rectangles as values and must
// never mutate a shape the caller still owns.
//
// Rules:
//  1. newPoint builds a Point from its two coordinates.
//  2. translate returns a moved copy; the input Point is unchanged.
//  3. area of a rect is (Max.X-Min.X)*(Max.Y-Min.Y), and an inverted or
//     empty rectangle has area 0.
//  4. The zero Point is (0, 0).
package main

import "fmt"

// Point is a 2D coordinate.
type Point struct {
	X int
	Y int
}

// Rect is an axis-aligned rectangle.
type Rect struct {
	Min Point
	Max Point
}

func newPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func translate(p Point, dx, dy int) Point {
	return Point{X: p.X + dx, Y: p.Y + dy}
}

func area(r Rect) int {
	width := r.Max.X - r.Min.X
	height := r.Max.Y - r.Min.Y
	if width <= 0 || height <= 0 {
		return 0
	}
	return width * height
}

func main() {
	p := newPoint(1, 2)
	q := translate(p, 3, 4)
	r := Rect{Min: p, Max: q}
	fmt.Println("point:", p)
	fmt.Println("moved:", q)
	fmt.Println("area:", area(r))
}
