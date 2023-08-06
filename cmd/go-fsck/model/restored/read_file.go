package model

import (
	"encoding/json"
	"os"
)

// ReadFile loads the definitions from a json file
func ReadFile(inputPath string) ([]*Definition, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	var result []*Definition
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	for _, decl := range result {
		decl.Fill()
	}

	return result, nil
}
