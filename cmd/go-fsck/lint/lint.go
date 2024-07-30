package lint

import (
	"errors"
	"fmt"

	"github.com/TykTechnologies/exp/cmd/go-fsck/internal"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model/loader"
)

func getDefinitions(cfg *options) ([]*model.Definition, error) {
	// list current local packages
	packages, err := internal.ListCurrent()
	if err != nil {
		return nil, err
	}

	defs := []*model.Definition{}

	for _, pkgPath := range packages {
		d, err := loader.Load(pkgPath, cfg.verbose)
		if err != nil {
			return nil, err
		}

		defs = append(defs, d...)
	}

	return defs, nil
}

func lint(cfg *options) error {
	var lintErrors []error

	defs, err := getDefinitions(cfg)
	if err != nil {
		return err
	}

	for _, def := range defs {
		_, importCollisions := def.Imports.Map()
		for _, err := range importCollisions {
			lintErrors = append(lintErrors, err)
		}
	}

	if len(lintErrors) == 0 {
		return nil
	}

	for _, err := range lintErrors {
		fmt.Println(err)
	}

	return errors.New("Linter not passing")
}
