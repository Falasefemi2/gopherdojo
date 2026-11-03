// 022: Enums. Go builds enums from a named integer type, an iota block, and
// a String method for readable output.
//
// Scenario: a scheduling API stores weekdays as small integers but must log
// and parse them as English names.
//
// Rules:
//  1. Sunday is 0 and each following day adds one, ending at Saturday == 6.
//  2. String returns the English name, or "Weekday(n)" for unknown values.
//  3. parseWeekday accepts only exact English names and reports false
//     otherwise.
//  4. isWeekend is true only for Saturday and Sunday, and nextDay wraps
//     Saturday to Sunday while leaving unknown values unchanged.
//
// TODO: implement String, parseWeekday, isWeekend and nextDay.
package main

import "fmt"

// Weekday is a day of the week; zero is Sunday.
type Weekday int

// The days of the week, in order.
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// String returns the English day name.
func (d Weekday) String() string {
	return ""
}

// parseWeekday maps an exact English day name to its value.
func parseWeekday(s string) (Weekday, bool) {
	return Sunday, false
}

// isWeekend reports whether d is Saturday or Sunday.
func isWeekend(d Weekday) bool {
	return false
}

// nextDay returns the following day, wrapping Saturday to Sunday.
func nextDay(d Weekday) Weekday {
	return Sunday
}

func main() {
	d := Wednesday
	fmt.Println(d, "weekend?", isWeekend(d), "next:", nextDay(d))
	if parsed, ok := parseWeekday("Saturday"); ok {
		fmt.Println("parsed:", parsed, "weekend?", isWeekend(parsed))
	}
}
