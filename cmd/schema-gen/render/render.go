package render

import (
	"errors"
	"os"

	flag "github.com/spf13/pflag"
	"golang.org/x/exp/slices"
)

// Run is the entrypoint for `schema-gen extract`.
func Run() (err error) {
	var (
		outputFile = "schema.json"
		sourcePath = "."
	)
	flag.StringVar(&outputFile, "o", outputFile, "output file")
	flag.StringVar(&sourcePath, "i", sourcePath, "source path")
	flag.Parse()

	if slices.Contains(os.Args, "help") {
		flag.PrintDefaults()
		return nil
	}

	return errors.New("unimplemented")
}
