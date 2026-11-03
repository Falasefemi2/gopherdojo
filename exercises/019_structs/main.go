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
//
// TODO: implement the three functions below.
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

// newPoint builds a Point with the given coordinates.
func newPoint(x, y int) Point {
	return Point{}
}

// translate returns a copy of p moved by dx and dy.
func translate(p Point, dx, dy int) Point {
	return Point{}
}

// area returns the rectangle area, or 0 for an inverted or empty rectangle.
func area(r Rect) int {
	return 0
}

func main() {
	p := newPoint(1, 2)
	q := translate(p, 3, 4)
	r := Rect{Min: p, Max: q}
	fmt.Println("point:", p)
	fmt.Println("moved:", q)
	fmt.Println("area:", area(r))
}
