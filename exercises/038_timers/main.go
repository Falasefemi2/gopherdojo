// 038: Timers. time.NewTimer fires once after a duration and delivers a single
// time.Time on its C channel. Stop it when you stop waiting.
//
// Scenario: a payment gateway gets one short window to answer before the
// checkout gives up.
//
// Rules:
//  1. newTimer returns a running timer that fires after d (tests use 5-20ms).
//  2. resetTimer restarts the same timer for a new duration and returns it.
//  3. Callers own the timer and must Stop it when they are done with it.
//  4. A stopped timer must not deliver a stale tick to a waiting select.
//
// TODO: implement newTimer and resetTimer. Keep main() as the demo harness.
package main

import (
	"fmt"
	"time"
)

// newTimer starts a timer that fires after d.
func newTimer(d time.Duration) *time.Timer {
	return time.NewTimer(time.Hour)
}

// resetTimer restarts t with a new duration and returns it.
func resetTimer(t *time.Timer, d time.Duration) *time.Timer {
	return t
}

func main() {
	timer := newTimer(10 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("timer fired")
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timer did not fire yet")
	}
}
