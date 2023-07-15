package extract

import (
	"encoding/json"
	"os"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

func extract(cfg *options) error {
	definitions, err := model.Load(cfg.sourcePath)
	if err != nil {
		return err
	}

	encode := func(in interface{}) ([]byte, error) {
		if cfg.prettyJSON {
			return json.MarshalIndent(in, "", "  ")
		}
		return json.Marshal(in)
	}

	body, err := encode(definitions)
	if err != nil {
		return err
	}

	return os.WriteFile(cfg.outputFile, body, 0644)
}
