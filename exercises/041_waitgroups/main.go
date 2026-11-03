// 041: WaitGroups. sync.WaitGroup counts live goroutines: Add before starting,
// Done when finishing, Wait to block until the count reaches zero.
//
// Scenario: ten auditors each stamp one ledger row before the close is signed.
//
// Rules:
//  1. concurrentTotal starts n goroutines; each adds step to a shared total and
//     signals Done. Wait for all of them, then return the exact total.
//  2. n <= 0 returns 0 without spawning goroutines.
//  3. writeIndexes writes i*10 into slot i from goroutine i, then waits.
//  4. Never call Add from inside the goroutine it guards; Add before `go`.
//
// TODO: implement concurrentTotal and writeIndexes. Keep main() as the demo.
package main

import "fmt"

// concurrentTotal adds step once per goroutine, n times in total.
func concurrentTotal(n, step int) int64 {
	return 0
}

// writeIndexes returns a slice where slot i holds i*10.
func writeIndexes(n int) []int {
	return nil
}

func main() {
	fmt.Println(concurrentTotal(100, 2))
	fmt.Println(writeIndexes(4))
}
