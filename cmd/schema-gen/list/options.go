package list

import (
	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string

	includeVariables bool
	includeFunctions bool

	sorted bool

	json       bool
	prettyJSON bool
}

func NewOptions() *options {
	cfg := &options{
		inputFile:        "schema.json",
		includeVariables: true,
	}
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file (glob)")
	flag.BoolVar(&cfg.json, "json", cfg.json, "print json")
	flag.BoolVar(&cfg.includeVariables, "includeVariables", cfg.includeVariables, "print variables")
	flag.BoolVar(&cfg.includeFunctions, "includeFunctions", cfg.includeFunctions, "print functions")
	flag.BoolVar(&cfg.sorted, "sorted", cfg.sorted, "print sorted")
	flag.BoolVar(&cfg.prettyJSON, "pretty-json", cfg.prettyJSON, "pretty print json (json implied)")
	flag.Parse()
	return cfg
}
