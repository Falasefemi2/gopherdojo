// 012: Multiple Return Values. A function can return several values at once,
// which Go uses for results plus an ok flag.
//
// Scenario: a till splits a bill, a scoreboard reports the low and high score,
// and a name parser separates first and last names.
//
// Rules:
//  1. divmod returns quotient, remainder and true. Division by zero returns
//     (0, 0, false). Quotients truncate toward zero, like Go's / operator.
//  2. minmax returns the smallest and largest of three integers.
//  3. splitName splits at the last space after trimming surrounding spaces,
//     returning ("", "") for empty input and (name, "") for a single word.
//
// TODO: implement the functions below. Keep main() as the demo harness.
package main

import "fmt"

// divmod returns a/b, a%b and whether the division was valid.
func divmod(a, b int) (int, int, bool) {
	return 0, 0, false
}

// minmax returns the smallest and largest of a, b and c.
func minmax(a, b, c int) (int, int) {
	return 0, 0
}

// splitName separates a full name at its last space.
func splitName(full string) (string, string) {
	return "", ""
}

func main() {
	q, r, ok := divmod(17, 5)
	fmt.Println(q, r, ok)
	lo, hi := minmax(3, 1, 2)
	fmt.Println(lo, hi)
	first, last := splitName("Ada Lovelace")
	fmt.Println(first, last)
}
