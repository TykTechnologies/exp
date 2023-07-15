package extract

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	sourcePath string
	outputFile string

	prettyJSON bool
}

func NewOptions() *options {
	cfg := &options{
		sourcePath: ".",
		outputFile: "go-fsck.json",
	}
	flag.StringVarP(&cfg.outputFile, "output-file", "o", cfg.outputFile, "output file")
	flag.StringVarP(&cfg.sourcePath, "source-path", "i", cfg.sourcePath, "source path")
	flag.BoolVar(&cfg.prettyJSON, "pretty-json", cfg.prettyJSON, "print pretty json")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s extract <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
