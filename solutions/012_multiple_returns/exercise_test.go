package main

import "testing"

func TestDivmod(t *testing.T) {
	tests := []struct {
		name   string
		a      int
		b      int
		wantQ  int
		wantR  int
		wantOK bool
	}{
		{"positive", 17, 5, 3, 2, true},
		{"negative dividend", -17, 5, -3, -2, true},
		{"negative divisor", 17, -5, -3, 2, true},
		{"both negative", -17, -5, 3, -2, true},
		{"zero dividend", 0, 7, 0, 0, true},
		{"divide by zero", 5, 0, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQ, gotR, gotOK := divmod(tt.a, tt.b)
			if gotQ != tt.wantQ || gotR != tt.wantR || gotOK != tt.wantOK {
				t.Errorf("divmod(%d, %d) = (%d, %d, %t), want (%d, %d, %t)",
					tt.a, tt.b, gotQ, gotR, gotOK, tt.wantQ, tt.wantR, tt.wantOK)
			}
		})
	}
}

func TestMinmax(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c int
		wantLo  int
		wantHi  int
	}{
		{"distinct", 3, 1, 2, 1, 3},
		{"all equal", 5, 5, 5, 5, 5},
		{"mixed signs", -1, 0, -5, -5, 0},
		{"largest in the middle", 9, -2, 9, -2, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLo, gotHi := minmax(tt.a, tt.b, tt.c)
			if gotLo != tt.wantLo || gotHi != tt.wantHi {
				t.Errorf("minmax(%d, %d, %d) = (%d, %d), want (%d, %d)",
					tt.a, tt.b, tt.c, gotLo, gotHi, tt.wantLo, tt.wantHi)
			}
		})
	}
}

func TestSplitName(t *testing.T) {
	tests := []struct {
		name      string
		full      string
		wantFirst string
		wantLast  string
	}{
		{"two words", "Ada Lovelace", "Ada", "Lovelace"},
		{"three words", "Ada King Lovelace", "Ada King", "Lovelace"},
		{"single word", "Ada", "Ada", ""},
		{"empty", "", "", ""},
		{"whitespace only", "   ", "", ""},
		{"extra spaces", "  Ada  Lovelace  ", "Ada", "Lovelace"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotLast := splitName(tt.full)
			if gotFirst != tt.wantFirst || gotLast != tt.wantLast {
				t.Errorf("splitName(%q) = (%q, %q), want (%q, %q)",
					tt.full, gotFirst, gotLast, tt.wantFirst, tt.wantLast)
			}
		})
	}
}
