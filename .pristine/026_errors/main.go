// 026: Errors. Functions return an error as a normal value; sentinel errors
// identify known failures and errors.Is sees through wrapping.
//
// Scenario: a checkout looks up item prices and must distinguish an empty
// item name, a missing item, and a corrupt negative price.
//
// Rules:
//  1. An empty item name returns an error wrapping ErrEmptyName.
//  2. A missing item returns an error wrapping ErrNotFound and mentioning
//     the item.
//  3. A stored negative price returns an error wrapping ErrInvalidPrice.
//  4. Total sums items in order, returns 0 on any failure, and wraps the
//     failing item's error so errors.Is still works.
//
// TODO: implement Price and Total. Use %w when adding context.
package main

import (
	"errors"
	"fmt"
)

// Errors returned by the price lookup.
var (
	ErrEmptyName    = errors.New("empty item name")
	ErrNotFound     = errors.New("item not found")
	ErrInvalidPrice = errors.New("invalid price")
)

// Price returns the price of item.
func Price(prices map[string]int, item string) (int, error) {
	return 0, nil
}

// Total sums the prices of items and stops at the first error.
func Total(prices map[string]int, items []string) (int, error) {
	return 0, nil
}

func main() {
	prices := map[string]int{"apple": 120, "pear": 80, "debt": -5}
	fmt.Println(Total(prices, []string{"apple", "pear"}))
	if _, err := Total(prices, []string{"apple", "plum"}); err != nil {
		fmt.Println("error:", err, "is not found:", errors.Is(err, ErrNotFound))
	}
}
