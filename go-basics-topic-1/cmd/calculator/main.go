// Package main is the entry point for the CLI Calculator application.
// This demonstrates the main package, command-line flags, and program initialization.
package main

import (
	business "cli-calculator/internal/business"
	"cli-calculator/internal/constants"
	"cli-calculator/internal/logger"
	"flag"
	"fmt"
	"os"
)

// Command-line flags
// This demonstrates flag declaration and usage
var (
	flagVersion   = flag.Bool("version", false, "Show version information")
	flagHelp      = flag.Bool("help", false, "Show help information")
	flagVerbose   = flag.Bool("verbose", false, "Enable verbose logging (debug level)")
	flagNoColor   = flag.Bool("no-color", false, "Disable colored output")
	flagPrecision = flag.Int("precision", constants.DefaultPrecision, "Number of decimal places for results (0-15)")
)

// main is the entry point of the application.
// This demonstrates program initialization and error handling.
func main() {
	// Parse command-line flags
	flag.Parse()

	// Handle special flags
	if *flagVersion {
		showVersion()
		os.Exit(int(constants.ExitSuccess))
	}

	if *flagHelp {
		showHelp()
		os.Exit(int(constants.ExitSuccess))
	}

	// Configure logging based on flags
	if *flagVerbose {
		logger.SetLevel(constants.LogLevelDebug)
		logger.Info("Verbose logging enabled")
	}

	// Log application start
	logger.Info("Starting %s v%s", constants.AppName, constants.AppVersion)

	// Create and initialize the service
	service, err := business.NewService()
	if err != nil {
		logger.Error("Failed to initialize service: %v", err)
		fmt.Fprintf(os.Stderr, "Error: Failed to initialize application: %v\n", err)
		os.Exit(int(constants.ExitError))
	}

	// Apply command-line flag overrides to configuration
	if *flagPrecision != constants.DefaultPrecision {
		if *flagPrecision < 0 || *flagPrecision > 15 {
			logger.Error("Invalid precision value: %d (must be 0-15)", *flagPrecision)
			fmt.Fprintf(os.Stderr, "Error: Precision must be between 0 and 15\n")
			os.Exit(int(constants.ExitInvalidInput))
		}
		service.Config.Precision = *flagPrecision
		logger.Debug("Precision set to %d via command-line flag", *flagPrecision)
	}

	if *flagNoColor {
		service.Config.ColorOutput = false
		logger.Debug("Color output disabled via command-line flag")
	}

	// Run the application
	// This demonstrates proper error handling and exit codes
	if err := service.Run(); err != nil {
		logger.Error("Application error: %v", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(int(constants.ExitError))
	}

	// Successful exit
	logger.Info("Application terminated successfully")
	os.Exit(int(constants.ExitSuccess))
}

// showVersion displays version information.
func showVersion() {
	fmt.Printf("%s version %s\n", constants.AppName, constants.AppVersion)
	fmt.Println("A production-grade CLI calculator demonstrating Go best practices")
}

// showHelp displays help information.
func showHelp() {
	fmt.Printf("%s - A production-grade CLI calculator\n\n", constants.AppName)
	fmt.Println("USAGE:")
	fmt.Printf("  %s [options]\n\n", os.Args[0])
	fmt.Println("OPTIONS:")
	flag.PrintDefaults()
	fmt.Println("\nEXAMPLES:")
	fmt.Println("  Start calculator:")
	fmt.Printf("    %s\n\n", os.Args[0])
	fmt.Println("  Start with high precision:")
	fmt.Printf("    %s -precision 5\n\n", os.Args[0])
	fmt.Println("  Start with verbose logging:")
	fmt.Printf("    %s -verbose\n\n", os.Args[0])
	fmt.Println("\nFEATURES:")
	fmt.Println("  - Basic arithmetic operations (+, -, *, /)")
	fmt.Println("  - Advanced operations (power, square root, modulo, factorial)")
	fmt.Println("  - Calculation history with statistics")
	fmt.Println("  - Configurable settings with file persistence")
	fmt.Println("  - Comprehensive error handling")
	fmt.Println("  - Structured logging")
}
