package lint

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	verbose bool
	args    []string
}

func NewOptions() *options {
	cfg := &options{}
	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	flag.Parse()

	cfg.args = flag.Args()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s lint <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
