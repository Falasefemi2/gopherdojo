// 014: Closures. A function literal can capture variables from its enclosing
// scope, and each call to the enclosing function creates fresh state.
//
// Scenario: a ticket desk issues counters, an accumulator tracks a running
// total through two functions, and a multiplier factory bakes in a factor.
//
// Rules:
//  1. newCounter returns a function that adds 1 and returns the new value, so
//     the first call on newCounter(0) returns 1. Two counters never share state.
//  2. newAccumulator returns add and value functions sharing one total. add
//     adds its argument and returns the new total; value returns the current
//     total without changing it. The total starts at 0.
//  3. newMultiplier captures its argument by value: later changes to the
//     caller's variable must not affect the returned function.
//
// TODO: implement the functions below. Keep main() as the demo harness.
package main

import "fmt"

// newCounter returns a counter starting just after start.
func newCounter(start int) func() int {
	return func() int { return 0 }
}

// newAccumulator returns an add function and a read-only value function.
func newAccumulator() (add func(int) int, value func() int) {
	return func(int) int { return 0 }, func() int { return 0 }
}

// newMultiplier returns a function multiplying by factor.
func newMultiplier(factor int) func(int) int {
	return func(int) int { return 0 }
}

func main() {
	counter := newCounter(0)
	fmt.Println(counter(), counter())
	add, value := newAccumulator()
	fmt.Println(add(5), value())
	fmt.Println(newMultiplier(3)(4))
}
