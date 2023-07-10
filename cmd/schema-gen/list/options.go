package list

import (
	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "schema.json",
	}
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.Parse()
	return cfg
}
