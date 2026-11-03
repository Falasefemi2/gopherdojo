// 043: Atomic Counters. sync/atomic updates a shared counter in one
// indivisible step, so concurrent goroutines cannot lose an increment.
//
// Scenario: a metrics endpoint is hit by many background workers at once, and
// every hit must be counted exactly once.
//
// Rules:
//  1. Use an atomic.Int64 counter (sync/atomic), never a plain int.
//  2. Spawn exactly n goroutines, each performing one increment.
//  3. Join every goroutine with a sync.WaitGroup before reading the counter.
//  4. n may be zero; return 0 and still shut down cleanly.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// countHits starts n goroutines that each record one hit and returns the
// final counter value after all of them have finished.
func countHits(n int) int64 {
	var counter atomic.Int64
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()

	return counter.Load()
}

func main() {
	fmt.Printf("hits after 100 workers: %d\n", countHits(100))
}
