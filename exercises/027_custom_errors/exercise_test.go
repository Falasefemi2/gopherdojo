package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestValidationErrorFormat(t *testing.T) {
	tests := []struct {
		name string
		err  *ValidationError
		want string
	}{
		{
			name: "age",
			err:  &ValidationError{Field: "age", Reason: "must be non-negative"},
			want: "age: must be non-negative",
		},
		{
			name: "name",
			err:  &ValidationError{Field: "name", Reason: "must not be empty"},
			want: "name: must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("ValidationError.Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name       string
		age        int
		wantErr    bool
		wantReason string
	}{
		{"negative", -1, true, "must be non-negative"},
		{"zero is valid", 0, false, ""},
		{"ordinary", 30, false, ""},
		{"upper bound", 150, false, ""},
		{"too old", 151, true, "must be at most 150"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)
			if !tt.wantErr {
				if err != nil {
					t.Fatalf("ValidateAge(%d) error = %v, want nil", tt.age, err)
				}
				return
			}
			if err == nil {
				t.Fatalf("ValidateAge(%d) error = nil, want *ValidationError", tt.age)
			}

			var ve *ValidationError
			if !errors.As(err, &ve) {
				t.Fatalf("ValidateAge(%d) error type = %T, want *ValidationError", tt.age, err)
			}
			if ve.Field != "age" {
				t.Errorf("ValidateAge(%d) field = %q, want %q", tt.age, ve.Field, "age")
			}
			if ve.Reason != tt.wantReason {
				t.Errorf("ValidateAge(%d) reason = %q, want %q", tt.age, ve.Reason, tt.wantReason)
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"ordinary", "Ada", false},
		{"empty", "", true},
		{"spaces only", "   ", true},
		{"tab only", "\t", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.in)
			if !tt.want {
				if err != nil {
					t.Fatalf("ValidateName(%q) error = %v, want nil", tt.in, err)
				}
				return
			}
			if err == nil {
				t.Fatalf("ValidateName(%q) error = nil, want *ValidationError", tt.in)
			}

			var ve *ValidationError
			if !errors.As(err, &ve) {
				t.Fatalf("ValidateName(%q) error type = %T, want *ValidationError", tt.in, err)
			}
			if ve.Field != "name" {
				t.Errorf("ValidateName(%q) field = %q, want %q", tt.in, ve.Field, "name")
			}
			if ve.Reason != "must not be empty" {
				t.Errorf("ValidateName(%q) reason = %q, want %q", tt.in, ve.Reason, "must not be empty")
			}
		})
	}
}

func TestFieldOf(t *testing.T) {
	base := &ValidationError{Field: "email", Reason: "must contain @"}
	wrapped := fmt.Errorf("signup rejected: %w", base)

	tests := []struct {
		name      string
		err       error
		wantField string
		wantOK    bool
	}{
		{"direct", base, "email", true},
		{"wrapped", wrapped, "email", true},
		{"nil", nil, "", false},
		{"plain error", errors.New("boom"), "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, ok := FieldOf(tt.err)
			if ok != tt.wantOK {
				t.Fatalf("FieldOf(%v) ok = %v, want %v", tt.err, ok, tt.wantOK)
			}
			if field != tt.wantField {
				t.Errorf("FieldOf(%v) field = %q, want %q", tt.err, field, tt.wantField)
			}
		})
	}

	t.Run("wrapped message keeps context", func(t *testing.T) {
		if !strings.Contains(wrapped.Error(), "signup rejected") {
			t.Errorf("wrapped error = %q, want it to keep the outer context", wrapped)
		}
	})
}
