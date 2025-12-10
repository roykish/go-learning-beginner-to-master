// Package util provides UI and input/output utility functions.
// This demonstrates string formatting, I/O operations, and user interaction.
package util

import (
	"bufio"
	"cli-calculator/internal/constants"
	"cli-calculator/internal/errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// DisplayWelcome displays the welcome banner.
// This demonstrates multi-line string output and formatting.
func DisplayWelcome() {
	fmt.Println("╔══════════════════════════════════════════════════════╗")
	fmt.Printf("║              %s v%s              ║\n", constants.AppName, constants.AppVersion)
	fmt.Println("╠══════════════════════════════════════════════════════╣")
	fmt.Println("║  A simple yet powerful command-line calculator       ║")
	fmt.Println("║  with support for basic and advanced operations      ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝")
	fmt.Println()
}

// DisplayMainMenu displays the main menu options.
func DisplayMainMenu() {
	fmt.Println("MAIN MENU:")
	fmt.Println("════════════════════════════════════════════════════════")
	fmt.Println("1. Basic Calculator (+, -, *, /)")
	fmt.Println("2. Advanced Calculator (^, √, %, !)")
	fmt.Println("3. Batch Calculations (multiple operations)")
	fmt.Println("4. Calculation History")
	fmt.Println("5. Settings")
	fmt.Println("6. Help & Instructions")
	fmt.Println("7. Exit")
	fmt.Println("════════════════════════════════════════════════════════")
}

// DisplayBasicCalculatorMenu displays the basic calculator menu.
func DisplayBasicCalculatorMenu() {
	fmt.Println("BASIC CALCULATOR MENU:")
	fmt.Println("════════════════════════════════════════════════════════")
	fmt.Println("Available Operations:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("0. Back to Main Menu")
	fmt.Println("════════════════════════════════════════════════════════")
}

// DisplayAdvancedCalculatorMenu displays the advanced calculator menu.
func DisplayAdvancedCalculatorMenu() {
	fmt.Println("ADVANCED CALCULATOR MENU:")
	fmt.Println("════════════════════════════════════════════════════════")
	fmt.Println("Available Operations:")
	fmt.Println("1. Power (x^y)")
	fmt.Println("2. Square Root (√x)")
	fmt.Println("3. Modulo (x % y)")
	fmt.Println("4. Factorial (x!)")
	fmt.Println("0. Back to Main Menu")
	fmt.Println("════════════════════════════════════════════════════════")
}

// DisplayHelp displays help information.
func DisplayHelp() {
	fmt.Println("HELP & INSTRUCTIONS:")
	fmt.Println("════════════════════════════════════════════════════════")
	fmt.Println("BASIC OPERATIONS:")
	fmt.Println("  Addition       : Adds two or more numbers")
	fmt.Println("  Subtraction    : Subtracts second number from first")
	fmt.Println("  Multiplication : Multiplies two or more numbers")
	fmt.Println("  Division       : Divides first number by second")
	fmt.Println()
	fmt.Println("ADVANCED OPERATIONS:")
	fmt.Println("  Power          : Raises first number to power of second")
	fmt.Println("  Square Root    : Calculates square root of a number")
	fmt.Println("  Modulo         : Calculates remainder of division")
	fmt.Println("  Factorial      : Calculates factorial (n!)")
	fmt.Println()
	fmt.Println("FEATURES:")
	fmt.Println("  - History tracking of all calculations")
	fmt.Println("  - Configurable precision for results")
	fmt.Println("  - Persistent settings saved to disk")
	fmt.Println("  - Error handling with detailed messages")
	fmt.Println("════════════════════════════════════════════════════════")
}

// ClearScreen clears the terminal screen.
// This demonstrates platform-specific behavior.
func ClearScreen() {
	// ANSI escape sequence works on Unix-like systems and Windows 10+
	if runtime.GOOS == "windows" {
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

// GetUserInput prompts the user and reads a line of input.
// This demonstrates I/O operations and error handling.
func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Wrap(err, "failed to read input")
	}

	// Trim whitespace and handle Windows line endings
	input = strings.TrimSpace(input)
	input = strings.TrimSuffix(input, "\r")

	return input, nil
}

// Confirm asks the user a yes/no question.
// This demonstrates boolean return values and user interaction.
func Confirm(prompt string) (bool, error) {
	input, err := GetUserInput(prompt + " (y/n): ")
	if err != nil {
		return false, err
	}

	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "yes", nil
}

// PrintSuccess prints a success message.
func PrintSuccess(message string) {
	fmt.Printf("✓ %s\n", message)
}

// PrintError prints an error message.
func PrintError(err error) {
	fmt.Printf("✗ Error: %v\n", err)
}

// PrintWarning prints a warning message.
func PrintWarning(message string) {
	fmt.Printf("⚠ Warning: %s\n", message)
}

// PrintInfo prints an informational message.
func PrintInfo(message string) {
	fmt.Printf("ℹ %s\n", message)
}

// PrintDivider prints a horizontal divider line.
func PrintDivider() {
	fmt.Println("════════════════════════════════════════════════════════")
}

// PrintResult prints a formatted calculation result.
func PrintResult(operation string, expression string, result string) {
	fmt.Println()
	PrintDivider()
	fmt.Printf("Operation : %s\n", operation)
	fmt.Printf("Expression: %s\n", expression)
	fmt.Printf("Result    : %s\n", result)
	PrintDivider()
	fmt.Println()
}

// PressEnterToContinue waits for the user to press Enter.
func PressEnterToContinue() {
	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
