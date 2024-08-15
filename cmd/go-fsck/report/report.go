package report

import (
	"fmt"

	"github.com/TykTechnologies/exp/cmd/go-fsck/internal"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model/loader"
	"github.com/TykTechnologies/exp/cmd/go-fsck/report/testusage"
)

func getDefinitions(cfg *options) ([]*model.Definition, error) {
	// Read the exported go-fsck.json data.
	defs, err := loader.ReadFile(cfg.inputFile)
	if err == nil {
		return defs, nil
	}

	// list current local package
	packages, err := internal.ListPackages(".", ".")
	if err != nil {
		return nil, err
	}

	defs = []*model.Definition{}

	for _, pkg := range packages {
		d, err := loader.Load(pkg.Path, cfg.verbose)
		if err != nil {
			return nil, err
		}

		defs = append(defs, d...)
	}

	return defs, nil
}

func report(cfg *options) error {
	defs, err := getDefinitions(cfg)
	if len(defs) == 0 || err != nil {
		return fmt.Errorf("error getting definitions: %w, len %d", err, len(defs))
	}

	report, err := testusage.NewReport(defs)
	if err != nil {
		return fmt.Errorf("error generating report: %w", err)
	}

	fmt.Println(report)

	return nil
}