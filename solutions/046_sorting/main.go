// 046: Sorting. sort.Ints and sort.Strings order slices in place, so copy the
// input first when the caller still needs the original order.
//
// Scenario: a report needs ranked numbers and alphabetical labels while the
// raw data stays available for an audit trail.
//
// Rules:
//  1. sortedInts returns a new ascending slice; nil or empty becomes a
//     non-nil empty slice.
//  2. sortedStrings returns a new ascending (case-sensitive) slice; nil or
//     empty becomes a non-nil empty slice.
//  3. The input slice must never be modified or aliased by the result.
//  4. Duplicates and negative values are kept.
package main

import (
	"fmt"
	"sort"
)

// sortedInts returns a sorted copy of in, leaving in untouched.
func sortedInts(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	sort.Ints(out)
	return out
}

// sortedStrings returns a sorted copy of in, leaving in untouched.
func sortedStrings(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	sort.Strings(out)
	return out
}

func main() {
	fmt.Println(sortedInts([]int{3, 1, 2}))
	fmt.Println(sortedStrings([]string{"pear", "apple", "fig"}))
}
