package stats

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string
	filter    []string
	verbose   bool
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "go-fsck.json",
	}

	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.StringSliceVar(&cfg.filter, "filter", cfg.filter, "filter imports that match (csv)")
	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s stats <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
