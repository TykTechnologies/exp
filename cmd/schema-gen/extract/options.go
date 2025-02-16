package extract

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	sourcePath string
	outputFile string

	includeUnexported bool
	includeFunctions  bool
	includeTests      bool
	includeInternal   bool
	ignoreFiles       []string

	prettyJSON bool
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
	flag.BoolVar(&cfg.includeUnexported, "include-unexported", cfg.includeUnexported, "include unexported symbols")
	flag.BoolVar(&cfg.includeTests, "include-tests", cfg.includeTests, "include test files")
	flag.BoolVar(&cfg.includeInternal, "include-internal", cfg.includeTests, "include internal packages")
	flag.StringSliceVarP(&cfg.ignoreFiles, "ignore-files", "", cfg.ignoreFiles, "ignore files (csv)")
	flag.BoolVar(&cfg.prettyJSON, "pretty-json", cfg.prettyJSON, "print pretty json")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: schema-gen extract <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
