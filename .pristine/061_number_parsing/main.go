// 061: Number Parsing. Text from files, flags and forms is just bytes; strconv
// turns it into numbers and reports an error instead of guessing.
//
// Scenario: a billing importer reads quantity, price and a "gift" flag from a
// CSV row and must reject malformed rows with a clear error.
//
// Rules:
//  1. parseInt accepts base-10 ints, including a leading + or -.
//  2. parseFloat accepts decimal and exponent notation as float64.
//  3. parseBool accepts exactly "true" and "false" (lowercase only).
//  4. On any error the function returns the zero value and a non-nil error.
//
// TODO: implement parseInt, parseFloat and parseBool below. Keep main().
package main

import (
	"fmt"
)

// parseInt reads a base-10 integer.
func parseInt(s string) (int, error) {
	// TODO
	return 0, nil
}

// parseFloat reads a decimal floating-point number.
func parseFloat(s string) (float64, error) {
	// TODO
	return 0, nil
}

// parseBool reads exactly "true" or "false".
func parseBool(s string) (bool, error) {
	// TODO
	return false, nil
}

func main() {
	n, err := parseInt("42")
	fmt.Println("int:", n, "err:", err)

	f, err := parseFloat("19.95")
	fmt.Println("float:", f, "err:", err)

	b, err := parseBool("true")
	fmt.Println("bool:", b, "err:", err)
}
