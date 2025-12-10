// Package validation provides input validation functions with tests.
// This demonstrates testing validation logic and error cases.
package validation

import (
	"cli-calculator/internal/constants"
	"testing"
)

// TestValidateMenuOption tests menu option validation.
func TestValidateMenuOption(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected constants.MenuOption
		hasError bool
	}{
		{"valid option 1", "1", constants.MenuBasicCalculator, false},
		{"valid option 7", "7", constants.MenuExit, false},
		{"invalid option 0", "0", 0, true},
		{"invalid option 8", "8", 0, true},
		{"non-numeric", "abc", 0, true},
		{"empty string", "", 0, true},
		{"with spaces", " 3 ", constants.MenuBatchCalculations, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateMenuOption(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("%s: expected error, got nil", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tt.name, err)
				}
				if result != tt.expected {
					t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}

// TestValidateBasicOperation tests basic operation validation.
func TestValidateBasicOperation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected constants.Operation
		hasError bool
	}{
		{"addition", "1", constants.OpAddition, false},
		{"subtraction", "2", constants.OpSubtraction, false},
		{"multiplication", "3", constants.OpMultiplication, false},
		{"division", "4", constants.OpDivision, false},
		{"invalid option 0", "0", 0, true},
		{"invalid option 5", "5", 0, true},
		{"non-numeric", "xyz", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateBasicOperation(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("%s: expected error, got nil", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tt.name, err)
				}
				if result != tt.expected {
					t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}

// TestValidateNumber tests number validation.
func TestValidateNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		hasError bool
	}{
		{"positive integer", "42", 42.0, false},
		{"negative integer", "-15", -15.0, false},
		{"positive float", "3.14", 3.14, false},
		{"negative float", "-2.5", -2.5, false},
		{"zero", "0", 0.0, false},
		{"scientific notation", "1.5e2", 150.0, false},
		{"with spaces", " 10.5 ", 10.5, false},
		{"non-numeric", "abc", 0, true},
		{"empty string", "", 0, true},
		{"just a dot", ".", 0, true},
		{"multiple dots", "1.2.3", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateNumber(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("%s: expected error, got nil", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tt.name, err)
				}
				if result != tt.expected {
					t.Errorf("%s: expected %f, got %f", tt.name, tt.expected, result)
				}
			}
		})
	}
}

// TestValidatePrecision tests precision validation.
func TestValidatePrecision(t *testing.T) {
	tests := []struct {
		name      string
		precision int
		hasError  bool
	}{
		{"valid 0", 0, false},
		{"valid 5", 5, false},
		{"valid 15", 15, false},
		{"invalid negative", -1, true},
		{"invalid too high", 16, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePrecision(tt.precision)

			if tt.hasError {
				if err == nil {
					t.Errorf("%s: expected error, got nil", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tt.name, err)
				}
			}
		})
	}
}

// TestValidateYesNo tests yes/no validation.
func TestValidateYesNo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		hasError bool
	}{
		{"yes lowercase", "yes", true, false},
		{"yes uppercase", "YES", true, false},
		{"y", "y", true, false},
		{"Y", "Y", true, false},
		{"1", "1", true, false},
		{"true", "true", true, false},
		{"no lowercase", "no", false, false},
		{"no uppercase", "NO", false, false},
		{"n", "n", false, false},
		{"N", "N", false, false},
		{"0", "0", false, false},
		{"false", "false", false, false},
		{"with spaces", " yes ", true, false},
		{"invalid", "maybe", false, true},
		{"empty", "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateYesNo(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("%s: expected error, got nil", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tt.name, err)
				}
				if result != tt.expected {
					t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}
