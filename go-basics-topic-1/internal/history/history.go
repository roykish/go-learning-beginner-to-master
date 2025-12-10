// Package history manages calculation history with file persistence.
// This demonstrates slices, file I/O, time handling, and data structures.
package history

import (
	"cli-calculator/internal/errors"
	"encoding/json"
	"os"
	"time"
)

// Entry represents a single calculation history entry.
// This demonstrates struct tags for JSON serialization.
type Entry struct {
	Timestamp time.Time `json:"timestamp"` // When the calculation was performed
	Operation string    `json:"operation"` // The operation performed (e.g., "Addition")
	Expression string   `json:"expression"` // The full expression (e.g., "10 + 5")
	Result    float64   `json:"result"`    // The result of the calculation
	Success   bool      `json:"success"`   // Whether the calculation succeeded
	Error     string    `json:"error,omitempty"` // Error message if failed
}

// History manages a collection of calculation entries.
// This demonstrates slice usage and methods on structs.
type History struct {
	Entries  []Entry `json:"entries"`  // Slice of history entries
	MaxSize  int     `json:"max_size"` // Maximum number of entries to keep
	FilePath string  `json:"-"`        // Path to history file (not saved in JSON)
}

// NewHistory creates a new History instance with the given parameters.
func NewHistory(filePath string, maxSize int) *History {
	return &History{
		Entries:  make([]Entry, 0, maxSize), // Pre-allocate slice capacity
		MaxSize:  maxSize,
		FilePath: filePath,
	}
}

// Add adds a new entry to the history.
// This demonstrates slice append and capacity management.
func (h *History) Add(entry Entry) {
	// Add timestamp if not set
	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now()
	}

	// Append to slice
	h.Entries = append(h.Entries, entry)

	// Trim if exceeds max size (keep most recent entries)
	if len(h.Entries) > h.MaxSize {
		// Remove oldest entries
		excess := len(h.Entries) - h.MaxSize
		h.Entries = h.Entries[excess:]
	}
}

// AddSuccess adds a successful calculation to history.
func (h *History) AddSuccess(operation, expression string, result float64) {
	h.Add(Entry{
		Operation:  operation,
		Expression: expression,
		Result:     result,
		Success:    true,
	})
}

// AddError adds a failed calculation to history.
func (h *History) AddError(operation, expression string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}

	h.Add(Entry{
		Operation:  operation,
		Expression: expression,
		Success:    false,
		Error:      errorMsg,
	})
}

// GetRecent returns the most recent n entries.
// This demonstrates slice slicing and bounds checking.
func (h *History) GetRecent(n int) []Entry {
	if n <= 0 {
		return []Entry{}
	}

	if n >= len(h.Entries) {
		return h.Entries
	}

	// Return last n entries
	return h.Entries[len(h.Entries)-n:]
}

// GetAll returns all history entries.
func (h *History) GetAll() []Entry {
	return h.Entries
}

// Count returns the number of entries in history.
func (h *History) Count() int {
	return len(h.Entries)
}

// Clear removes all entries from history.
func (h *History) Clear() {
	h.Entries = make([]Entry, 0, h.MaxSize)
}

// Load loads history from the file.
// This demonstrates file reading and JSON unmarshaling with error handling.
func (h *History) Load() error {
	// Check if file exists
	if _, err := os.Stat(h.FilePath); os.IsNotExist(err) {
		// File doesn't exist, start with empty history (not an error)
		return nil
	}

	// Read file
	data, err := os.ReadFile(h.FilePath)
	if err != nil {
		return errors.NewFileError(h.FilePath, "read", err)
	}

	// Unmarshal JSON
	var loaded History
	if err := json.Unmarshal(data, &loaded); err != nil {
		return errors.WrapWithContext(err, "failed to parse history file")
	}

	// Update entries (preserve FilePath and MaxSize)
	h.Entries = loaded.Entries

	// Trim if loaded history exceeds current max size
	if len(h.Entries) > h.MaxSize {
		excess := len(h.Entries) - h.MaxSize
		h.Entries = h.Entries[excess:]
	}

	return nil
}

// Save saves history to the file.
// This demonstrates JSON marshaling and file writing with error handling.
func (h *History) Save() error {
	// Marshal to JSON with indentation
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return errors.WrapWithContext(err, "failed to marshal history")
	}

	// Write to file
	if err := os.WriteFile(h.FilePath, data, 0644); err != nil {
		return errors.NewFileError(h.FilePath, "write", err)
	}

	return nil
}

// GetStatistics calculates statistics from history.
// This demonstrates iteration, conditionals, and working with slices.
type Statistics struct {
	TotalCalculations   int
	SuccessfulCount     int
	FailedCount         int
	MostUsedOperation   string
	AverageResult       float64
	FirstCalculation    *time.Time
	LastCalculation     *time.Time
}

// GetStatistics returns statistics about the calculation history.
func (h *History) GetStatistics() Statistics {
	stats := Statistics{
		TotalCalculations: len(h.Entries),
	}

	if len(h.Entries) == 0 {
		return stats
	}

	// Track operation counts
	operationCounts := make(map[string]int)
	var totalResult float64
	var successfulResults int

	// Iterate through entries
	for i := range h.Entries {
		entry := &h.Entries[i] // Use pointer to avoid copying

		// Count success/failure
		if entry.Success {
			stats.SuccessfulCount++
			totalResult += entry.Result
			successfulResults++
		} else {
			stats.FailedCount++
		}

		// Count operations
		operationCounts[entry.Operation]++

		// Track first and last calculation times
		if stats.FirstCalculation == nil || entry.Timestamp.Before(*stats.FirstCalculation) {
			t := entry.Timestamp
			stats.FirstCalculation = &t
		}
		if stats.LastCalculation == nil || entry.Timestamp.After(*stats.LastCalculation) {
			t := entry.Timestamp
			stats.LastCalculation = &t
		}
	}

	// Calculate average result
	if successfulResults > 0 {
		stats.AverageResult = totalResult / float64(successfulResults)
	}

	// Find most used operation
	maxCount := 0
	for op, count := range operationCounts {
		if count > maxCount {
			maxCount = count
			stats.MostUsedOperation = op
		}
	}

	return stats
}

// Filter returns entries matching a predicate function.
// This demonstrates function parameters and filtering.
func (h *History) Filter(predicate func(Entry) bool) []Entry {
	filtered := make([]Entry, 0)

	for _, entry := range h.Entries {
		if predicate(entry) {
			filtered = append(filtered, entry)
		}
	}

	return filtered
}

// GetSuccessful returns only successful calculations.
func (h *History) GetSuccessful() []Entry {
	return h.Filter(func(e Entry) bool {
		return e.Success
	})
}

// GetFailed returns only failed calculations.
func (h *History) GetFailed() []Entry {
	return h.Filter(func(e Entry) bool {
		return !e.Success
	})
}
