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
package main

import "fmt"

// sumTo returns 1 + 2 + ... + n.
func sumTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// countdown returns n down to 1.
func countdown(n int) []int {
	out := []int{}
	for n > 0 {
		out = append(out, n)
		n--
	}
	return out
}

// nextPowerOfTwo returns the smallest power of two >= n.
func nextPowerOfTwo(n int) int {
	p := 1
	for {
		if p >= n {
			return p
		}
		p *= 2
	}
}

func main() {
	fmt.Println(sumTo(5))
	fmt.Println(countdown(3))
	fmt.Println(nextPowerOfTwo(17))
}
