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
//
// TODO: implement the three functions using pointer parameters.
package main

import "fmt"

// increment adds one to the value pointed at by p.
func increment(p *int) {
}

// swap exchanges the values pointed at by a and b.
func swap(a, b *int) {
}

// setIfPositive stores v through p when v is positive; it reports success.
func setIfPositive(p *int, v int) bool {
	return false
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
