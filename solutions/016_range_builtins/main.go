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
package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func firstIndex(words []string, target string) int {
	for i, word := range words {
		if word == target {
			return i
		}
	}
	return -1
}

func vowelCount(s string) int {
	count := 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	fmt.Println("sum:", rangeSum([]int{1, 2, 3, 4}))
	fmt.Println("index:", firstIndex([]string{"go", "range", "runes"}, "range"))
	fmt.Println("vowels:", vowelCount("Hello, Gopher!"))
	fmt.Println("keys:", sortedKeys(map[string]int{"b": 2, "a": 1, "c": 3}))
}
