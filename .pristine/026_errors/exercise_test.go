package main

import (
	"errors"
	"strings"
	"testing"
)

func TestPrice(t *testing.T) {
	prices := map[string]int{"apple": 120, "pear": 80, "debt": -5}

	tests := []struct {
		name    string
		item    string
		want    int
		wantErr error
	}{
		{"known item", "apple", 120, nil},
		{"second known item", "pear", 80, nil},
		{"missing item", "plum", 0, ErrNotFound},
		{"empty name", "", 0, ErrEmptyName},
		{"negative price", "debt", 0, ErrInvalidPrice},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Price(prices, tt.item)
			if tt.wantErr == nil {
				if err != nil {
					t.Fatalf("Price(%q) error = %v, want nil", tt.item, err)
				}
				if got != tt.want {
					t.Errorf("Price(%q) = %d, want %d", tt.item, got, tt.want)
				}
				return
			}

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Price(%q) error = %v, want errors.Is(_, %v)", tt.item, err, tt.wantErr)
			}
			if got != 0 {
				t.Errorf("Price(%q) = %d on error, want 0", tt.item, got)
			}
		})
	}

	t.Run("missing item is named in the error", func(t *testing.T) {
		_, err := Price(prices, "plum")
		if err == nil {
			t.Fatal("Price(\"plum\") error = nil, want ErrNotFound")
		}
		if !strings.Contains(err.Error(), "plum") {
			t.Errorf("Price(\"plum\") error = %q, want it to mention the item", err)
		}
	})
}

func TestTotal(t *testing.T) {
	prices := map[string]int{"apple": 120, "pear": 80, "debt": -5}

	tests := []struct {
		name    string
		items   []string
		want    int
		wantErr error
	}{
		{"two items", []string{"apple", "pear"}, 200, nil},
		{"empty order", nil, 0, nil},
		{"single item", []string{"pear"}, 80, nil},
		{"missing second item", []string{"apple", "plum"}, 0, ErrNotFound},
		{"empty item name", []string{"apple", ""}, 0, ErrEmptyName},
		{"negative price", []string{"debt"}, 0, ErrInvalidPrice},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Total(prices, tt.items)
			if tt.wantErr == nil {
				if err != nil {
					t.Fatalf("Total(%v) error = %v, want nil", tt.items, err)
				}
				if got != tt.want {
					t.Errorf("Total(%v) = %d, want %d", tt.items, got, tt.want)
				}
				return
			}

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Total(%v) error = %v, want errors.Is(_, %v)", tt.items, err, tt.wantErr)
			}
			if got != 0 {
				t.Errorf("Total(%v) = %d on error, want 0", tt.items, got)
			}
		})
	}

	t.Run("wrapped error keeps the failing item", func(t *testing.T) {
		_, err := Total(prices, []string{"apple", "plum"})
		if err == nil {
			t.Fatal("Total with a missing item error = nil, want ErrNotFound")
		}
		if !strings.Contains(err.Error(), "plum") {
			t.Errorf("Total error = %q, want it to mention plum", err)
		}
	})
}
