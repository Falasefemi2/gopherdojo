// 008: Arrays. An array [N]T has a fixed length that is part of its type, and
// assigning or passing one copies every element.
//
// Scenario: a five-slot LED panel is summed, scaled, and rotated. The caller's
// panel must never change unless the function says so.
//
// Rules:
//  1. sum5 returns the sum of the five elements.
//  2. double5 returns a new array with every element doubled; the input is a
//     copy and must be left untouched.
//  3. rotateRight moves every element one slot to the right and wraps the last
//     element to the front.
package main

import "fmt"

// sum5 returns the total of a.
func sum5(a [5]int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// double5 returns a copy of a with every element doubled.
func double5(a [5]int) [5]int {
	for i := range a {
		a[i] *= 2
	}
	return a
}

// rotateRight returns a rotated one slot to the right.
func rotateRight(a [5]int) [5]int {
	var out [5]int
	out[0] = a[len(a)-1]
	for i := 1; i < len(a); i++ {
		out[i] = a[i-1]
	}
	return out
}

func main() {
	panel := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum5(panel))
	fmt.Println(double5(panel))
	fmt.Println(rotateRight(panel))
}
