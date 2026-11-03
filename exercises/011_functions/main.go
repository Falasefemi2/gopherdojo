// 011: Functions. Functions take parameters, return values, and can be passed
// around like any other value; one function may call another.
//
// Scenario: a checkout computes a subtotal, adds tax by calling the subtotal
// helper, and a pipeline applies a transformation to a value.
//
// Rules:
//  1. subtotal returns priceCents * qty, but any non-positive qty or negative
//     price yields 0.
//  2. totalWithTax calls subtotal and adds floor(subtotal/10) as 10% tax.
//  3. apply calls f exactly once with x and returns its result.
//
// TODO: implement the functions below. Keep main() as the demo harness.
package main

import "fmt"

// subtotal returns priceCents * qty.
func subtotal(priceCents, qty int) int {
	return 0
}

// totalWithTax returns the subtotal plus 10% tax.
func totalWithTax(priceCents, qty int) int {
	return 0
}

// apply returns f(x).
func apply(f func(int) int, x int) int {
	return 0
}

func main() {
	fmt.Println(subtotal(250, 3))
	fmt.Println(totalWithTax(250, 3))
	fmt.Println(apply(func(n int) int { return n * n }, 4))
}
