package modfile

import (
	"encoding/json"
	"io"
	"os"

	"golang.org/x/exp/slices"
	"golang.org/x/mod/module"
)

// Run is the entrypoint for the plugin.
func Run() (err error) {
	cfg := NewOptions()

	if slices.Contains(os.Args, "help") {
		PrintHelp()
		return nil
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	listOfModules, err := Parse(input, cfg.Contains)
	if err != nil {
		return err
	}

	output := []module.Version{} //modfile.Mod{}
	for _, v := range listOfModules {
		output = append(output, v.Mod)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(output)
	return nil
}
