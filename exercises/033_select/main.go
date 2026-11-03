// 033: Select. select waits on several channel operations and runs the one
// that is ready; a case whose channel is never ready is simply skipped.
//
// Scenario: an alert desk watches two feeds, but only one of them is wired up
// today.
//
// Rules:
//  1. selectReady must use select over both ready and idle.
//  2. It returns the value from ready, reporting ok == true.
//  3. idle is never written by selectReady, so it must never be chosen.
//  4. If ready already holds a value, the call returns it without extra delay.
//
// TODO: implement selectReady. Keep main() as the demo harness.
package main

import "fmt"

// selectReady waits on ready and idle and returns the value from ready.
// idle is never sent on, so it must never win the select.
func selectReady(ready, idle <-chan int) (int, bool) {
	return 0, false
}

func main() {
	ready := make(chan int, 1)
	ready <- 99
	fmt.Println(selectReady(ready, make(chan int)))
}
