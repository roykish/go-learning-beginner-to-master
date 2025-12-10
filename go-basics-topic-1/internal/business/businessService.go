// Package businessService handles business logic and orchestrates the calculator operations.
// This demonstrates package organization, error handling, and function composition.
package businessService

import (
	"cli-calculator/internal/calculator"
	"cli-calculator/internal/config"
	"cli-calculator/internal/constants"
	"cli-calculator/internal/errors"
	"cli-calculator/internal/history"
	"cli-calculator/internal/logger"
	"cli-calculator/internal/util"
	"cli-calculator/internal/validation"
	"fmt"
)

// Service holds the application state and dependencies.
// This demonstrates struct composition and dependency injection.
type Service struct {
	Config  *config.Config  // Application configuration
	History *history.History // Calculation history
}

// NewService creates a new Service instance with loaded configuration and history.
// This demonstrates constructor functions and initialization.
func NewService() (*Service, error) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load configuration: %v", err)
		cfg = config.DefaultConfig() // Use defaults on error
	}

	// Initialize history
	var hist *history.History
	if cfg.HistoryPath != nil {
		hist = history.NewHistory(*cfg.HistoryPath, cfg.MaxHistory)
		if err := hist.Load(); err != nil {
			logger.Warn("Failed to load history: %v", err)
		}
	} else {
		hist = history.NewHistory("", cfg.MaxHistory)
	}

	return &Service{
		Config:  cfg,
		History: hist,
	}, nil
}

// Run starts the main application loop.
// This demonstrates control flow and menu-driven interfaces.
func (s *Service) Run() error {
	// Display welcome message if configured
	if s.Config.ShowWelcome {
		util.DisplayWelcome()
	}

	// Main loop
	for {
		util.DisplayMainMenu()

		input, err := util.GetUserInput("Enter your choice (1-7): ")
		if err != nil {
			return errors.Wrap(err, "failed to read menu input")
		}

		// Validate menu option
		option, err := validation.ValidateMenuOption(input)
		if err != nil {
			util.PrintError(err)
			continue
		}

		// Handle the menu option
		shouldExit, err := s.handleMenuOption(option)
		if err != nil {
			util.PrintError(err)
			util.PressEnterToContinue()
		}

		if shouldExit {
			return nil
		}
	}
}

// handleMenuOption processes a menu selection and returns whether to exit.
func (s *Service) handleMenuOption(option constants.MenuOption) (bool, error) {
	logger.Debug("Handling menu option: %d", option)

	switch option {
	case constants.MenuBasicCalculator:
		return false, s.handleBasicCalculator()
	case constants.MenuAdvancedCalculator:
		return false, s.handleAdvancedCalculator()
	case constants.MenuBatchCalculations:
		return false, s.handleBatchCalculations()
	case constants.MenuHistory:
		return false, s.handleHistory()
	case constants.MenuSettings:
		return false, s.handleSettings()
	case constants.MenuHelp:
		return false, s.handleHelp()
	case constants.MenuExit:
		return s.handleExit()
	default:
		return false, errors.NewValidationError("menu_option", fmt.Sprintf("%d", option), "invalid menu option")
	}
}

// handleBasicCalculator handles the basic calculator submenu.
func (s *Service) handleBasicCalculator() error {
	if s.Config.ClearScreen {
		util.ClearScreen()
	}

	util.DisplayBasicCalculatorMenu()

	for {
		input, err := util.GetUserInput("Enter operation (1-4) or 0 to go back: ")
		if err != nil {
			return err
		}

		// Check for back option
		if input == "0" {
			return nil
		}

		// Validate operation
		operation, err := validation.ValidateBasicOperation(input)
		if err != nil {
			util.PrintError(err)
			continue
		}

		// Perform calculation
		if err := s.performCalculation(operation); err != nil {
			util.PrintError(err)
		}

		util.PressEnterToContinue()
		return nil
	}
}

// handleAdvancedCalculator handles the advanced calculator submenu.
func (s *Service) handleAdvancedCalculator() error {
	if s.Config.ClearScreen {
		util.ClearScreen()
	}

	util.DisplayAdvancedCalculatorMenu()

	for {
		input, err := util.GetUserInput("Enter operation (1-4) or 0 to go back: ")
		if err != nil {
			return err
		}

		// Check for back option
		if input == "0" {
			return nil
		}

		// Validate operation
		operation, err := s.validateAdvancedOperation(input)
		if err != nil {
			util.PrintError(err)
			continue
		}

		// Perform calculation
		if err := s.performCalculation(operation); err != nil {
			util.PrintError(err)
		}

		util.PressEnterToContinue()
		return nil
	}
}

// validateAdvancedOperation validates advanced calculator input.
func (s *Service) validateAdvancedOperation(input string) (constants.Operation, error) {
	// Parse input
	num := 0
	_, err := fmt.Sscanf(input, "%d", &num)
	if err != nil {
		return 0, errors.NewValidationError("operation", input, "not a valid number")
	}

	// Map to operations
	operations := map[int]constants.Operation{
		1: constants.OpPower,
		2: constants.OpSquareRoot,
		3: constants.OpModulo,
		4: constants.OpFactorial,
	}

	op, ok := operations[num]
	if !ok {
		return 0, errors.NewValidationError("operation", input, "must be between 1 and 4")
	}

	return op, nil
}

