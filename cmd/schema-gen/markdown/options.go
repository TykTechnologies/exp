package markdown

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile   string
	outputFile  string
	rootElement string
}

func NewOptions() *options {
	cfg := &options{
		inputFile:   "schema.json",
		outputFile:  "schema.md",
		rootElement: "",
	}
	flag.StringVarP(&cfg.outputFile, "output-file", "o", cfg.outputFile, "output file")
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.StringVar(&cfg.rootElement, "root", cfg.rootElement, "root type to put first in output")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: schema-gen markdown <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
