package main

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"negative", -1, "negative"},
		{"zero", 0, "zero"},
		{"single digit", 7, "digit"},
		{"nine", 9, "digit"},
		{"ten", 10, "large"},
		{"large", 42, "large"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classify(tt.n); got != tt.want {
				t.Errorf("classify(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		name  string
		month int
		leap  bool
		want  int
	}{
		{"january", 1, false, 31},
		{"april", 4, false, 30},
		{"july", 7, false, 31},
		{"november", 11, false, 30},
		{"december", 12, false, 31},
		{"february common", 2, false, 28},
		{"february leap", 2, true, 29},
		{"month zero", 0, false, 0},
		{"month thirteen", 13, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysInMonth(tt.month, tt.leap); got != tt.want {
				t.Errorf("daysInMonth(%d, %t) = %d, want %d", tt.month, tt.leap, got, tt.want)
			}
		})
	}
}

func TestShippingCents(t *testing.T) {
	tests := []struct {
		name    string
		weightG int
		want    int
	}{
		{"heavy", 1200, 1000},
		{"just over one kilo", 1001, 1000},
		{"exactly one kilo", 1000, 500},
		{"just over half kilo", 501, 500},
		{"exactly half kilo", 500, 200},
		{"light", 1, 200},
		{"zero", 0, 0},
		{"negative", -5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shippingCents(tt.weightG); got != tt.want {
				t.Errorf("shippingCents(%d) = %d, want %d", tt.weightG, got, tt.want)
			}
		})
	}
}
