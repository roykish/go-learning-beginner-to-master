// Package validation provides input validation functions.
// This demonstrates string manipulation, type conversion, and validation patterns.
package validation

import (
	"cli-calculator/internal/constants"
	"cli-calculator/internal/errors"
	"strconv"
	"strings"
)

// ValidateMenuOption validates main menu input.
// This demonstrates validation with custom error types.
func ValidateMenuOption(input string) (constants.MenuOption, error) {
	// Clean the input
	trimmed := strings.TrimSpace(input)

	// Convert to number
	num, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, errors.NewValidationError("menu_option", trimmed, "not a valid number")
	}

	// Check range
	if num < constants.MinMenuOption || num > constants.MaxMenuOption {
		return 0, errors.NewValidationError(
			"menu_option",
			trimmed,
			"must be between 1 and 7",
		)
	}

	return constants.MenuOption(num), nil
}

// ValidateBasicOperation validates basic calculator operation input.
func ValidateBasicOperation(input string) (constants.Operation, error) {
	// Clean the input
	trimmed := strings.TrimSpace(input)

	// Convert to number
	num, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, errors.NewValidationError("operation", trimmed, "not a valid number")
	}

	// Check range (1-4 for basic operations)
	if num < constants.MinBasicCalcOption || num > constants.MaxBasicCalcOption {
		return 0, errors.NewValidationError(
			"operation",
			trimmed,
			"must be between 1 and 4",
		)
	}

	// Map to operation constants
	operations := []constants.Operation{
		constants.OpAddition,
		constants.OpSubtraction,
		constants.OpMultiplication,
		constants.OpDivision,
	}

	return operations[num-1], nil
}

// ValidateNumber validates and parses a number input.
// This demonstrates float parsing with validation and error handling.
func ValidateNumber(input string) (float64, error) {
	// Clean the input
	trimmed := strings.TrimSpace(input)

	// Check for empty input
	if trimmed == "" {
		return 0, errors.NewValidationError("number", trimmed, "cannot be empty")
	}

	// Parse as float64
	num, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return 0, errors.NewValidationError("number", trimmed, "not a valid number")
	}

	// Validate range
	if num > constants.MaxNumberInputValue || num < constants.MinNumberInputValue {
		return 0, errors.NewValidationError(
			"number",
			trimmed,
			"value out of allowed range",
		)
	}

	return num, nil
}

// ValidatePrecision validates precision input for number formatting.
func ValidatePrecision(precision int) error {
	if precision < 0 || precision > 15 {
		return errors.NewValidationError(
			"precision",
			strconv.Itoa(precision),
			"must be between 0 and 15",
		)
	}
	return nil
}

// ValidateYesNo validates yes/no input.
// This demonstrates case-insensitive string comparison.
func ValidateYesNo(input string) (bool, error) {
	trimmed := strings.TrimSpace(strings.ToLower(input))

	switch trimmed {
	case "y", "yes", "1", "true":
		return true, nil
	case "n", "no", "0", "false":
		return false, nil
	default:
		return false, errors.NewValidationError(
			"yes_no",
			input,
			"must be yes/no, y/n, or true/false",
		)
	}
}