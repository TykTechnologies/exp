package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"

	"github.com/TykTechnologies/exp/scripts/models2jsonshema/converter"
)

func main() {
	var (
		userDir  string
		rootType string
		outFile  string
	)

	pflag.StringVar(&userDir, "dir", "", "Path to the user's Go module (required)")
	pflag.StringVar(&rootType, "type", "", "Root type to generate schema for (required)")
	pflag.StringVar(&outFile, "out", "schema.json", "Output file name (optional)")

	pflag.Parse()

	// Check required flags
	if userDir == "" {
		fmt.Fprintln(os.Stderr, "Error: --dir is required")
		pflag.Usage()
		os.Exit(1)
	}
	if rootType == "" {
		fmt.Fprintln(os.Stderr, "Error: --type is required")
		pflag.Usage()
		os.Exit(1)
	}

	if err := converter.ParseAndConvertStruct(userDir, rootType, outFile); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
