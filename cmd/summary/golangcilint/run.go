package golangcilint

import (
	"encoding/json"
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

	return run(cfg)
}

func run(cfg *options) error {
	decoder := json.NewDecoder(os.Stdin)

	input := &Root{}
	err := decoder.Decode(input)
	if err != nil {
		return err
	}

	output := Convert(input)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(output)
}
