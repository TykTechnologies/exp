package extract

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	sourcePath string
	outputFile string

	includeUnexposed bool
	includeFunctions bool
	includeTests     bool

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
	flag.BoolVar(&cfg.includeFunctions, "include-functions", cfg.includeFunctions, "include functions")
	flag.BoolVar(&cfg.includeUnexposed, "include-unexposed", cfg.includeUnexposed, "include unexposed symbols")
	flag.BoolVar(&cfg.includeTests, "include-tests", cfg.includeTests, "include test files")
	flag.StringSliceVarP(&cfg.ignoreFiles, "ignore-files", "", cfg.ignoreFiles, "ignore files (csv)")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: schema-gen restore <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
