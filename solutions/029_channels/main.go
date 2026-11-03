// 029: Channels. A channel is a typed conduit: a send blocks until a receiver
// is ready, and a receive blocks until a sender arrives.
//
// Scenario: a ticket desk hands requests to the only clerk on duty, one at a
// time, and waits for each ticket to be taken.
//
// Rules:
//  1. roundTrip must send v over an unbuffered channel from a goroutine and
//     return the value it receives.
//  2. relayAll must hand over every value one at a time, preserving order.
//  3. An empty input returns an empty (non-nil) slice.
//  4. No goroutine may stay blocked when a function returns.
package main

import "fmt"

// roundTrip sends v across an unbuffered channel from a goroutine and returns
// the value it receives.
func roundTrip(v int) int {
	ch := make(chan int)
	go func() { ch <- v }()
	return <-ch
}

// relayAll sends each value of vals in order over an unbuffered channel from a
// goroutine and receives each one in the calling goroutine.
func relayAll(vals []int) []int {
	out := make([]int, 0, len(vals))
	ch := make(chan int)
	go func() {
		for _, v := range vals {
			ch <- v
		}
	}()
	for range vals {
		out = append(out, <-ch)
	}
	return out
}

func main() {
	fmt.Println(roundTrip(7))
	fmt.Println(relayAll([]int{1, 2, 3}))
}
