// 002: Values. Go has a small set of basic value types: int, float64, string
// and bool. Money is safest kept in integer cents; convert only when printing.
//
// Scenario: a corner shop prints receipts. Prices are stored in cents so that
// 2.50 never becomes 2.4999999, and bulk orders get a marker on the line.
//
// Rules:
//  1. subtotal multiplies integer cents; a non-positive quantity or unit price
//     yields 0.
//  2. isBulk reports true when the quantity is at least 10.
//  3. receiptLine renders "<qty> x $<unit> = $<total>" with exactly two
//     decimals, and appends " (bulk)" for bulk orders.
//  4. Negative quantities or prices are treated as zero on the printed line.
//
// TODO: implement the functions below. Keep main() as the demo harness.
package main

import "fmt"

// subtotal returns qty * unitCents in cents.
func subtotal(qty, unitCents int) int {
	return 0
}

// isBulk reports whether qty qualifies as a bulk order.
func isBulk(qty int) bool {
	return false
}

// receiptLine renders one receipt line.
func receiptLine(qty, unitCents int) string {
	return ""
}

func main() {
	fmt.Println(receiptLine(3, 250))
	fmt.Println(receiptLine(12, 250))
	fmt.Println(subtotal(3, 250), isBulk(12))
}
