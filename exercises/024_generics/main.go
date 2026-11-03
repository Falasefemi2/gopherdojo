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
//
// TODO: implement Sum, Max and Map with real type parameters.
package main

import "fmt"

// Number is the numeric constraint used by the helpers.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum adds every value, returning T's zero value for an empty slice.
func Sum[T Number](values []T) T {
	var zero T
	return zero
}

// Max returns the largest value, or false when values is empty.
func Max[T Number](values []T) (T, bool) {
	var zero T
	return zero, false
}

// Map applies f to every value and keeps the input order.
func Map[T, U any](values []T, f func(T) U) []U {
	return nil
}

func main() {
	fmt.Println("sum:", Sum([]int{1, 2, 3, 4}))
	if largest, ok := Max([]float64{1.5, 4.25, -2}); ok {
		fmt.Println("max:", largest)
	}
	fmt.Println("squares:", Map([]int{1, 2, 3}, func(n int) int { return n * n }))
}
