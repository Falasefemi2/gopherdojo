// 050: Recover. recover stops a panicking goroutine, but it only works when
// called directly from a deferred function.
//
// Scenario: a pricing pipeline divides by a user-supplied quantity. Bad input
// must become an ordinary error instead of crashing the server.
//
// Rules:
//  1. Integer division truncates toward zero, as Go's / operator does.
//  2. b != 0 returns the quotient and a nil error.
//  3. b == 0 must not panic: recover the runtime panic and return an error
//     whose message contains "divide by zero".
//  4. The result is 0 whenever an error is returned.
//
// TODO: implement safeDivide. Keep main() as the demo harness.
package main

import "fmt"

// safeDivide returns a/b, converting a divide-by-zero panic into an error.
func safeDivide(a, b int) (result int, err error) {
	// TODO
	return 0, nil
}

func main() {
	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(1, 0))
}
