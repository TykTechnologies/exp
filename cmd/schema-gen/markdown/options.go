package markdown

import (
	"fmt"
	"os"
	"strings"

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

	full    bool
	keep    []string
	skip    []string
	replace map[string]string

	trim         string
	fieldSpacing bool
}

func flagNameFromEnvironmentName(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, "_", "-", -1)
	return s
}

func containsFlag(haystack []string, needle string) bool {
	for _, v := range haystack {
		if strings.HasPrefix(v, needle) {
			return true
		}
	}
	return false
}

func Parse() {
	for _, v := range os.Environ() {
		vals := strings.SplitN(v, "=", 2)
		flagName := flagNameFromEnvironmentName(vals[0])
		if fn := flag.CommandLine.Lookup(flagName); fn == nil {
			continue
		}
		flagOption := "--" + flagName
		if containsFlag(os.Args, flagOption) {
			continue
		}
		os.Args = append(os.Args, flagOption, vals[1])
	}
	flag.Parse()
}

func NewOptions() *options {
	cfg := &options{
		inputFile:        "schema.json",
		outputFile:       "schema.md",
		headingFormat:    "# %s",
		fieldFormat:      "**Field: `%s` ([%s](#%s))**",
		fieldFormatKnown: "**Field: `%s` (`%s`)**",
		replace:          make(map[string]string),
		rootElement:      "",
	}

	flag.StringVarP(&cfg.outputFile, "output-file", "o", cfg.outputFile, "output file")
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.StringVarP(&cfg.packageName, "package-name", "p", cfg.packageName, "package name")

	flag.StringVar(&cfg.title, "title", cfg.title, "title of markdown doc")
	flag.StringVar(&cfg.headingFormat, "heading-format", cfg.headingFormat, "heading format")
	flag.StringVar(&cfg.fieldFormat, "field-format", cfg.fieldFormat, "field format")
	flag.StringVar(&cfg.fieldFormatKnown, "field-format-known", cfg.fieldFormatKnown, "field format for known types")

	flag.BoolVar(&cfg.full, "full", cfg.full, "print package info with symbols")
	flag.StringSliceVar(&cfg.keep, "keep", cfg.keep, "type definition names to keep (default: all)")
	flag.StringSliceVar(&cfg.skip, "skip", cfg.skip, "type definition names to skip (default: none)")
	flag.StringToStringVar(&cfg.replace, "replace", cfg.replace, "type replacement string to string csv (default: none)")

	flag.StringVar(&cfg.trim, "trim", cfg.trim, "trim lines from docs output")
	flag.StringVar(&cfg.rootElement, "root", cfg.rootElement, "root type to put first in output")

	Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: schema-gen markdown <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
