// 010: Maps. A map stores key/value pairs with fast lookup. The comma-ok form
// tells a missing key apart from a stored zero value.
//
// Scenario: a log analyser counts words in a message, reads a counter with a
// default, and drops noisy keys from a shared tally.
//
// Rules:
//  1. wordCount lowercases the text, splits it on whitespace, and counts each
//     word. Empty or whitespace-only text yields a non-nil empty map.
//  2. getOr returns the stored value, or fallback when the key is absent. A
//     stored 0 is a present value, not a miss.
//  3. removeKey deletes key in place; removing a missing key is a no-op.
package main

import (
	"fmt"
	"strings"
)

// wordCount counts words in text.
func wordCount(text string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(strings.ToLower(text)) {
		counts[word]++
	}
	return counts
}

// getOr returns m[key] or fallback.
func getOr(m map[string]int, key string, fallback int) int {
	if v, ok := m[key]; ok {
		return v
	}
	return fallback
}

// removeKey deletes key from m.
func removeKey(m map[string]int, key string) {
	delete(m, key)
}

func main() {
	counts := wordCount("the cat and the hat")
	fmt.Println(counts)
	fmt.Println(getOr(counts, "the", 0), getOr(counts, "dog", -1))
	removeKey(counts, "cat")
	fmt.Println(counts)
}
