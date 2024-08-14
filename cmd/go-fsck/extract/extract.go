package extract

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TykTechnologies/exp/cmd/go-fsck/internal"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model/loader"
)

func getDefinitions(cfg *options) ([]*model.Definition, error) {
	// list current local packages
	var pattern string
	if pattern = "."; cfg.recursive {
		pattern = "./..."
	}

	packages, err := internal.ListPackages(cfg.sourcePath, pattern)
	if err != nil {
		return nil, err
	}

	defs := []*model.Definition{}

	for _, pkg := range packages {
		d, err := loader.Load(pkg.Path, cfg.verbose)
		if err != nil {
			return nil, err
		}
		for _, v := range d {
			v.Package.ImportPath = pkg.ImportPath
			v.Package.Path = pkg.Path
			v.Package.Package = pkg.Package

			if !cfg.includeTests && pkg.TestPackage {
				continue
			}

			defs = append(defs, d...)
		}
	}

	if !cfg.includeSources {
		for _, def := range defs {
			def.ClearSource()
		}
	}

	return defs, nil
}

func extract(cfg *options) error {
	definitions, err := getDefinitions(cfg)
	if err != nil {
		return err
	}

	output := os.Stdout
	switch cfg.outputFile {
	case "", "-":
	default:
		fmt.Println(cfg.outputFile)
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
