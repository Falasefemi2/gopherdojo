// 040: Worker Pools. A fixed set of goroutines pulls jobs from one channel and
// publishes results to another; closing jobs tells every worker to stop.
//
// Scenario: an invoice batch is squared up by a pool of clerks, then the totals
// are added together.
//
// Rules:
//  1. poolSquare returns the sum of each job squared.
//  2. It uses exactly `workers` goroutines; workers <= 0 is treated as 1.
//  3. The producer closes jobs after the last job; workers exit when jobs closes.
//  4. results is closed after every worker finishes, so the caller can range it.
package main

import (
	"fmt"
	"sync"
)

// poolSquare processes jobs with a fixed worker pool and returns the sum of
// the squares of the jobs.
func poolSquare(jobs []int, workers int) int {
	if workers <= 0 {
		workers = 1
	}

	jobsCh := make(chan int)
	resultsCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobsCh {
				resultsCh <- job * job
			}
		}()
	}

	go func() {
		for _, job := range jobs {
			jobsCh <- job
		}
		close(jobsCh)
	}()

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	total := 0
	for result := range resultsCh {
		total += result
	}
	return total
}

func main() {
	fmt.Println(poolSquare([]int{1, 2, 3, 4}, 2))
}
