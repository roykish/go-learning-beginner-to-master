// Package constants defines application-wide constants using Go's iota for enumerations.
// This demonstrates proper constant declaration and the iota identifier.
package constants

// ExitCode represents application exit status codes.
// Using iota to create enumerated constants starting from 0.
type ExitCode int

const (
	ExitSuccess ExitCode = iota // 0 - successful execution
	ExitError                   // 1 - general error
	ExitInvalidInput            // 2 - invalid user input
	ExitFileError               // 3 - file operation error
	ExitConfigError             // 4 - configuration error
)

// Operation represents calculator operation types.
type Operation uint8

const (
	OpUnknown Operation = iota
	OpAddition
	OpSubtraction
	OpMultiplication
	OpDivision
	OpPower
	OpSquareRoot
	OpModulo
	OpFactorial
)

// String returns the string representation of an operation.
func (o Operation) String() string {
	switch o {
	case OpAddition:
		return "Addition"
	case OpSubtraction:
		return "Subtraction"
	case OpMultiplication:
		return "Multiplication"
	case OpDivision:
		return "Division"
	case OpPower:
		return "Power"
	case OpSquareRoot:
		return "Square Root"
	case OpModulo:
		return "Modulo"
	case OpFactorial:
		return "Factorial"
	default:
		return "Unknown"
	}
}

// Symbol returns the mathematical symbol for the operation.
func (o Operation) Symbol() string {
	switch o {
	case OpAddition:
		return "+"
	case OpSubtraction:
		return "-"
	case OpMultiplication:
		return "*"
	case OpDivision:
		return "/"
	case OpPower:
		return "^"
	case OpSquareRoot:
		return "√"
	case OpModulo:
		return "%"
	case OpFactorial:
		return "!"
	default:
		return "?"
	}
}

// MenuOption represents main menu choices.
type MenuOption uint8

const (
	MenuBasicCalculator MenuOption = iota + 1 // Start from 1
	MenuAdvancedCalculator
	MenuBatchCalculations
	MenuHistory
	MenuSettings
	MenuHelp
	MenuExit
)

// LogLevel represents logging severity levels.
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// String returns the string representation of a log level.
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Application constants
const (
	AppName           = "CLI Calculator"
	AppVersion        = "1.0.0"
	ConfigFileName    = ".calculator_config.json"
	HistoryFileName   = ".calculator_history.json"
	MaxHistoryEntries = 100
	DefaultPrecision  = 2
)

// Validation constants
const (
	MinMenuOption       = 1
	MaxMenuOption       = 7
	MinBasicCalcOption  = 1
	MaxBasicCalcOption  = 4
	MaxNumberInputValue = 1e15  // Maximum safe number for calculations
	MinNumberInputValue = -1e15 // Minimum safe number for calculations
)
