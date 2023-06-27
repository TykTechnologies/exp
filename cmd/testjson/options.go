package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "test-log.json",
	}
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Println("Usage: testjson <options>:")
	fmt.Println()
	flag.PrintDefaults()
}
