package restore

import (
	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile   string
	outputFile  string
	packageName string
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
	return cfg
}
