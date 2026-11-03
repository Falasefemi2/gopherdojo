// 042: Rate Limiting. A ticker paces work: wait for one tick before each item
// so bursts are smoothed out.
//
// Scenario: a bulk API client is allowed only one request per millisecond, so
// it takes a ticket from the pacer before every call.
//
// Rules:
//  1. rateLimited processes items one at a time, waiting for a tick from a
//     ticker with period d before each item.
//  2. The returned slice holds the same items in the same order.
//  3. Empty input returns an empty (non-nil) slice.
//  4. The ticker is stopped before returning; tests use d = 1ms.
//
// TODO: implement rateLimited. Keep main() as the demo harness.
package main

import (
	"fmt"
	"time"
)

// rateLimited returns items after pacing them through a ticker.
func rateLimited(items []int, d time.Duration) []int {
	return nil
}

func main() {
	fmt.Println(rateLimited([]int{1, 2, 3}, time.Millisecond))
}
