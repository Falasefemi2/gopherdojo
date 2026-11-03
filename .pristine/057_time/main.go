// 057: Time. time.Time values are immutable: Add and Sub return new values,
// and Before/After compare two instants without changing either one.
//
// Scenario: a deploy bot schedules maintenance windows in UTC and needs to
// shift a window, measure its length, and tell which window comes first.
//
// Rules:
//  1. plus returns t shifted by d and must leave t unchanged.
//  2. gap returns exactly to.Sub(from); the result may be negative.
//  3. earlier reports whether a is strictly before b, so equal times are false.
//  4. A shifted time keeps the location of its input (UTC in the tests).
//
// TODO: implement plus, gap and earlier below. Keep main() as the demo harness.
package main

import (
	"fmt"
	"time"
)

// plus returns t shifted by d.
func plus(t time.Time, d time.Duration) time.Time {
	// TODO
	return time.Time{}
}

// gap returns how far to is from from.
func gap(from, to time.Time) time.Duration {
	// TODO
	return 0
}

// earlier reports whether a happens strictly before b.
func earlier(a, b time.Time) bool {
	// TODO
	return false
}

func main() {
	start := time.Date(2026, time.March, 2, 9, 0, 0, 0, time.UTC)
	end := plus(start, 90*time.Minute)
	fmt.Printf("window: %s -> %s (%v)\n", start.Format("15:04"), end.Format("15:04"), gap(start, end))
	fmt.Println("ends before it starts:", earlier(end, start))
}
