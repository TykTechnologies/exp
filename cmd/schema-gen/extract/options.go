package extract

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	sourcePath string
	outputFile string

	includeFunctions     bool
	includeTestFunctions bool

	ignoreFiles []string
}

func NewOptions() *options {
	cfg := &options{
		sourcePath:  ".",
		outputFile:  "schema.json",
		ignoreFiles: []string{},
	}
	flag.StringVarP(&cfg.outputFile, "output-file", "o", cfg.outputFile, "output file")
	flag.StringVarP(&cfg.sourcePath, "source-path", "i", cfg.sourcePath, "source path")
	flag.StringSliceVarP(&cfg.ignoreFiles, "ignore-files", "", cfg.ignoreFiles, "ignore files (csv)")
	flag.BoolVar(&cfg.includeFunctions, "include-functions", cfg.includeFunctions, "include functions")
	flag.BoolVar(&cfg.includeTestFunctions, "include-test-functions", cfg.includeTestFunctions, "include test functions")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: schema-gen restore <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
