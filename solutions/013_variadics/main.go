// 013: Variadic Functions. A final ...T parameter accepts any number of
// arguments, and a slice can be spread into it with slice....
//
// Scenario: a scoreboard totals an arbitrary number of rounds, picks the best
// score, and prints them joined by a separator.
//
// Rules:
//  1. sum adds every argument and returns 0 with no arguments; negatives count.
//  2. maxOf returns the largest argument and true, or (0, false) when empty.
//  3. joinInts renders each number in decimal, joined by sep; with no numbers
//     it returns "". Passing nums... must match passing the elements.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// sum adds all of nums.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// maxOf returns the largest of nums and whether any were given.
func maxOf(nums ...int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	best := nums[0]
	for _, n := range nums[1:] {
		if n > best {
			best = n
		}
	}
	return best, true
}

// joinInts joins nums in decimal with sep.
func joinInts(sep string, nums ...int) string {
	parts := make([]string, len(nums))
	for i, n := range nums {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, sep)
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(maxOf(4, 9, 2))
	fmt.Println(joinInts("-", 1, 2, 3))
}
