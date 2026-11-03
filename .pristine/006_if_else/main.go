// 006: If/Else. Conditions drive branches; an if can declare a variable in an
// init statement, and an else-if chain picks exactly one branch.
//
// Scenario: an exam board converts scores to letter grades while a gateway
// labels HTTP status codes by class and a thermostat clamps a target.
//
// Rules:
//  1. grade returns "A" for 90-100, "B" for 80-89, "C" for 70-79, "D" for
//     60-69 and "F" for 0-59. Anything outside 0-100 is "invalid".
//  2. httpClass uses an if with an init statement on code/100 and returns
//     "success" for 2xx, "redirect" for 3xx, "client error" for 4xx,
//     "server error" for 5xx, and "unknown" otherwise.
//  3. clamp returns lo when n < lo, hi when n > hi, and n otherwise.
//
// TODO: implement the functions below. Keep main() as the demo harness.
package main

import "fmt"

// grade maps a score to a letter grade.
func grade(score int) string {
	return ""
}

// httpClass names the class of an HTTP status code.
func httpClass(code int) string {
	return ""
}

// clamp limits n to the inclusive range lo..hi.
func clamp(n, lo, hi int) int {
	return 0
}

func main() {
	fmt.Println(grade(87))
	fmt.Println(httpClass(404))
	fmt.Println(clamp(42, 0, 10))
}
