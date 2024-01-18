package markdown

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile   string
	outputFile  string
	packageName string
	rootElement string

	title         string
	headingFormat string

	fieldFormat      string
	fieldFormatKnown string

	keep []string
	skip []string

	trim         string
	fieldSpacing bool
}

func NewOptions() *options {
	cfg := &options{
		inputFile:        "schema.json",
		outputFile:       "schema.md",
		headingFormat:    "# %s",
		fieldFormat:      "**Field: `%s` ([%s](#%s))**",
		fieldFormatKnown: "**Field: `%s` (`%s`)**",
		rootElement:      "",
	}
	flag.StringVarP(&cfg.outputFile, "output-file", "o", cfg.outputFile, "output file")
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.StringVarP(&cfg.packageName, "package-name", "p", cfg.packageName, "package name")

	flag.StringVar(&cfg.title, "title", cfg.title, "title of markdown doc")
	flag.StringVar(&cfg.headingFormat, "heading-format", cfg.headingFormat, "heading format")
	flag.StringVar(&cfg.fieldFormat, "field-format", cfg.fieldFormat, "field format")
	flag.StringVar(&cfg.fieldFormatKnown, "field-format-known", cfg.fieldFormatKnown, "field format for known types")

	flag.StringSliceVar(&cfg.keep, "keep", cfg.keep, "type definition names to keep (default: all)")
	flag.StringSliceVar(&cfg.skip, "skip", cfg.skip, "type definition names to skip (default: none)")
	flag.StringVar(&cfg.trim, "trim", cfg.trim, "trim lines from docs output")
	flag.StringVar(&cfg.rootElement, "root", cfg.rootElement, "root type to put first in output")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: schema-gen markdown <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
