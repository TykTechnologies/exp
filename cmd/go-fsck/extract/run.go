package extract

import (
	"os"

	"golang.org/x/exp/slices"
)

// Run is the entrypoint for `go-fsck extract`.
func Run() (err error) {
	cfg := NewOptions()

	if slices.Contains(os.Args, "help") {
		PrintHelp()
		return nil
	}

	return extract(cfg)
}
