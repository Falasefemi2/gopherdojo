// 034: Timeouts. select plus time.After gives up on a channel that stays quiet.
//
// Scenario: the checkout calls a flaky pricing service but never waits longer
// than the SLA allows.
//
// Rules:
//  1. awaitValue returns (v, true) when a value arrives before d elapses.
//  2. It returns (0, false) when the timer fires first.
//  3. A closed channel reports (0, false) instead of a phantom zero value.
//  4. A received zero is still a success: trust the bool, not the value.
package main

import (
	"fmt"
	"time"
)

// awaitValue waits at most d for a value on ch.
func awaitValue(ch <-chan int, d time.Duration) (int, bool) {
	select {
	case v, ok := <-ch:
		if !ok {
			return 0, false
		}
		return v, true
	case <-time.After(d):
		return 0, false
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 5
	fmt.Println(awaitValue(ch, 20*time.Millisecond))

	v, ok := awaitValue(make(chan int), 10*time.Millisecond)
	fmt.Println(v, ok)
}
