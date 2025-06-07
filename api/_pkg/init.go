package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Pals is the global slice containing all Pal data
var Pals []Pal

// init automatically runs when the package is imported
func init() {
	if err := Initialize(); err != nil {
		fmt.Printf("Failed to initialize pals data: %v\n", err)
	}
}

// Initialize loads the pals data from the JSON file
func Initialize() error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	// Build full path to JSON file
	jsonPath := filepath.Join(cwd, "api", "pals.json")
	// Read and parse the JSON file
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("error reading pals.json: %v", err)
	}

	err = json.Unmarshal(data, &Pals)
	if err != nil {
		return fmt.Errorf("error parsing pals.json: %v", err)
	}

	return nil
} 