// 021: Interfaces. A type satisfies an interface implicitly by implementing
// its method set; callers can then treat different concrete types alike.
//
// Scenario: a quoting service sums a mixed basket of shapes and needs to
// report the largest one without switching on concrete types.
//
// Rules:
//  1. Shape is satisfied by any type with Name() string and Area() float64.
//  2. totalArea sums every shape's area; a nil or empty slice totals 0.
//  3. maxArea returns the shape with the greatest area and true, or nil and
//     false for an empty slice. On a tie the first shape wins.
//  4. Rectangle area is W*H and circle area is pi*R*R.
package main

import (
	"fmt"
	"math"
)

// Shape is anything that can report its name and area.
type Shape interface {
	Name() string
	Area() float64
}

// Rectangle is a shape with width and height.
type Rectangle struct {
	W float64
	H float64
}

// Name returns "rectangle".
func (r Rectangle) Name() string { return "rectangle" }

// Area returns W*H.
func (r Rectangle) Area() float64 { return r.W * r.H }

// Circle is a shape with radius.
type Circle struct {
	R float64
}

// Name returns "circle".
func (c Circle) Name() string { return "circle" }

// Area returns pi*R*R.
func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func maxArea(shapes []Shape) (Shape, bool) {
	if len(shapes) == 0 {
		return nil, false
	}
	largest := shapes[0]
	for _, shape := range shapes[1:] {
		if shape.Area() > largest.Area() {
			largest = shape
		}
	}
	return largest, true
}

func main() {
	shapes := []Shape{Rectangle{W: 2, H: 3}, Circle{R: 2}}
	fmt.Println("total:", totalArea(shapes))
	if biggest, ok := maxArea(shapes); ok {
		fmt.Printf("largest: %s (%.2f)\n", biggest.Name(), biggest.Area())
	}
}
