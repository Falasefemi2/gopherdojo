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
package main

import "fmt"

// squareChan returns a channel of the squares 1..n, closed after the last send.
func squareChan(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

// sumChannel consumes ch with range and returns the sum of its values.
func sumChannel(ch <-chan int) int {
	total := 0
	for v := range ch {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sumChannel(squareChan(4)))
}
