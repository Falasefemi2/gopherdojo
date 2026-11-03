// 045: Stateful Goroutines. Instead of sharing memory, one owner goroutine
// holds the state and other goroutines ask it to change or read that state.
//
// Scenario: a bank ledger has a single writer. Callers send requests over a
// channel and wait for the reply on their own channel.
//
// Rules:
//  1. Exactly one owner goroutine is allowed to mutate the running total.
//  2. Each request carries a delta and a reply channel for the new total.
//  3. Requests are applied in input order; result i is the running total
//     after deltas[0..i].
//  4. An empty input returns a non-nil empty slice.
package main

import "fmt"

type tallyRequest struct {
	delta int
	reply chan int
}

// tally applies each delta in order through a single owner goroutine and
// returns the running total after every request.
func tally(deltas []int) []int {
	requests := make(chan tallyRequest)
	done := make(chan struct{})

	go func() {
		defer close(done)
		total := 0
		for req := range requests {
			total += req.delta
			req.reply <- total
		}
	}()

	results := make([]int, 0, len(deltas))
	for _, delta := range deltas {
		reply := make(chan int, 1)
		requests <- tallyRequest{delta: delta, reply: reply}
		results = append(results, <-reply)
	}

	close(requests)
	<-done
	return results
}

func main() {
	fmt.Println(tally([]int{10, -3, 5}))
}
