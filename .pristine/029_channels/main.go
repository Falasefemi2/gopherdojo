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
//
// TODO: implement roundTrip and relayAll below. Keep main() as the demo harness.
package main

import "fmt"

// roundTrip sends v across an unbuffered channel from a goroutine and returns
// the value it receives.
func roundTrip(v int) int {
	return 0
}

// relayAll sends each value of vals in order over an unbuffered channel from a
// goroutine and receives each one in the calling goroutine.
func relayAll(vals []int) []int {
	return nil
}

func main() {
	fmt.Println(roundTrip(7))
	fmt.Println(relayAll([]int{1, 2, 3}))
}
