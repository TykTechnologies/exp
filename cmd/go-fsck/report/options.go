package report

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile string

	json    bool
	verbose bool
	args    []string
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "go-fsck.json",
	}

	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.BoolVar(&cfg.json, "json", cfg.json, "print results as json")
	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	flag.Parse()

	cfg.args = flag.Args()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s report <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
