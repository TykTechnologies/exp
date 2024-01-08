package search

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string

	name      string
	reference string

	all     bool
	json    bool
	verbose bool
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "go-fsck.json",
	}

	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")

	flag.StringVar(&cfg.name, "name", cfg.name, "function name match (case sensitive)")
	flag.StringVar(&cfg.reference, "reference", cfg.reference, "reference symbol (e.g. 'oas', or 'oas.OAS')")

	flag.BoolVar(&cfg.all, "all", cfg.all, "traverse all packages (./...)")
	flag.BoolVar(&cfg.json, "json", cfg.json, "print results as json")
	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s search <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
