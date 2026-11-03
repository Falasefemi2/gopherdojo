// 017: Pointers. & takes an address, * reads or writes the value at that
// address, and a nil pointer must never be dereferenced.
//
// Scenario: a settings package updates counters in place while tolerating
// optional (nil) fields that callers did not provide.
//
// Rules:
//  1. increment adds one through p; a nil p is a silent no-op.
//  2. swap exchanges the two pointed-at values; if either side is nil,
//     nothing changes and no panic occurs.
//  3. setIfPositive stores v only when p is non-nil and v > 0, and reports
//     whether a store happened.
//  4. A successful call must mutate the caller's variable, not a copy.
package main

import "fmt"

func increment(p *int) {
	if p == nil {
		return
	}
	*p++
}

func swap(a, b *int) {
	if a == nil || b == nil {
		return
	}
	*a, *b = *b, *a
}

func setIfPositive(p *int, v int) bool {
	if p == nil || v <= 0 {
		return false
	}
	*p = v
	return true
}

func main() {
	n := 1
	increment(&n)
	fmt.Println("incremented:", n)

	a, b := 10, 20
	swap(&a, &b)
	fmt.Println("swapped:", a, b)

	target := 0
	fmt.Println("set:", setIfPositive(&target, 9), target)
}
