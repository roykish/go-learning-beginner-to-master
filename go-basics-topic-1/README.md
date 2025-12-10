# Go Basics - Topic 1: Production-Grade CLI Calculator

A comprehensive, production-grade command-line calculator built to demonstrate Go fundamentals and best practices.

## Table of Contents

- [Overview](#overview)
- [Learning Objectives](#learning-objectives)
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Key Concepts Demonstrated](#key-concepts-demonstrated)
- [Testing](#testing)
- [Architecture](#architecture)

## Overview

This project implements a fully-featured CLI calculator that covers all fundamental Go concepts in a production-ready manner. It goes beyond basic arithmetic to demonstrate industry-standard practices for Go applications.

## Learning Objectives

This project covers all the learning objectives from the main README:

### 1. Go Modules & Dependency Management
- ✅ Module initialization (`go mod init`)
- ✅ Module structure and organization
- ✅ Internal package visibility

### 2. Package Structure & Organization
- ✅ Multi-package architecture
- ✅ `internal/` directory for private packages
- ✅ `cmd/` directory for entry points
- ✅ Clear separation of concerns

### 3. Main Package & Entry Point
- ✅ Proper `main` function structure
- ✅ Command-line flag parsing
- ✅ Application initialization
- ✅ Exit codes and graceful shutdown

### 4. Variables, Constants, and Data Types
- ✅ Variable declarations (short, full, package-level)
- ✅ Constants and `const` blocks
- ✅ Enumerated constants using `iota`
- ✅ Type definitions and custom types
- ✅ Basic types: `int`, `float64`, `string`, `bool`
- ✅ Composite types: structs, slices, maps

### 5. Control Flow
- ✅ `if/else` conditionals
- ✅ `switch` statements (value and type switches)
- ✅ `for` loops (traditional, range, infinite)
- ✅ Loop control (`break`, `continue`)

### 6. Functions
- ✅ Function parameters and return values
- ✅ Multiple return values
- ✅ Named return values
- ✅ Variadic functions
- ✅ Function types and closures
- ✅ Methods on types

### 7. Error Handling
- ✅ Error interface implementation
- ✅ Custom error types
- ✅ Error wrapping with `fmt.Errorf` and `%w`
- ✅ Sentinel errors
- ✅ Error unwrapping with `errors.Is` and `errors.As`
- ✅ Panic and recovery (where appropriate)

### 8. Pointers
- ✅ Pointer declaration and dereferencing
- ✅ Passing by value vs. reference
- ✅ Pointer receivers on methods
- ✅ Nil pointer checks
- ✅ Pointers in structs (optional fields)

## Features

### Core Functionality
- **Basic Operations**: Addition, subtraction, multiplication, division
- **Advanced Operations**: Power, square root, modulo, factorial
- **History Tracking**: Persistent calculation history with statistics
- **Configuration**: User preferences saved to disk
- **Precision Control**: Configurable decimal places (0-15)

### Production Features
- **Structured Logging**: Multiple log levels (Debug, Info, Warn, Error)
- **Error Handling**: Comprehensive error types and graceful degradation
- **Input Validation**: Robust validation with helpful error messages
- **File I/O**: JSON-based config and history persistence
- **Command-line Flags**: Full CLI argument support
- **Exit Codes**: Proper exit status codes
- **Unit Tests**: Comprehensive test coverage
- **Documentation**: Full GoDoc comments

## Project Structure

```
go-basics-topic-1/
├── cmd/
│   └── calculator/
│       └── main.go              # Application entry point with CLI flags
├── internal/
│   ├── business/
│   │   └── businessService.go   # Business logic orchestration
│   ├── calculator/
│   │   ├── calculator.go        # Core calculation functions
│   │   └── calculator_test.go   # Unit tests
│   ├── config/
│   │   ├── config.go            # Configuration management & file I/O
│   │   └── config_test.go       # Configuration tests
│   ├── constants/
│   │   └── constants.go         # Application constants with iota
│   ├── errors/
│   │   └── errors.go            # Custom error types
│   ├── history/
│   │   └── history.go           # Calculation history with persistence
│   ├── logger/
│   │   └── logger.go            # Structured logging
│   ├── util/
│   │   └── utility.go           # UI and I/O utilities
│   └── validation/
│       ├── validation.go        # Input validation
│       └── validation_test.go   # Validation tests
├── go.mod                       # Module definition
└── README.md                    # This file
```

## Installation

### Prerequisites
- Go 1.20 or higher
- Terminal/Command Prompt

### Build from Source

```bash
# Navigate to project directory
cd go-basics-topic-1

# Build the application
go build -o bin/calculator cmd/calculator/main.go

# Or run directly
go run cmd/calculator/main.go
```

## Usage

### Basic Usage

```bash
# Run the calculator
./bin/calculator

# Or with go run
go run cmd/calculator/main.go
```

### Command-line Flags

```bash
# Show version
./bin/calculator -version

# Show help
./bin/calculator -help

# Set precision
./bin/calculator -precision 5

# Enable verbose logging
./bin/calculator -verbose

# Disable colored output
./bin/calculator -no-color
```

### Interactive Menu

Once running, you'll see a menu with options:

1. **Basic Calculator** - Arithmetic operations (+, -, *, /)
2. **Advanced Calculator** - Power, square root, modulo, factorial
3. **Batch Calculations** - (Coming soon)
4. **Calculation History** - View past calculations with statistics
5. **Settings** - View current configuration
6. **Help & Instructions** - Detailed help information
7. **Exit** - Quit the application

## Key Concepts Demonstrated

### 1. Constants and Enumerations (`internal/constants/`)

```go
// Using iota for enumerated constants
type Operation uint8

const (
    OpUnknown Operation = iota
    OpAddition
    OpSubtraction
    OpMultiplication
    OpDivision
)
```

**Learn:** How to use `iota` for auto-incrementing constants and creating type-safe enumerations.

### 2. Custom Error Types (`internal/errors/`)

```go
// Custom error type with context
type ValidationError struct {
    Field   string
    Value   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error for %s='%s': %s",
        e.Field, e.Value, e.Message)
}
```

**Learn:** How to create custom error types that implement the `error` interface and provide rich context.

### 3. Pointers in Structs (`internal/config/`)

```go
type Config struct {
    Precision  int
    ConfigPath *string  // Pointer allows nil value
}

func (c *Config) Save() error {
    // Method with pointer receiver
    if c.ConfigPath == nil {
        return errors.New("config path is nil")
    }
    // ...
}
```

**Learn:** When and why to use pointers, including optional fields and method receivers.

### 4. File I/O (`internal/config/`, `internal/history/`)

```go
// JSON marshaling and file writing
data, err := json.MarshalIndent(config, "", "  ")
if err != nil {
    return err
}
return os.WriteFile(path, data, 0644)
```

**Learn:** Reading and writing files, JSON serialization, and error handling with I/O operations.

### 5. Table-Driven Tests (`*_test.go`)

```go
func TestCalculateFactorial(t *testing.T) {
    tests := []struct {
        name     string
        input    float64
        expected float64
        hasError bool
    }{
        {"factorial of 5", 5, 120, false},
        {"factorial of negative", -5, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

**Learn:** How to write idiomatic Go tests using table-driven test patterns.

### 6. Error Wrapping (`internal/errors/`)

```go
// Error wrapping with %w for error chains
if err != nil {
    return fmt.Errorf("failed to read config: %w", err)
}

// Checking wrapped errors
if errors.Is(err, ErrDivisionByZero) {
    // Handle specific error
}
```

**Learn:** Modern Go error handling with wrapping and unwrapping.

### 7. Slices and Maps (`internal/history/`)

```go
// Slice operations
h.Entries = append(h.Entries, entry)
if len(h.Entries) > h.MaxSize {
    h.Entries = h.Entries[excess:]  // Slice slicing
}

// Map usage
operationCounts := make(map[string]int)
operationCounts[entry.Operation]++
```

**Learn:** Working with slices (append, capacity, slicing) and maps (initialization, access, iteration).

### 8. Method Receivers (`internal/logger/`)

```go
// Value receiver (doesn't modify original)
func (o Operation) String() string {
    // ...
}

// Pointer receiver (can modify original)
func (l *Logger) SetLevel(level LogLevel) {
    l.config.Level = level
}
```

**Learn:** Difference between value and pointer receivers and when to use each.

## Testing

### Run All Tests

```bash
cd go-basics-topic-1
go test ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

### Run Specific Package Tests

```bash
go test ./internal/calculator
go test ./internal/validation
go test ./internal/config
```

### Run Benchmarks

```bash
go test -bench=. ./internal/calculator
```

### Verbose Test Output

```bash
go test -v ./...
```

## Architecture

### Layered Architecture

```
┌─────────────────────────────────┐
│  Presentation Layer (main.go)  │  ← CLI, flags, user interaction
├─────────────────────────────────┤
│  Business Layer (business/)     │  ← Orchestration, workflows
├─────────────────────────────────┤
│  Domain Layer (calculator/)     │  ← Core logic, calculations
├─────────────────────────────────┤
│  Infrastructure Layer           │  ← Config, logging, history
│  (config/, logger/, history/)   │
├─────────────────────────────────┤
│  Utility Layer                  │  ← Validation, utilities, errors
│  (validation/, util/, errors/)  │
└─────────────────────────────────┘
```

### Design Principles Applied

1. **Separation of Concerns**: Each package has a single, well-defined responsibility
2. **Dependency Injection**: Dependencies passed explicitly rather than global state
3. **Interface Segregation**: Small, focused interfaces
4. **Error Handling**: Errors bubble up with context
5. **Testability**: Pure functions and mockable dependencies

## What Makes This Production-Grade?

1. ✅ **Proper Error Handling**: Custom error types, error wrapping, graceful degradation
2. ✅ **Logging**: Structured logging with configurable levels
3. ✅ **Configuration**: Persistent settings with validation
4. ✅ **Testing**: Comprehensive unit tests and benchmarks
5. ✅ **Documentation**: Full GoDoc comments on all exported items
6. ✅ **Input Validation**: Robust validation with clear error messages
7. ✅ **Exit Codes**: Proper Unix exit codes for different failure modes
8. ✅ **File I/O**: Safe file operations with error handling
9. ✅ **User Experience**: Clear UI, helpful messages, history tracking
10. ✅ **Code Organization**: Clean package structure following Go conventions

## Next Steps

After mastering this project, you'll be ready to:

- Add concurrent operations (goroutines and channels)
- Implement interfaces and polymorphism
- Add HTTP API endpoints
- Connect to databases
- Work with external packages
- Build microservices

## License

This project is part of a learning repository. Feel free to use and modify for educational purposes.
