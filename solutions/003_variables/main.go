// 003: Variables. Variables are declared with var or created with :=, and
// several can be assigned in one statement.
//
// Scenario: a seating chart is reshuffled between courses. The host swaps two
// seats, rotates three seats, and prints the resulting order.
//
// Rules:
//  1. swap returns its two arguments in reverse order using a single
//     multiple-assignment statement, not a temporary variable.
//  2. rotate3 moves each value one seat left: (a, b, c) becomes (b, c, a).
//  3. swapReport formats the change as "before: <a>, <b>; after: <b>, <a>".
//  4. Zero and negative integers must work exactly like positive ones.
package main

import "fmt"

// swap returns b, a.
func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// rotate3 returns b, c, a.
func rotate3(a, b, c int) (int, int, int) {
	a, b, c = b, c, a
	return a, b, c
}

// swapReport describes the swap in one line.
func swapReport(a, b int) string {
	afterA, afterB := swap(a, b)
	return fmt.Sprintf("before: %d, %d; after: %d, %d", a, b, afterA, afterB)
}

func main() {
	x, y := swap(3, 7)
	fmt.Println(x, y)
	fmt.Println(rotate3(1, 2, 3))
	fmt.Println(swapReport(3, 7))
}
