package restore

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	fix bool

	inputFile   string
	outputFile  string
	packageName string
	rootElement string

	keep             []string
	includeFunctions []string
}

func NewOptions() *options {
	cfg := &options{
		inputFile:   "schema.json",
		outputFile:  "schema.go.txt",
		packageName: "schema",
	}
	flag.StringVarP(&cfg.outputFile, "output-file", "o", cfg.outputFile, "output file")
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.StringVarP(&cfg.packageName, "package-name", "p", cfg.packageName, "package name")

	flag.StringSliceVar(&cfg.keep, "keep", cfg.keep, "type definition names to keep (default: all)")
	flag.StringSliceVar(&cfg.includeFunctions, "include-functions", cfg.includeFunctions, "include functions (default: none)")
	flag.BoolVar(&cfg.fix, "fix", cfg.fix, "generate output files according to struct names")
	flag.StringVar(&cfg.rootElement, "root", cfg.rootElement, "root type to put first in output")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: schema-gen restore <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
