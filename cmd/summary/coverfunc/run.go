package coverfunc

import (
	"os"

	"golang.org/x/exp/slices"
)

// Run is the entrypoint for the plugin.
func Run() (err error) {
	cfg := NewOptions()

	if slices.Contains(os.Args, "help") {
		PrintHelp()
		return nil
	}

	return coverfunc(cfg)
}
