// 035: Non-Blocking Channel Operations. select with a default case never
// waits: when no other case is ready, default runs immediately.
//
// Scenario: a telemetry agent drops samples instead of stalling its caller.
//
// Rules:
//  1. trySend sends only when the channel can accept the value right now.
//  2. tryReceive reports (value, true) only when a value is buffered or a
//     sender is already waiting.
//  3. A full channel, an empty channel, and a channel with no partner all take
//     the default path immediately.
//  4. A closed channel yields (0, false) on receive.
package main

import "fmt"

// trySend attempts a non-blocking send of v on ch.
func trySend(ch chan<- int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

// tryReceive attempts a non-blocking receive from ch.
func tryReceive(ch <-chan int) (int, bool) {
	select {
	case v, ok := <-ch:
		return v, ok
	default:
		return 0, false
	}
}

func main() {
	ch := make(chan int, 1)
	fmt.Println(trySend(ch, 1))
	fmt.Println(trySend(ch, 2))
	fmt.Println(tryReceive(ch))
	fmt.Println(tryReceive(ch))
}
