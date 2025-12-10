// Package config provides configuration management with file I/O operations.
// This demonstrates JSON marshaling/unmarshaling, file I/O, and pointer usage.
package config

import (
	"cli-calculator/internal/constants"
	"cli-calculator/internal/errors"
	"encoding/json"
	"os"
	"path/filepath"
)

// Config represents the application configuration.
// Using pointers for optional fields allows distinguishing between zero values and unset values.
type Config struct {
	// Display settings
	Precision       int  `json:"precision"`        // Number of decimal places
	ShowWelcome     bool `json:"show_welcome"`     // Show welcome message
	ClearScreen     bool `json:"clear_screen"`     // Clear screen between operations
	ColorOutput     bool `json:"color_output"`     // Enable colored output

	// Behavior settings
	SaveHistory     bool `json:"save_history"`     // Save calculation history
	MaxHistory      int  `json:"max_history"`      // Maximum history entries
	AutoSave        bool `json:"auto_save"`        // Auto-save config changes
	ConfirmExit     bool `json:"confirm_exit"`     // Ask confirmation before exit

	// Advanced settings
	UseRadians      bool    `json:"use_radians"`      // Use radians for trig (for future)
	ScientificMode  bool    `json:"scientific_mode"`  // Enable scientific notation
	ThousandSep     bool    `json:"thousand_sep"`     // Use thousand separator

	// File paths (using pointers to show optional string fields)
	ConfigPath  *string `json:"-"` // Path to config file (not saved in JSON)
	HistoryPath *string `json:"-"` // Path to history file (not saved in JSON)
}

// DefaultConfig returns a configuration with default values.
// This demonstrates function returning a pointer to a struct.
func DefaultConfig() *Config {
	// Get user's home directory for storing config files
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "." // Fallback to current directory
	}

	configPath := filepath.Join(homeDir, constants.ConfigFileName)
	historyPath := filepath.Join(homeDir, constants.HistoryFileName)

	return &Config{
		Precision:      constants.DefaultPrecision,
		ShowWelcome:    true,
		ClearScreen:    true,
		ColorOutput:    false,
		SaveHistory:    true,
		MaxHistory:     constants.MaxHistoryEntries,
		AutoSave:       true,
		ConfirmExit:    false,
		UseRadians:     false,
		ScientificMode: false,
		ThousandSep:    false,
		ConfigPath:     &configPath,
		HistoryPath:    &historyPath,
	}
}

// Load loads configuration from the config file.
// If the file doesn't exist, it returns the default configuration.
// This demonstrates file reading and error handling.
func Load() (*Config, error) {
	config := DefaultConfig()

	// Check if config file exists
	if config.ConfigPath == nil {
		return config, nil
	}

	data, err := os.ReadFile(*config.ConfigPath)
	if err != nil {
		// If file doesn't exist, return default config (not an error)
		if os.IsNotExist(err) {
			return config, nil
		}
		return nil, errors.NewFileError(*config.ConfigPath, "read", err)
	}

	// Unmarshal JSON into config struct
	if err := json.Unmarshal(data, config); err != nil {
		return nil, errors.WrapWithContext(err, "failed to parse config file")
	}

	// Restore paths (they're not saved in JSON)
	configPath := *config.ConfigPath
	historyPath := *config.HistoryPath
	config.ConfigPath = &configPath
	config.HistoryPath = &historyPath

	return config, nil
}

// Save saves the configuration to the config file.
// This demonstrates JSON marshaling and file writing.
func (c *Config) Save() error {
	if c.ConfigPath == nil {
		return errors.Wrap(errors.ErrConfigInvalid, "config path is nil")
	}

	// Marshal config to JSON with indentation for readability
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return errors.WrapWithContext(err, "failed to marshal config")
	}

	// Write to file with appropriate permissions (0644 = rw-r--r--)
	if err := os.WriteFile(*c.ConfigPath, data, 0644); err != nil {
		return errors.NewFileError(*c.ConfigPath, "write", err)
	}

	return nil
}

// Validate validates the configuration values.
// This demonstrates validation logic and error handling.
func (c *Config) Validate() error {
	// Validate precision
	if c.Precision < 0 || c.Precision > 15 {
		return errors.NewValidationError("precision", string(rune(c.Precision)), "must be between 0 and 15")
	}

	// Validate max history
	if c.MaxHistory < 0 || c.MaxHistory > 10000 {
		return errors.NewValidationError("max_history", string(rune(c.MaxHistory)), "must be between 0 and 10000")
	}

	return nil
}

// Reset resets the configuration to default values.
func (c *Config) Reset() {
	defaultCfg := DefaultConfig()

	// Preserve file paths
	configPath := c.ConfigPath
	historyPath := c.HistoryPath

	// Copy all values from default
	*c = *defaultCfg

	// Restore paths
	c.ConfigPath = configPath
	c.HistoryPath = historyPath
}

// Clone creates a deep copy of the configuration.
// This demonstrates pointer handling and value copying.
func (c *Config) Clone() *Config {
	clone := *c // Copy all fields

	// Deep copy pointer fields
	if c.ConfigPath != nil {
		path := *c.ConfigPath
		clone.ConfigPath = &path
	}
	if c.HistoryPath != nil {
		path := *c.HistoryPath
		clone.HistoryPath = &path
	}

	return &clone
}
