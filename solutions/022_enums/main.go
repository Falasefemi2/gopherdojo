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

func (d Weekday) String() string {
	if d < Sunday || d > Saturday {
		return fmt.Sprintf("Weekday(%d)", int(d))
	}
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

func parseWeekday(s string) (Weekday, bool) {
	for d := Sunday; d <= Saturday; d++ {
		if d.String() == s {
			return d, true
		}
	}
	return Sunday, false
}

func isWeekend(d Weekday) bool {
	return d == Saturday || d == Sunday
}

func nextDay(d Weekday) Weekday {
	if d < Sunday || d > Saturday {
		return d
	}
	if d == Saturday {
		return Sunday
	}
	return d + 1
}

func main() {
	d := Wednesday
	fmt.Println(d, "weekend?", isWeekend(d), "next:", nextDay(d))
	if parsed, ok := parseWeekday("Saturday"); ok {
		fmt.Println("parsed:", parsed, "weekend?", isWeekend(parsed))
	}
}
