// 028: Goroutines. `go f()` starts a concurrent call; a WaitGroup or a done
// channel tells the parent when all workers have finished.
//
// Scenario: a batch reporter splits a large slice across a fixed number of
// worker goroutines, then assembles deterministic results.
//
// Rules:
//  1. SumConcurrent returns the exact total for any workers value; workers
//     below 1 is treated as 1 and above len(values) is capped.
//  2. SumConcurrent must wait for every worker before returning.
//  3. MapConcurrent applies f concurrently but returns results in input
//     order.
//  4. MapConcurrent on an empty input returns an empty, non-nil slice.
//
// TODO: implement SumConcurrent and MapConcurrent using goroutines. Do not
// use sleeps. Aggregate through a channel or synchronize shared writes.
package main

import "fmt"

// SumConcurrent sums values using at most workers goroutines.
func SumConcurrent(values []int, workers int) int {
	return 0
}

// MapConcurrent applies f to every value concurrently, preserving order.
func MapConcurrent(values []int, f func(int) int) []int {
	return nil
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("sum:", SumConcurrent(values, 4))
	fmt.Println("squares:", MapConcurrent(values, func(n int) int { return n * n }))
}
