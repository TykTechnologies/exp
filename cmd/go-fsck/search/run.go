package search

import (
	"os"

	"golang.org/x/exp/slices"
)

// Run is the entrypoint for `schema-gen search`.
func Run() (err error) {
	cfg := NewOptions()

	if slices.Contains(os.Args, "help") {
		PrintHelp()
		return nil
	}

	return search(cfg)
}
