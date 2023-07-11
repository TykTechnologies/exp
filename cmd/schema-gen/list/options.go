package list

import (
	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string
	summary   bool
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "schema.json",
	}
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.BoolVar(&cfg.summary, "summary", cfg.summary, "print summary")
	flag.Parse()
	return cfg
}
