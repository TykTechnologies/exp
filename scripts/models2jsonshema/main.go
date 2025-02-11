package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"models2jsonshema/converter"

	"github.com/TykTechnologies/exp/cmd/schema-gen/extract"
)

func main() {
	pkgInfos, err := extract.Extract("/Users/itachisasuke/projects/dc/tyk/config/.", &extract.ExtractOptions{})
	if err != nil {
		log.Fatalf("Failed to extract types: %v", err)
	}
	rootType := "Config"

	schema, err := converter.ConvertToJSONSchema(pkgInfos[0], rootType, NewDefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal schema: %v", err)
	}
	err = os.WriteFile("schema.json", jsonBytes, 0o644)
	if err != nil {
		log.Fatalf("Failed to write schema: %v", err)
	}

	fmt.Println("Successfully generated JSON Schema in schema.json")
}

func NewDefaultConfig() *converter.RequiredFieldsConfig {
	return &converter.RequiredFieldsConfig{
		Fields: map[string][]string{
			"User":  {"ID", "Name"}, // Only ID and Name are required for User
			"Inner": {"Name"},
		},
	}
}
