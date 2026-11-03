// 048: Panic. panic stops normal execution and unwinds the stack; callers
// that can continue must recover it deliberately.
//
// Scenario: an age field arrives from an untrusted form, and invalid values
// are programmer errors that must fail loudly during validation.
//
// Rules:
//  1. age < 0 panics with exactly "age must be >= 0".
//  2. age > 130 panics with exactly "age must be <= 130".
//  3. 0..17 returns "minor", 18..64 returns "adult", 65..130 returns "senior".
//  4. Boundaries 0, 17, 18, 64, 65 and 130 are valid.
//
// TODO: implement classifyAge. Keep main() as the demo harness.
package main

import "fmt"

// classifyAge returns the age band or panics when age is outside 0..130.
func classifyAge(age int) string {
	// TODO
	return ""
}

func main() {
	fmt.Println(classifyAge(30))
}
