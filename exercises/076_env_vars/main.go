// 076: Environment Variables. os.Getenv returns "" both for an unset key and
// for an empty value, so defaults belong in one helper.
//
// Scenario: the service reads DOJO_ENV and PORT from the process environment.
//
// Rules:
//  1. envOr(key, fallback) returns the variable value, or fallback when the
//     variable is unset or empty.
//  2. envInt(key, fallback) parses the value as a base-10 int, or returns
//     fallback when it is unset, empty or not an integer.
//  3. Neither function mutates the environment.
//  4. Tests set variables with t.Setenv, never os.Setenv.
//
// TODO: implement envOr and envInt below. Keep main() as the demo harness.
package main

import "fmt"

// envOr returns the value of key, or fallback when it is unset or empty.
func envOr(key, fallback string) string {
	// TODO
	return ""
}

// envInt returns the integer value of key, or fallback when unusable.
func envInt(key string, fallback int) int {
	// TODO
	return 0
}

func main() {
	fmt.Println("env:", envOr("DOJO_ENV", "development"))
	fmt.Println("port:", envInt("PORT", 8080))
}
