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
package main

import (
	"fmt"
	"sync"
)

func SumConcurrent(values []int, workers int) int {
	if len(values) == 0 {
		return 0
	}
	if workers < 1 {
		workers = 1
	}
	if workers > len(values) {
		workers = len(values)
	}

	chunk := (len(values) + workers - 1) / workers
	results := make(chan int, workers)
	var wg sync.WaitGroup

	for start := 0; start < len(values); start += chunk {
		end := start + chunk
		if end > len(values) {
			end = len(values)
		}
		part := values[start:end]

		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := 0
			for _, v := range part {
				sum += v
			}
			results <- sum
		}()
	}

	wg.Wait()
	close(results)

	total := 0
	for sum := range results {
		total += sum
	}
	return total
}

func MapConcurrent(values []int, f func(int) int) []int {
	out := make([]int, len(values))
	var wg sync.WaitGroup

	for i, v := range values {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out[i] = f(v)
		}()
	}

	wg.Wait()
	return out
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("sum:", SumConcurrent(values, 4))
	fmt.Println("squares:", MapConcurrent(values, func(n int) int { return n * n }))
}
