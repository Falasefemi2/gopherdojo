// 024: Generics. Type parameters plus a constraint interface let one
// function work over several numeric types without losing type safety.
//
// Scenario: a reporting service sums and scans both integer counts and
// float64 revenue, then maps either into a display type.
//
// Rules:
//  1. Number permits int, int64 and float64 (including named types with
//     those underlying types).
//  2. Sum of an empty slice is the zero value of T.
//  3. Max returns the largest value and true, or the zero value and false
//     for an empty slice.
//  4. Map returns a new slice of the same length in input order; an empty
//     input produces an empty, non-nil slice.
package main

import "fmt"

// Number is the numeric constraint used by the helpers.
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

func Max[T Number](values []T) (T, bool) {
	if len(values) == 0 {
		var zero T
		return zero, false
	}
	largest := values[0]
	for _, v := range values[1:] {
		if v > largest {
			largest = v
		}
	}
	return largest, true
}

func Map[T, U any](values []T, f func(T) U) []U {
	out := make([]U, len(values))
	for i, v := range values {
		out[i] = f(v)
	}
	return out
}

func main() {
	fmt.Println("sum:", Sum([]int{1, 2, 3, 4}))
	if largest, ok := Max([]float64{1.5, 4.25, -2}); ok {
		fmt.Println("max:", largest)
	}
	fmt.Println("squares:", Map([]int{1, 2, 3}, func(n int) int { return n * n }))
}
