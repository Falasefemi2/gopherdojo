// 007: Switch. A switch selects one case; without a tag it is a cleaner
// else-if chain, and fallthrough continues into the next case.
//
// Scenario: a warehouse classifies stock counts, reports the days in a month,
// and prices shipping by weight tier with fallthrough accumulation.
//
// Rules:
//  1. classify returns "negative", "zero", "digit" (1-9) or "large" (>=10).
//  2. daysInMonth returns 31 for 1,3,5,7,8,10,12; 30 for 4,6,9,11; 29 for
//     February in a leap year and 28 otherwise; 0 for any other month.
//  3. shippingCents adds a 200c base tier for any positive weight, plus 300c
//     when weight > 500g, plus 500c when weight > 1000g. Zero or negative
//     weight ships free (0).
package main

import "fmt"

// classify labels an integer.
func classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	case n < 10:
		return "digit"
	default:
		return "large"
	}
}

// daysInMonth returns the number of days in month.
func daysInMonth(month int, leap bool) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if leap {
			return 29
		}
		return 28
	default:
		return 0
	}
}

// shippingCents prices a parcel by weight in grams.
func shippingCents(weightG int) int {
	total := 0
	switch {
	case weightG > 1000:
		total += 500
		fallthrough
	case weightG > 500:
		total += 300
		fallthrough
	case weightG > 0:
		total += 200
	}
	return total
}

func main() {
	fmt.Println(classify(7))
	fmt.Println(daysInMonth(2, true))
	fmt.Println(shippingCents(750))
}
