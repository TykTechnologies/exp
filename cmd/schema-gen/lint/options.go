package lint

import (
	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string
	rules     []string
	verbose   bool
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "schema.json",
		rules:     []string{"godoc-fields"},
		verbose:   false,
	}
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	return cfg
}
