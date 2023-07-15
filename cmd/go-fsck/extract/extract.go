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

	body, err := json.Marshal(definitions)
	if err != nil {
		return err
	}

	return os.WriteFile(cfg.outputFile, body, 0644)
}
