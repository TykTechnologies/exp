package coverfunc

import (
	"fmt"
	"os"
	"path"

	flag "github.com/spf13/pflag"
)

type options struct {
	GroupByFiles bool

	RenderJSON bool
}

func NewOptions() *options {
	cfg := &options{}

	flag.BoolVarP(&cfg.GroupByFiles, "files", "f", false, "Group coverage by file")
	flag.BoolVar(&cfg.RenderJSON, "json", false, "Render output as json")
	flag.Parse()

	return cfg
}

func PrintHelp() {
	fmt.Printf("Usage: %s coverfunc <options>:\n\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}
