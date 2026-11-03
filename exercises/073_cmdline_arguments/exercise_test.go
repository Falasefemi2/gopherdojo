package main

import "testing"

func TestSumArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    int
		wantErr bool
	}{
		{"no args", nil, 0, false},
		{"empty slice", []string{}, 0, false},
		{"single", []string{"42"}, 42, false},
		{"several", []string{"10", "20", "12"}, 42, false},
		{"negative", []string{"5", "-8"}, -3, false},
		{"padded", []string{" 7 ", " 3"}, 10, false},
		{"zero", []string{"0"}, 0, false},
		{"not a number", []string{"10", "ten"}, 0, true},
		{"empty string", []string{""}, 0, true},
		{"float", []string{"1.5"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sumArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("sumArgs(%q) error = nil, want an error", tt.args)
				}
				if got != 0 {
					t.Errorf("sumArgs(%q) = %d on error, want 0", tt.args, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("sumArgs(%q) unexpected error: %v", tt.args, err)
			}
			if got != tt.want {
				t.Errorf("sumArgs(%q) = %d, want %d", tt.args, got, tt.want)
			}
		})
	}
}
