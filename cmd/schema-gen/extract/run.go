package extract

import (
	"fmt"
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
	flag.StringVarP(&outputFile, "output-file", "o", outputFile, "output file")
	flag.StringVarP(&sourcePath, "source-path", "i", sourcePath, "source path")
	flag.Parse()

	if slices.Contains(os.Args, "help") {
		fmt.Println("Usage: schema-gen extract <options>:")
		fmt.Println()
		flag.PrintDefaults()
		return nil
	}

	return write(outputFile, sourcePath)
}
