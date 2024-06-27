package lint

import (
	"fmt"

	"golang.org/x/exp/slices"

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
		errs.Combine(runLinter(cfg, NewLinter("require-comment", linterStructs), pkgInfo))
		errs.Combine(runLinter(cfg, NewLinter("require-field-comment", linterFields), pkgInfo))
		errs.Combine(runLinter(cfg, NewLinter("require-no-globals", linterNoGlobals), pkgInfo))
		errs.Combine(runLinter(cfg, NewLinter("require-dot-or-backtick", linterFields), pkgInfo))
		errs.Combine(runLinter(cfg, NewLinter("require-prefix", linterFields), pkgInfo))

		if errs.Empty() {
			return nil
		}
		return errs
	}

	return nil
}

func runLinter(cfg *options, linter Linter, pkgInfo *PackageInfo) *LintError {
	if !slices.Contains(cfg.GetRules(), linter.GetName()) {
		return nil
	}

	return linter.Do(cfg, pkgInfo)
}
