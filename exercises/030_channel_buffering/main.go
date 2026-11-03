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
//
// TODO: implement bufferStats and bufferedRoundTrip. Keep main() as the demo.
package main

import "fmt"

// bufferStats returns the length and capacity of a channel of size capacity
// after fill values have been buffered.
func bufferStats(size, fill int) (int, int) {
	return 0, 0
}

// bufferedRoundTrip sends vals into a channel buffered for exactly len(vals),
// closes it, and drains it back into a slice.
func bufferedRoundTrip(vals []int) []int {
	return nil
}

func main() {
	fmt.Println(bufferStats(3, 2))
	fmt.Println(bufferedRoundTrip([]int{1, 2, 3}))
}
