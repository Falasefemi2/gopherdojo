// 005: For. Go's only loop keyword is for, but it covers the classic counter
// loop, the while-style condition loop, and the infinite loop with break.
//
// Scenario: a fitness band counts steps up to a goal, plays a countdown, then
// finds the smallest power-of-two buffer that can hold the data.
//
// Rules:
//  1. sumTo adds 1 through n inclusive and returns 0 for n <= 0.
//  2. countdown returns n, n-1, ..., 1 and a non-nil empty slice for n <= 0.
//  3. nextPowerOfTwo returns the smallest power of two >= n, using an infinite
//     loop that breaks; it returns 1 for n <= 1.
//
// TODO: implement the functions below. Keep main() as the demo harness.
package main

import "fmt"

// sumTo returns 1 + 2 + ... + n.
func sumTo(n int) int {
	return 0
}

// countdown returns n down to 1.
func countdown(n int) []int {
	return nil
}

// nextPowerOfTwo returns the smallest power of two >= n.
func nextPowerOfTwo(n int) int {
	return 0
}

func main() {
	fmt.Println(sumTo(5))
	fmt.Println(countdown(3))
	fmt.Println(nextPowerOfTwo(17))
}
