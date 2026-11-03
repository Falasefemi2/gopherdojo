// 047: Sorting by Functions. sort.Slice accepts a less function, which lets
// you rank values by any rule you can express.
//
// Scenario: a reading list is ordered shortest-first, and equal lengths are
// broken alphabetically so the output is predictable.
//
// Rules:
//  1. Sort by length ascending using sort.Slice.
//  2. Break length ties by name ascending, so the order is total and stable
//     for identical entries.
//  3. Return a new slice and never modify or alias the input.
//  4. nil or empty input returns a non-nil empty slice.
package main

import (
	"fmt"
	"sort"
)

// byLengthThenName returns a copy of words ordered by length ascending, then
// alphabetically for equal lengths.
func byLengthThenName(words []string) []string {
	out := make([]string, len(words))
	copy(out, words)

	sort.Slice(out, func(i, j int) bool {
		if len(out[i]) != len(out[j]) {
			return len(out[i]) < len(out[j])
		}
		return out[i] < out[j]
	})
	return out
}

func main() {
	fmt.Println(byLengthThenName([]string{"pear", "fig", "apple", "kiwi"}))
}
