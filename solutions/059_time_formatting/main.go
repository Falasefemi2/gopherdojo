// 059: Time Formatting / Parsing. Go formats times with a reference layout,
// not with strftime codes: "2006-01-02 15:04" means year-month-day hour:minute.
//
// Scenario: a scheduler reads "YYYY-MM-DD HH:MM" timestamps from a config file
// and prints them back to the operator in the same shape.
//
// Rules:
//  1. formatStamp renders t with the exact layout "2006-01-02 15:04".
//  2. Seconds and smaller units never appear in the formatted string.
//  3. parseStamp accepts only that same layout and returns a UTC time.
//  4. Bad input returns an error and the zero time, never a guessed time.
package main

import (
	"fmt"
	"time"
)

const stampLayout = "2006-01-02 15:04"

// formatStamp renders t as "YYYY-MM-DD HH:MM".
func formatStamp(t time.Time) string {
	return t.Format(stampLayout)
}

// parseStamp reads a "YYYY-MM-DD HH:MM" stamp as UTC.
func parseStamp(s string) (time.Time, error) {
	return time.Parse(stampLayout, s)
}

func main() {
	when := time.Date(2026, time.July, 4, 8, 5, 0, 0, time.UTC)
	stamp := formatStamp(when)
	fmt.Println("formatted:", stamp)

	parsed, err := parseStamp("2026-12-31 23:59")
	fmt.Println("parsed:", parsed.Format(time.RFC3339), "err:", err)
}
