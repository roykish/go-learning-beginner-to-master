// Package calculator provides mathematical calculation functions.
// This demonstrates functions with multiple return values and error handling.
package calculator

import (
	"cli-calculator/internal/constants"
	"cli-calculator/internal/errors"
	"fmt"
	"math"
)

// Calculate performs a calculation based on the operation and operands.
// This demonstrates function parameters, return values, and error handling.
func Calculate(operation constants.Operation, operands []float64) (float64, error) {
	// Validate operation and operands
	if err := validateCalculation(operation, operands); err != nil {
		return 0, err
	}

	// Perform calculation based on operation
	switch operation {
	case constants.OpAddition:
		return add(operands), nil
	case constants.OpSubtraction:
		return subtract(operands), nil
	case constants.OpMultiplication:
		return multiply(operands), nil
	case constants.OpDivision:
		return divide(operands[0], operands[1])
	case constants.OpPower:
		return power(operands[0], operands[1]), nil
	case constants.OpSquareRoot:
		return squareRoot(operands[0])
	case constants.OpModulo:
		return modulo(operands[0], operands[1])
	case constants.OpFactorial:
		return factorial(operands[0])
	default:
		return 0, errors.NewCalculationError(
			operation.String(),
			operands,
			"unsupported operation",
			errors.ErrInvalidOperation,
		)
	}
}

// validateCalculation validates the operation and operands.
func validateCalculation(operation constants.Operation, operands []float64) error {
	// Check if we have operands
	if len(operands) == 0 {
		return errors.NewValidationError("operands", "none", "at least one operand is required")
	}

	// Validate operand count based on operation
	requiredOperands := getRequiredOperandCount(operation)
	if len(operands) < requiredOperands {
		return errors.NewValidationError(
			"operands",
			fmt.Sprintf("%d", len(operands)),
			fmt.Sprintf("%s requires %d operands, got %d", operation.String(), requiredOperands, len(operands)),
		)
	}

	// Validate operand ranges
	for i, val := range operands {
		if math.IsNaN(val) {
			return errors.NewValidationError(
				fmt.Sprintf("operand[%d]", i),
				"NaN",
				"operand cannot be NaN",
			)
		}
		if math.IsInf(val, 0) {
			return errors.NewValidationError(
				fmt.Sprintf("operand[%d]", i),
				"Inf",
				"operand cannot be infinity",
			)
		}
		if val > constants.MaxNumberInputValue || val < constants.MinNumberInputValue {
			return errors.NewValidationError(
				fmt.Sprintf("operand[%d]", i),
				fmt.Sprintf("%f", val),
				fmt.Sprintf("operand must be between %e and %e", constants.MinNumberInputValue, constants.MaxNumberInputValue),
			)
		}
	}

	return nil
}

// getRequiredOperandCount returns the number of operands required for an operation.
func getRequiredOperandCount(operation constants.Operation) int {
	switch operation {
	case constants.OpSquareRoot, constants.OpFactorial:
		return 1
	case constants.OpAddition, constants.OpSubtraction, constants.OpMultiplication:
		return 1 // Can work with 1+ operands
	default:
		return 2 // Binary operations
	}
}

// Basic arithmetic operations

// add adds multiple numbers together.
// This demonstrates variadic-style operations with slices.
func add(operands []float64) float64 {
	result := 0.0
	for _, val := range operands {
		result += val
	}
	return result
}

// subtract subtracts subsequent numbers from the first.
func subtract(operands []float64) float64 {
	if len(operands) == 0 {
		return 0
	}
	if len(operands) == 1 {
		return operands[0]
	}

	result := operands[0]
	for _, val := range operands[1:] {
		result -= val
	}
	return result
}

// multiply multiplies multiple numbers together.
func multiply(operands []float64) float64 {
	if len(operands) == 0 {
		return 0
	}

	result := 1.0
	for _, val := range operands {
		result *= val
	}
	return result
}

// divide divides the first number by the second.
// This demonstrates error handling for invalid operations.
func divide(a, b float64) (float64, error) {
	// Check for division by zero
	if b == 0 {
		return 0, errors.NewCalculationError(
			"Division",
			[]float64{a, b},
			"division by zero",
			errors.ErrDivisionByZero,
		)
	}

	result := a / b

	// Check for overflow/underflow
	if math.IsInf(result, 0) {
		return 0, errors.NewCalculationError(
			"Division",
			[]float64{a, b},
			"result is infinity (overflow)",
			nil,
		)
	}

	return result, nil
}

// Advanced operations

// power raises a to the power of b.
func power(a, b float64) float64 {
	return math.Pow(a, b)
}

// squareRoot calculates the square root of a number.
func squareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.NewCalculationError(
			"SquareRoot",
			[]float64{a},
			"cannot calculate square root of negative number",
			errors.ErrNegativeSquareRoot,
		)
	}
	return math.Sqrt(a), nil
}

// modulo calculates the remainder of a divided by b.
func modulo(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.NewCalculationError(
			"Modulo",
			[]float64{a, b},
			"division by zero in modulo operation",
			errors.ErrDivisionByZero,
		)
	}
	return math.Mod(a, b), nil
}

// factorial calculates the factorial of a number.
func factorial(n float64) (float64, error) {
	// Check if n is an integer
	if n != math.Floor(n) {
		return 0, errors.NewCalculationError(
			"Factorial",
			[]float64{n},
			"factorial requires an integer",
			errors.ErrInvalidInput,
		)
	}

	// Check if n is negative
	if n < 0 {
		return 0, errors.NewCalculationError(
			"Factorial",
			[]float64{n},
			"factorial of negative number is undefined",
			errors.ErrInvalidInput,
		)
	}

	// Check for overflow (factorial grows very quickly)
	if n > 170 {
		return 0, errors.NewCalculationError(
			"Factorial",
			[]float64{n},
			"factorial result would overflow (too large)",
			errors.ErrOutOfRange,
		)
	}

	// Calculate factorial iteratively
	result := 1.0
	for i := 2.0; i <= n; i++ {
		result *= i
	}

	return result, nil
}

// FormatResult formats a calculation result with the specified precision.
// This demonstrates string formatting and type conversion.
func FormatResult(result float64, precision int) string {
	// Handle special cases
	if math.IsNaN(result) {
		return "NaN"
	}
	if math.IsInf(result, 1) {
		return "+Inf"
	}
	if math.IsInf(result, -1) {
		return "-Inf"
	}

	// Format with specified precision
	format := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(format, result)
}
