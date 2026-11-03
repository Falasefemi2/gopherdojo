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

func Price(prices map[string]int, item string) (int, error) {
	if item == "" {
		return 0, fmt.Errorf("%w: name must not be empty", ErrEmptyName)
	}
	price, ok := prices[item]
	if !ok {
		return 0, fmt.Errorf("%w: %q", ErrNotFound, item)
	}
	if price < 0 {
		return 0, fmt.Errorf("%w: %q has price %d", ErrInvalidPrice, item, price)
	}
	return price, nil
}

func Total(prices map[string]int, items []string) (int, error) {
	total := 0
	for _, item := range items {
		price, err := Price(prices, item)
		if err != nil {
			return 0, fmt.Errorf("total at %q: %w", item, err)
		}
		total += price
	}
	return total, nil
}

func main() {
	prices := map[string]int{"apple": 120, "pear": 80, "debt": -5}
	fmt.Println(Total(prices, []string{"apple", "pear"}))
	if _, err := Total(prices, []string{"apple", "plum"}); err != nil {
		fmt.Println("error:", err, "is not found:", errors.Is(err, ErrNotFound))
	}
}
