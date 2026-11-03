// 039: Tickers. time.NewTicker repeats on a fixed period until Stop is called;
// if the receiver falls behind, only the latest tick is kept.
//
// Scenario: a metrics agent samples a counter every millisecond for a short
// run, then shuts the sampler down.
//
// Rules:
//  1. collectTicks starts a ticker with period d and collects exactly n ticks.
//  2. The ticker is stopped before collectTicks returns.
//  3. n <= 0 returns an empty (non-nil) slice and never starts a ticker.
//  4. The returned timestamps are in non-decreasing order.
//
// TODO: implement collectTicks. Keep main() as the demo harness.
package main

import (
	"fmt"
	"time"
)

// collectTicks returns n tick timestamps from a ticker with period d.
func collectTicks(n int, d time.Duration) []time.Time {
	return nil
}

func main() {
	ticks := collectTicks(3, time.Millisecond)
	for i, ts := range ticks {
		fmt.Println(i, ts.Format("15:04:05.000"))
	}
}
