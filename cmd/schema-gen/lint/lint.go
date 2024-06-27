package lint

import (
	"fmt"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

func lint(cfg *options) error {
	pkgInfos, err := model.Load(cfg.inputFile)
	if err != nil {
		return fmt.Errorf("Error loading package info: %w", err)
	}

	for _, pkgInfo := range pkgInfos {

		errs := NewLintError()
		errs.Combine(runLinter(cfg, NewLinter("lint structs", linterStructs), pkgInfo))
		errs.Combine(runLinter(cfg, NewLinter("lint fields", linterFields), pkgInfo))
		errs.Combine(runLinter(cfg, NewLinter("lint globals", linterNoGlobals), pkgInfo))

		if errs.Empty() {
			return nil
		}
		return errs
	}

	return nil
}

func runLinter(cfg *options, linter Linter, pkgInfo *PackageInfo) *LintError {
	return linter.Do(cfg, pkgInfo)
}
