// 009: Slices. A slice is a window onto a backing array: append may grow it,
// copy duplicates it, and two slices can alias the same elements.
//
// Scenario: an audio playlist is extended, reversed in place for shuffle, and
// filtered down to its even track numbers.
//
// Rules:
//  1. appendAll returns a new slice holding s followed by extra. It must not
//     share a backing array with s, and it never returns nil.
//  2. reverseInPlace reverses the caller's slice in place, sharing its backing
//     array, and returns nothing. Nil and empty slices are no-ops.
//  3. evens returns a new slice with the even values in their original order,
//     and a non-nil empty slice when there are none.
package main

import "fmt"

// appendAll returns s followed by extra in a fresh slice.
func appendAll(s, extra []int) []int {
	out := make([]int, 0, len(s)+len(extra))
	out = append(out, s...)
	out = append(out, extra...)
	return out
}

// reverseInPlace reverses s in place.
func reverseInPlace(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// evens returns the even values of s in order.
func evens(s []int) []int {
	out := make([]int, 0, len(s))
	for _, v := range s {
		if v%2 == 0 {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	playlist := []int{1, 2, 3, 4}
	fmt.Println(appendAll(playlist, []int{5, 6}))
	fmt.Println(evens(playlist))
	reverseInPlace(playlist)
	fmt.Println(playlist)
}
