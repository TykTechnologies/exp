package extract

import (
	"encoding/json"
	"os"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

func extract(cfg *options) error {
	definitions, err := model.Load(cfg.sourcePath, cfg.verbose)
	if err != nil {
		return err
	}

	output := os.Stdout
	switch cfg.outputFile {
	case "", "-":
	default:
		var err error
		output, err = os.Create(cfg.outputFile)
		if err != nil {
			return err
		}
	}

	encoder := json.NewEncoder(output)
	if cfg.prettyJSON {
		encoder.SetIndent("", "  ")
	}

	return encoder.Encode(definitions)
}
