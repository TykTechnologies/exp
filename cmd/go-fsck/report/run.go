package report

import (
	"os"

	"golang.org/x/exp/slices"
)

// Run is the entrypoint for `go-fsck report`.
func Run() (err error) {
	cfg := NewOptions()

	if slices.Contains(os.Args, "help") {
		PrintHelp()
		return nil
	}

	return report(cfg)
}