// performCalculation performs a calculation and updates history.
func (s *Service) performCalculation(operation constants.Operation) error {
	// Get operands based on operation
	operands, err := s.getOperands(operation)
	if err != nil {
		return err
	}

	// Build expression string
	expression := s.buildExpression(operation, operands)

	// Perform calculation
	result, err := calculator.Calculate(operation, operands)
	if err != nil {
		// Record failure in history
		if s.Config.SaveHistory {
			s.History.AddError(operation.String(), expression, err)
		}
		return err
	}

	// Format result
	resultStr := calculator.FormatResult(result, s.Config.Precision)

	// Display result
	util.PrintResult(operation.String(), expression, resultStr)

	// Add to history
	if s.Config.SaveHistory {
		s.History.AddSuccess(operation.String(), expression, result)

		// Auto-save history if configured
		if s.Config.AutoSave {
			if err := s.History.Save(); err != nil {
				logger.Warn("Failed to save history: %v", err)
			}
		}
	}

	logger.Info("Calculation completed: %s = %s", expression, resultStr)
	return nil
}

// getOperands prompts for and collects operands based on operation type.
func (s *Service) getOperands(operation constants.Operation) ([]float64, error) {
	switch operation {
	case constants.OpSquareRoot, constants.OpFactorial:
		// Single operand operations
		num, err := s.readNumber("Enter number: ")
		if err != nil {
			return nil, err
		}
		return []float64{num}, nil
	default:
		// Binary operations
		a, err := s.readNumber("Enter first number: ")
		if err != nil {
			return nil, err
		}
		b, err := s.readNumber("Enter second number: ")
		if err != nil {
			return nil, err
		}
		return []float64{a, b}, nil
	}
}

// readNumber prompts for and validates a number input.
func (s *Service) readNumber(prompt string) (float64, error) {
	input, err := util.GetUserInput(prompt)
	if err != nil {
		return 0, err
	}

	return validation.ValidateNumber(input)
}

// buildExpression builds a human-readable expression string.
func (s *Service) buildExpression(operation constants.Operation, operands []float64) string {
	switch operation {
	case constants.OpSquareRoot:
		return fmt.Sprintf("√%.2f", operands[0])
	case constants.OpFactorial:
		return fmt.Sprintf("%.0f!", operands[0])
	case constants.OpAddition, constants.OpSubtraction, constants.OpMultiplication, constants.OpDivision, constants.OpPower, constants.OpModulo:
		if len(operands) >= 2 {
			return fmt.Sprintf("%.2f %s %.2f", operands[0], operation.Symbol(), operands[1])
		}
	}
	return fmt.Sprintf("%s(%v)", operation.String(), operands)
}

// handleBatchCalculations handles batch calculation mode (placeholder).
func (s *Service) handleBatchCalculations() error {
	util.PrintInfo("Batch calculations feature coming soon!")
	util.PressEnterToContinue()
	return nil
}

// handleHistory displays calculation history.
func (s *Service) handleHistory() error {
	if s.Config.ClearScreen {
		util.ClearScreen()
	}

	fmt.Println("CALCULATION HISTORY:")
	util.PrintDivider()

	entries := s.History.GetAll()
	if len(entries) == 0 {
		util.PrintInfo("No calculation history available.")
	} else {
		for i, entry := range entries {
			status := "✓"
			if !entry.Success {
				status = "✗"
			}
			fmt.Printf("%d. [%s] %s: %s = ", i+1, status, entry.Timestamp.Format("15:04:05"), entry.Expression)
			if entry.Success {
				fmt.Printf("%.2f\n", entry.Result)
			} else {
				fmt.Printf("Error: %s\n", entry.Error)
			}
		}

		// Display statistics
		stats := s.History.GetStatistics()
		fmt.Println()
		util.PrintDivider()
		fmt.Printf("Total: %d | Successful: %d | Failed: %d\n",
			stats.TotalCalculations, stats.SuccessfulCount, stats.FailedCount)
		if stats.MostUsedOperation != "" {
			fmt.Printf("Most used operation: %s\n", stats.MostUsedOperation)
		}
	}

	util.PrintDivider()
	util.PressEnterToContinue()
	return nil
}

// handleSettings handles the settings menu (placeholder).
func (s *Service) handleSettings() error {
	if s.Config.ClearScreen {
		util.ClearScreen()
	}

	fmt.Println("SETTINGS:")
	util.PrintDivider()
	fmt.Printf("1. Precision: %d decimal places\n", s.Config.Precision)
	fmt.Printf("2. Save History: %v\n", s.Config.SaveHistory)
	fmt.Printf("3. Auto-save: %v\n", s.Config.AutoSave)
	fmt.Printf("4. Clear Screen: %v\n", s.Config.ClearScreen)
	util.PrintDivider()
	util.PrintInfo("Settings modification feature coming soon!")
	util.PressEnterToContinue()
	return nil
}

// handleHelp displays help information.
func (s *Service) handleHelp() error {
	if s.Config.ClearScreen {
		util.ClearScreen()
	}

	util.DisplayHelp()
	util.PressEnterToContinue()
	return nil
}

// handleExit handles application exit.
func (s *Service) handleExit() (bool, error) {
	// Confirm exit if configured
	if s.Config.ConfirmExit {
		confirm, err := util.Confirm("Are you sure you want to exit?")
		if err != nil {
			return false, err
		}
		if !confirm {
			return false, nil
		}
	}

	// Save history if auto-save is enabled
	if s.Config.AutoSave && s.Config.SaveHistory {
		if err := s.History.Save(); err != nil {
			logger.Error("Failed to save history: %v", err)
		}
	}

	fmt.Println("\nThank you for using CLI Calculator!")
	return true, nil
}
