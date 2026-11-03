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
//
// TODO: implement poolSquare. Keep main() as the demo harness.
package main

import "fmt"

// poolSquare processes jobs with a fixed worker pool and returns the sum of
// the squares of the jobs.
func poolSquare(jobs []int, workers int) int {
	return 0
}

func main() {
	fmt.Println(poolSquare([]int{1, 2, 3, 4}, 2))
}
