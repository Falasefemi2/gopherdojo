// 030: Channel Buffering. A buffered channel accepts up to cap(ch) sends
// without a receiver being ready.
//
// Scenario: a coffee counter queues up to n orders before the barista has to
// make the next customer wait.
//
// Rules:
//  1. bufferStats creates a channel of the given capacity, sends fill values,
//     and returns (len, cap).
//  2. make(chan T, 0) is unbuffered, so its capacity is 0.
//  3. bufferedRoundTrip sends every value without a sender goroutine, using a
//     capacity of exactly len(vals), then closes and drains it in order.
//  4. Empty input returns an empty (non-nil) slice.
package main

import "fmt"

// bufferStats returns the length and capacity of a channel of size capacity
// after fill values have been buffered.
func bufferStats(size, fill int) (int, int) {
	ch := make(chan int, size)
	for i := 0; i < fill; i++ {
		ch <- i
	}
	return len(ch), cap(ch)
}

// bufferedRoundTrip sends vals into a channel buffered for exactly len(vals),
// closes it, and drains it back into a slice.
func bufferedRoundTrip(vals []int) []int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)

	out := make([]int, 0, len(vals))
	for v := range ch {
		out = append(out, v)
	}
	return out
}

func main() {
	fmt.Println(bufferStats(3, 2))
	fmt.Println(bufferedRoundTrip([]int{1, 2, 3}))
}
