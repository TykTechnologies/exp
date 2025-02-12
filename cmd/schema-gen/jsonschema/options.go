package jsonschema

import (
	"fmt"

	"github.com/spf13/pflag"
)

type options struct {
	sourcePath string
	rootType   string
	outputFile string
}

func NewOptions() *options {
	cfg := &options{
		sourcePath: ".",
		outputFile: "schema.json",
	}

	pflag.StringVarP(&cfg.sourcePath, "dir", "i", cfg.sourcePath, "Path to the directory that contains the root type (required)")
	pflag.StringVarP(&cfg.rootType, "type", "t", cfg.rootType, "Root type to generate schema for (required)")
	pflag.StringVarP(&cfg.outputFile, "out", "o", cfg.outputFile, "Output file name (optional)")

	pflag.Parse()

	return cfg
}

// PrintHelp prints usage information for your CLI.
func PrintHelp() {
	fmt.Println("Usage: schema-gen extract [options]")
	fmt.Println()
	pflag.PrintDefaults()
}
