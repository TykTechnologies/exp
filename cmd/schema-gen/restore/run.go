package restore

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"golang.org/x/exp/slices"
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

// Run is the entrypoint for `schema-gen extract`.
func Run() (err error) {
	cfg := NewOptions()
	flag.Parse()

	if slices.Contains(os.Args, "help") {
		fmt.Println("Usage: schema-gen restore <options>:")
		fmt.Println()
		flag.PrintDefaults()
		return nil
	}

	return restore(cfg)
}
