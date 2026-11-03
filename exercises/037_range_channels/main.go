// 037: Range over Channels. `for v := range ch` receives until the channel is
// closed, then the loop ends.
//
// Scenario: a sensor stream emits a fixed run of readings and then shuts down.
//
// Rules:
//  1. squareChan returns a channel yielding 1² .. n² in order, closed after the
//     last value.
//  2. sumChannel ranges over ch until it is closed and returns the total.
//  3. n <= 0 yields an empty channel that is already closed.
//  4. The sending goroutine closes the channel; the receiver never does.
//
// TODO: implement squareChan and sumChannel. Keep main() as the demo harness.
package main

import "fmt"

// squareChan returns a channel of the squares 1..n, closed after the last send.
func squareChan(n int) <-chan int {
	return nil
}

// sumChannel consumes ch with range and returns the sum of its values.
func sumChannel(ch <-chan int) int {
	return 0
}

func main() {
	fmt.Println(sumChannel(squareChan(4)))
}
