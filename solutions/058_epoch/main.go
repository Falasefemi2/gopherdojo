// 058: Epoch. A time.Time converts to Unix seconds, milliseconds, or
// nanoseconds and back again; the numeric forms carry no location.
//
// Scenario: an audit log stores timestamps as epoch numbers, but reports need
// real UTC times. Convert both ways without losing precision.
//
// Rules:
//  1. unixSeconds, unixMillis and unixNanos return the epoch value of t.
//  2. fromUnixSeconds, fromUnixMillis and fromUnixNanos return UTC times.
//  3. Seconds drop sub-second precision; millis drop the last six digits.
//  4. A roundtrip through the matching pair must reproduce the same instant.
package main

import (
	"fmt"
	"time"
)

// unixSeconds returns t as seconds since the Unix epoch.
func unixSeconds(t time.Time) int64 {
	return t.Unix()
}

// fromUnixSeconds turns epoch seconds into a UTC time.
func fromUnixSeconds(sec int64) time.Time {
	return time.Unix(sec, 0).UTC()
}

// unixMillis returns t as milliseconds since the Unix epoch.
func unixMillis(t time.Time) int64 {
	return t.UnixMilli()
}

// fromUnixMillis turns epoch milliseconds into a UTC time.
func fromUnixMillis(ms int64) time.Time {
	return time.UnixMilli(ms).UTC()
}

// unixNanos returns t as nanoseconds since the Unix epoch.
func unixNanos(t time.Time) int64 {
	return t.UnixNano()
}

// fromUnixNanos turns epoch nanoseconds into a UTC time.
func fromUnixNanos(ns int64) time.Time {
	return time.Unix(0, ns).UTC()
}

func main() {
	base := time.Date(2024, time.February, 29, 12, 34, 56, 789000000, time.UTC)
	fmt.Println("seconds:", unixSeconds(base))
	fmt.Println("millis: ", unixMillis(base))
	fmt.Println("nanos:  ", unixNanos(base))
	fmt.Println("back:   ", fromUnixMillis(unixMillis(base)).Format(time.RFC3339Nano))
}
