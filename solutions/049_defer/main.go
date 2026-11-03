// 049: Defer. Deferred calls run when the surrounding function returns, and
// they run in last-in-first-out order, like a stack.
//
// Scenario: cleanup steps are registered while resources are acquired; the
// audit log must show them unwinding in reverse acquisition order.
//
// Rules:
//  1. Register one deferred call per label while iterating the input.
//  2. Deferred calls append their label, so the result is reversed input.
//  3. Capture the loop value at defer time (arguments are evaluated then).
//  4. nil or empty input returns a non-nil empty slice; input is not mutated.
package main

import "fmt"

// unwind registers a deferred append for every label and returns the labels
// in the order the deferred calls actually ran.
func unwind(labels []string) (order []string) {
	order = make([]string, 0, len(labels))
	for _, label := range labels {
		defer func(label string) {
			order = append(order, label)
		}(label)
	}
	return order
}

func main() {
	fmt.Println(unwind([]string{"open", "read", "close"}))
}
