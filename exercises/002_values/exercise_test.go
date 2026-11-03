package main

import "testing"

func TestSubtotal(t *testing.T) {
	tests := []struct {
		name string
		qty  int
		unit int
		want int
	}{
		{"three items", 3, 250, 750},
		{"single item", 1, 5, 5},
		{"bulk", 12, 99, 1188},
		{"zero quantity", 0, 250, 0},
		{"negative quantity", -2, 250, 0},
		{"negative unit price", 3, -250, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtotal(tt.qty, tt.unit); got != tt.want {
				t.Errorf("subtotal(%d, %d) = %d, want %d", tt.qty, tt.unit, got, tt.want)
			}
		})
	}
}

func TestIsBulk(t *testing.T) {
	tests := []struct {
		name string
		qty  int
		want bool
	}{
		{"below threshold", 9, false},
		{"at threshold", 10, true},
		{"above threshold", 12, true},
		{"zero", 0, false},
		{"negative", -3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBulk(tt.qty); got != tt.want {
				t.Errorf("isBulk(%d) = %t, want %t", tt.qty, got, tt.want)
			}
		})
	}
}

func TestReceiptLine(t *testing.T) {
	tests := []struct {
		name string
		qty  int
		unit int
		want string
	}{
		{"plain", 3, 250, "3 x $2.50 = $7.50"},
		{"bulk", 12, 250, "12 x $2.50 = $30.00 (bulk)"},
		{"small change", 1, 5, "1 x $0.05 = $0.05"},
		{"zero quantity", 0, 250, "0 x $2.50 = $0.00"},
		{"negative quantity", -2, 250, "0 x $2.50 = $0.00"},
		{"negative price", 2, -250, "2 x $0.00 = $0.00"},
		{"whole dollars", 4, 100, "4 x $1.00 = $4.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := receiptLine(tt.qty, tt.unit); got != tt.want {
				t.Errorf("receiptLine(%d, %d) = %q, want %q", tt.qty, tt.unit, got, tt.want)
			}
		})
	}
}
