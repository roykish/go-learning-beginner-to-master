// Package errors provides custom error types and error handling patterns.
// This demonstrates proper error creation, wrapping, and sentinel errors in Go.
package errors

import (
	"errors"
	"fmt"
)

// Sentinel errors - predefined errors that can be compared using errors.Is()
var (
	ErrInvalidInput      = errors.New("invalid input provided")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrNegativeSquareRoot = errors.New("cannot calculate square root of negative number")
	ErrInvalidOperation  = errors.New("invalid operation")
	ErrOutOfRange        = errors.New("value out of allowed range")
	ErrFileNotFound      = errors.New("file not found")
	ErrFileReadFailed    = errors.New("failed to read file")
	ErrFileWriteFailed   = errors.New("failed to write file")
	ErrConfigInvalid     = errors.New("configuration is invalid")
	ErrHistoryFull       = errors.New("history is full")
)

// ValidationError represents an input validation error with context.
// This is a custom error type that implements the error interface.
type ValidationError struct {
	Field   string // The field that failed validation
	Value   string // The invalid value
	Message string // Human-readable error message
}

// Error implements the error interface for ValidationError.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for %s='%s': %s", e.Field, e.Value, e.Message)
}

// NewValidationError creates a new ValidationError with the given details.
func NewValidationError(field, value, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}

// CalculationError represents an error that occurred during calculation.
type CalculationError struct {
	Operation string  // The operation being performed
	Operands  []float64 // The operands involved
	Reason    string  // The reason for failure
	Err       error   // The underlying error (if any)
}

// Error implements the error interface for CalculationError.
func (e *CalculationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("calculation error in %s: %s (caused by: %v)", e.Operation, e.Reason, e.Err)
	}
	return fmt.Sprintf("calculation error in %s: %s", e.Operation, e.Reason)
}

// Unwrap returns the underlying error, allowing errors.Is and errors.As to work.
func (e *CalculationError) Unwrap() error {
	return e.Err
}

// NewCalculationError creates a new CalculationError.
func NewCalculationError(operation string, operands []float64, reason string, err error) *CalculationError {
	return &CalculationError{
		Operation: operation,
		Operands:  operands,
		Reason:    reason,
		Err:       err,
	}
}

// FileError represents a file operation error.
type FileError struct {
	Path      string // The file path
	Operation string // The operation being performed (read/write/delete)
	Err       error  // The underlying error
}

// Error implements the error interface for FileError.
func (e *FileError) Error() string {
	return fmt.Sprintf("file error during %s on '%s': %v", e.Operation, e.Path, e.Err)
}

// Unwrap returns the underlying error.
func (e *FileError) Unwrap() error {
	return e.Err
}

// NewFileError creates a new FileError.
func NewFileError(path, operation string, err error) *FileError {
	return &FileError{
		Path:      path,
		Operation: operation,
		Err:       err,
	}
}

// Wrap wraps an error with additional context using fmt.Errorf and %w verb.
// This is a helper function to demonstrate error wrapping.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// WrapWithContext wraps an error with formatted context.
func WrapWithContext(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	message := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %w", message, err)
}
