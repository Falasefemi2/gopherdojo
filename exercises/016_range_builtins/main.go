// 016: Range over Built-in Types. range walks a slice by index/value,
// a map in unspecified order, and a string by rune.
//
// Scenario: an analytics job scans sample slices and a map of page tags
// before printing a stable, sorted report.
//
// Rules:
//  1. rangeSum adds every element of nums; nil and empty slices sum to 0.
//  2. firstIndex returns the first matching index, or -1 when absent.
//  3. vowelCount counts ASCII vowels a, e, i, o, u in s, ignoring case,
//     by ranging over the string's runes (so "café" counts one).
//  4. sortedKeys returns the map keys in ascending order; an empty map
//     returns an empty, non-nil slice.
//
// TODO: implement the four functions using range. Do not index slices by
// hand where range expresses the loop directly.
package main

import "fmt"

// rangeSum returns the sum of nums.
func rangeSum(nums []int) int {
	return 0
}

// firstIndex returns the index of target in words, or -1.
func firstIndex(words []string, target string) int {
	return 0
}

// vowelCount counts ASCII vowels in s, ignoring case.
func vowelCount(s string) int {
	return 0
}

// sortedKeys returns the keys of m sorted ascending.
func sortedKeys(m map[string]int) []string {
	return nil
}

func main() {
	fmt.Println("sum:", rangeSum([]int{1, 2, 3, 4}))
	fmt.Println("index:", firstIndex([]string{"go", "range", "runes"}, "range"))
	fmt.Println("vowels:", vowelCount("Hello, Gopher!"))
	fmt.Println("keys:", sortedKeys(map[string]int{"b": 2, "a": 1, "c": 3}))
}
