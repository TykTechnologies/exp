package lint

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"golang.org/x/exp/slices"
)

// Run is the entrypoint for `schema-gen restore`.
func Run() (err error) {
	cfg := NewOptions()
	flag.Parse()

	if slices.Contains(os.Args, "help") {
		fmt.Println("Usage: schema-gen lint <options>:")
		fmt.Println()
		flag.PrintDefaults()
		return nil
	}

	return lint(cfg)
}
