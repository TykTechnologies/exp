package jsonschema

import (
	"os"
	"slices"
)

// Run is the entrypoint for `schema-gen jsonschema`.
func Run() (err error) {
	cfg := NewOptions()

	if slices.Contains(os.Args, "help") {
		PrintHelp()
		return nil
	}

	return parseAndConvertStruct(cfg.sourcePath, cfg.rootType, cfg.outputFile)
}
