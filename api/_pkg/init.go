package pkg

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Pals is the global slice containing all Pal data
var Pals []Pal

//go:embed "pals.json"
var palsJSON []byte

// init automatically runs when the package is imported
func init() {
	if err := Initialize(); err != nil {
		fmt.Printf("Failed to initialize pals data: %v\n", err)
	}
}

// Initialize loads the pals data from the embedded JSON
func Initialize() error {
	err := json.Unmarshal(palsJSON, &Pals)
	if err != nil {
		return fmt.Errorf("error parsing pals.json: %v", err)
	}
	return nil
}
