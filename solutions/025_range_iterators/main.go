// 025: Range over Iterators. A function returning iter.Seq[T] can be used
// directly in a for-range loop; yield reports whether the caller wants more.
//
// Scenario: a scheduler streams countdown ticks and Fibonacci slots without
// allocating a slice up front, and callers may stop early.
//
// Rules:
//  1. Countdown(n) yields n, n-1, ..., 1; n <= 0 yields nothing.
//  2. Fibonacci(n) yields the first n values 0, 1, 1, 2, 3, ...; n <= 0
//     yields nothing.
//  3. Every sequence must honour yield returning false and stop immediately.
//  4. Sequences must not require a slice of all values in memory.
package main

import (
	"fmt"
	"iter"
)

func Countdown(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for v := n; v > 0; v-- {
			if !yield(v) {
				return
			}
		}
	}
}

func Fibonacci(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	var values []int
	for v := range Countdown(5) {
		values = append(values, v)
	}
	fmt.Println("countdown:", values)

	count := 0
	for v := range Fibonacci(8) {
		if count == 5 {
			break
		}
		count++
		fmt.Print(v, " ")
	}
	fmt.Println()
}
