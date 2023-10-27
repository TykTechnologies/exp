package lint

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	lintImports bool

	verbose bool
}

func NewOptions() *options {
	cfg := &options{
		lintImports: true,
	}

	flag.BoolVar(&cfg.lintImports, "lint-all", cfg.lintImports, "lint imports")

	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s lint <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
