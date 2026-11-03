// 060: Random Numbers. math/rand needs an explicit seed for reproducible
// streams: rand.New(rand.NewSource(seed)) gives a generator you own, and
// Intn(n) returns an int in [0, n).
//
// Scenario: a load test replays a "random" user journey, but CI must produce
// the exact same journey on every run so failures are reproducible.
//
// Rules:
//  1. rolls returns exactly n numbers, each in the half-open range [0, 100).
//  2. Two calls with the same seed must return identical sequences.
//  3. n <= 0 returns a non-nil empty slice, never nil.
//  4. Use a private generator: never the package-level rand functions.
//
// TODO: implement rolls below using rand.New(rand.NewSource(seed)). Keep main().
package main

import (
	"fmt"
)

// rolls returns n pseudo-random values in [0, 100) from the given seed.
func rolls(seed int64, n int) []int {
	// TODO
	return nil
}

func main() {
	fmt.Println("seed 42:", rolls(42, 5))
	fmt.Println("seed 42:", rolls(42, 5))
	fmt.Println("seed 7: ", rolls(7, 5))
}
