// Package calculator provides mathematical calculation functions with tests.
// This demonstrates Go's testing framework and table-driven tests.
package calculator

import (
	"cli-calculator/internal/constants"
	"math"
	"testing"
)

// TestCalculateAddition tests the addition operation.
// This demonstrates basic test function structure.
func TestCalculateAddition(t *testing.T) {
	result, err := Calculate(constants.OpAddition, []float64{5, 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("Expected 8, got %f", result)
	}
}

// TestCalculateSubtraction tests the subtraction operation.
func TestCalculateSubtraction(t *testing.T) {
	result, err := Calculate(constants.OpSubtraction, []float64{10, 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 7 {
		t.Errorf("Expected 7, got %f", result)
	}
}

// TestCalculateMultiplication tests the multiplication operation.
func TestCalculateMultiplication(t *testing.T) {
	result, err := Calculate(constants.OpMultiplication, []float64{4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 20 {
		t.Errorf("Expected 20, got %f", result)
	}
}

// TestCalculateDivision tests the division operation.
func TestCalculateDivision(t *testing.T) {
	result, err := Calculate(constants.OpDivision, []float64{10, 2})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %f", result)
	}
}

// TestCalculateDivisionByZero tests division by zero error handling.
// This demonstrates testing for expected errors.
func TestCalculateDivisionByZero(t *testing.T) {
	_, err := Calculate(constants.OpDivision, []float64{10, 0})
	if err == nil {
		t.Error("Expected error for division by zero, got nil")
	}
}

// TestCalculatePower tests the power operation.
func TestCalculatePower(t *testing.T) {
	result, err := Calculate(constants.OpPower, []float64{2, 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("Expected 8, got %f", result)
	}
}

// TestCalculateSquareRoot tests the square root operation.
func TestCalculateSquareRoot(t *testing.T) {
	result, err := Calculate(constants.OpSquareRoot, []float64{16})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 4 {
		t.Errorf("Expected 4, got %f", result)
	}
}

// TestCalculateSquareRootNegative tests square root of negative number.
func TestCalculateSquareRootNegative(t *testing.T) {
	_, err := Calculate(constants.OpSquareRoot, []float64{-16})
	if err == nil {
		t.Error("Expected error for square root of negative number, got nil")
	}
}

// TestCalculateModulo tests the modulo operation.
func TestCalculateModulo(t *testing.T) {
	result, err := Calculate(constants.OpModulo, []float64{10, 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected 1, got %f", result)
	}
}

// TestCalculateFactorial tests the factorial operation.
// This demonstrates table-driven tests.
func TestCalculateFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
		hasError bool
	}{
		{"factorial of 0", 0, 1, false},
		{"factorial of 1", 1, 1, false},
		{"factorial of 5", 5, 120, false},
		{"factorial of 10", 10, 3628800, false},
		{"factorial of negative", -5, 0, true},
		{"factorial of non-integer", 3.5, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(constants.OpFactorial, []float64{tt.input})

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

// TestFormatResult tests the result formatting function.
func TestFormatResult(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		precision int
		expected  string
	}{
		{"integer", 5.0, 2, "5.00"},
		{"decimal", 3.14159, 2, "3.14"},
		{"high precision", 3.14159265359, 5, "3.14159"},
		{"zero", 0.0, 2, "0.00"},
		{"negative", -7.5, 1, "-7.5"},
		{"NaN", math.NaN(), 2, "NaN"},
		{"positive infinity", math.Inf(1), 2, "+Inf"},
		{"negative infinity", math.Inf(-1), 2, "-Inf"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatResult(tt.value, tt.precision)
			if result != tt.expected {
				t.Errorf("%s: expected '%s', got '%s'", tt.name, tt.expected, result)
			}
		})
	}
}

// TestCalculateInvalidOperation tests handling of invalid operations.
func TestCalculateInvalidOperation(t *testing.T) {
	_, err := Calculate(constants.Operation(99), []float64{1, 2})
	if err == nil {
		t.Error("Expected error for invalid operation, got nil")
	}
}

// TestCalculateEmptyOperands tests validation with empty operands.
func TestCalculateEmptyOperands(t *testing.T) {
	_, err := Calculate(constants.OpAddition, []float64{})
	if err == nil {
		t.Error("Expected error for empty operands, got nil")
	}
}

// BenchmarkCalculateAddition benchmarks the addition operation.
// This demonstrates benchmark functions in Go.
func BenchmarkCalculateAddition(b *testing.B) {
	operands := []float64{123.456, 789.012}
	for i := 0; i < b.N; i++ {
		Calculate(constants.OpAddition, operands)
	}
}

// BenchmarkCalculateFactorial benchmarks the factorial operation.
func BenchmarkCalculateFactorial(b *testing.B) {
	operands := []float64{20}
	for i := 0; i < b.N; i++ {
		Calculate(constants.OpFactorial, operands)
	}
}
