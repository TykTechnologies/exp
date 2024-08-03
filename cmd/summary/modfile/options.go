package modfile

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	Contains string
}

func NewOptions() *options {
	cfg := &options{}

	flag.StringVar(&cfg.Contains, "contains", "", "Filter imports containing pattern")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s lsof <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
