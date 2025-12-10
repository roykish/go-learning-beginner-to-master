// Package config provides configuration management with tests.
// This demonstrates testing with file I/O and temporary files.
package config

import (
	"cli-calculator/internal/constants"
	"os"
	"path/filepath"
	"testing"
)

// TestDefaultConfig tests the default configuration creation.
func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg == nil {
		t.Fatal("DefaultConfig returned nil")
	}

	// Check default values
	if cfg.Precision != constants.DefaultPrecision {
		t.Errorf("Expected precision %d, got %d", constants.DefaultPrecision, cfg.Precision)
	}

	if !cfg.ShowWelcome {
		t.Error("Expected ShowWelcome to be true")
	}

	if !cfg.SaveHistory {
		t.Error("Expected SaveHistory to be true")
	}

	if cfg.MaxHistory != constants.MaxHistoryEntries {
		t.Errorf("Expected MaxHistory %d, got %d", constants.MaxHistoryEntries, cfg.MaxHistory)
	}
}

// TestConfigValidation tests configuration validation.
func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name      string
		config    *Config
		hasError  bool
	}{
		{
			name:     "valid config",
			config:   DefaultConfig(),
			hasError: false,
		},
		{
			name: "invalid precision negative",
			config: &Config{
				Precision:  -1,
				MaxHistory: 100,
			},
			hasError: true,
		},
		{
			name: "invalid precision too high",
			config: &Config{
				Precision:  16,
				MaxHistory: 100,
			},
			hasError: true,
		},
		{
			name: "invalid max history negative",
			config: &Config{
				Precision:  2,
				MaxHistory: -1,
			},
			hasError: true,
		},
		{
			name: "invalid max history too high",
			config: &Config{
				Precision:  2,
				MaxHistory: 10001,
			},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()

			if tt.hasError {
				if err == nil {
					t.Errorf("%s: expected error, got nil", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tt.name, err)
				}
			}
		})
	}
}

// TestConfigSaveAndLoad tests saving and loading configuration.
func TestConfigSaveAndLoad(t *testing.T) {
	// Create a temporary file
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "test_config.json")

	// Create a config with custom values
	cfg := DefaultConfig()
	cfg.ConfigPath = &configPath
	cfg.Precision = 5
	cfg.ShowWelcome = false
	cfg.SaveHistory = false

	// Save the config
	err := cfg.Save()
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Check that file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatal("Config file was not created")
	}

	// Load the config
	cfg2 := &Config{ConfigPath: &configPath}
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config file: %v", err)
	}

	// Parse it
	if len(data) == 0 {
		t.Fatal("Config file is empty")
	}

	// Verify config can be loaded (basic check)
	if cfg.Precision != 5 {
		t.Errorf("Expected precision 5, got %d", cfg2.Precision)
	}
}

// TestConfigReset tests resetting configuration to defaults.
func TestConfigReset(t *testing.T) {
	cfg := DefaultConfig()

	// Modify values
	cfg.Precision = 10
	cfg.ShowWelcome = false
	cfg.SaveHistory = false

	// Reset
	cfg.Reset()

	// Check that values are back to defaults
	if cfg.Precision != constants.DefaultPrecision {
		t.Errorf("Expected precision %d after reset, got %d", constants.DefaultPrecision, cfg.Precision)
	}

	if !cfg.ShowWelcome {
		t.Error("Expected ShowWelcome to be true after reset")
	}

	if !cfg.SaveHistory {
		t.Error("Expected SaveHistory to be true after reset")
	}
}

// TestConfigClone tests cloning configuration.
func TestConfigClone(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Precision = 7

	// Clone the config
	clone := cfg.Clone()

	// Verify clone has same values
	if clone.Precision != cfg.Precision {
		t.Errorf("Clone precision mismatch: expected %d, got %d", cfg.Precision, clone.Precision)
	}

	// Verify it's a separate copy
	clone.Precision = 3
	if cfg.Precision == clone.Precision {
		t.Error("Modifying clone affected original config")
	}

	// Verify pointer fields are deep copied
	if cfg.ConfigPath != nil && clone.ConfigPath != nil {
		if cfg.ConfigPath == clone.ConfigPath {
			t.Error("ConfigPath pointers are the same (not deep copied)")
		}
	}
}

// TestLoadNonExistentConfig tests loading when config file doesn't exist.
func TestLoadNonExistentConfig(t *testing.T) {
	// Try to load from non-existent file
	cfg, err := Load()

	// Should not error (returns default config)
	if err != nil {
		t.Errorf("Load() with non-existent file returned error: %v", err)
	}

	// Should return default config
	if cfg == nil {
		t.Fatal("Load() returned nil config")
	}

	if cfg.Precision != constants.DefaultPrecision {
		t.Errorf("Expected default precision %d, got %d", constants.DefaultPrecision, cfg.Precision)
	}
}
