// 027: Custom Errors. A struct can implement error, carry structured data,
// and be recovered from a wrapped chain with errors.As.
//
// Scenario: a signup form returns field-level validation failures that the
// HTTP layer needs to inspect by field name.
//
// Rules:
//  1. ValidationError implements error with a pointer receiver and formats
//     as "field: reason".
//  2. ValidateAge rejects negative ages with reason "must be non-negative"
//     and ages over 150 with reason "must be at most 150"; 0..150 is valid.
//  3. ValidateName rejects empty or whitespace-only names with reason
//     "must not be empty".
//  4. FieldOf finds a ValidationError anywhere in err's chain and reports
//     its field; it returns false for nil and non-validation errors.
//
// TODO: implement Error, ValidateAge, ValidateName and FieldOf.
package main

import "fmt"

// ValidationError describes one invalid field.
type ValidationError struct {
	Field  string
	Reason string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return ""
}

// ValidateAge checks that age is within 0..150.
func ValidateAge(age int) error {
	return nil
}

// ValidateName checks that name is not blank.
func ValidateName(name string) error {
	return nil
}

// FieldOf extracts the field name from a ValidationError in err's chain.
func FieldOf(err error) (string, bool) {
	return "", false
}

func main() {
	if err := ValidateAge(-1); err != nil {
		fmt.Println("age error:", err)
		if field, ok := FieldOf(err); ok {
			fmt.Println("field:", field)
		}
	}
	fmt.Println("valid name error:", ValidateName("Ada"))
}
