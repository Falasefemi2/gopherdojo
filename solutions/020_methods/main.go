// 020: Methods. A value receiver works on a copy; a pointer receiver can
// mutate the original value.
//
// Scenario: a metrics collector hands out snapshots while live counters keep
// changing, so incrementing must not require reassigning the struct.
//
// Rules:
//  1. Inc adds one through a pointer receiver and is visible to the caller.
//  2. Add adds delta (which may be negative) through a pointer receiver.
//  3. Value reports the current count and never mutates the receiver.
//  4. WithAdded uses a value receiver: it returns a new Counter and leaves
//     the original untouched.
package main

import "fmt"

// Counter is a running tally.
type Counter struct {
	n int
}

func (c *Counter) Inc() {
	c.n++
}

func (c *Counter) Add(delta int) {
	c.n += delta
}

func (c Counter) Value() int {
	return c.n
}

func (c Counter) WithAdded(delta int) Counter {
	return Counter{n: c.n + delta}
}

func main() {
	var c Counter
	c.Inc()
	c.Add(4)
	fmt.Println("count:", c.Value())
	fmt.Println("snapshot:", c.WithAdded(10).Value())
}
