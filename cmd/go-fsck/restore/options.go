package restore

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFile  string
	outputPath string

	ignoreFiles []string
}

func NewOptions() *options {
	cfg := &options{
		inputFile:   "go-fsck.json",
		outputPath:  ".",
		ignoreFiles: []string{"pkg.go"},
	}
	flag.StringVarP(&cfg.outputPath, "output-path", "o", cfg.outputPath, "output path")
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.StringSliceVar(&cfg.ignoreFiles, "ignore-files", cfg.ignoreFiles, "ignore files when writing output")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s restore <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
