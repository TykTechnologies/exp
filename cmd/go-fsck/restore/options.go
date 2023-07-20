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

	removeUnexported bool
	removeTests      bool

	save        bool
	packageName string
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

	flag.BoolVar(&cfg.removeUnexported, "remove-unexported", cfg.removeUnexported, "remove unexported symbols")
	flag.BoolVar(&cfg.removeTests, "remove-tests", cfg.removeTests, "remove tests")

	flag.BoolVar(&cfg.save, "save", cfg.save, "write out to files")
	flag.StringVarP(&cfg.packageName, "package-name", "p", cfg.packageName, "package name for --save")

	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s restore <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
