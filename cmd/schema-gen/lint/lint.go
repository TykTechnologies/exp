package lint

import (
	"fmt"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
	"golang.org/x/exp/slices"
)

func lint(cfg *options) error {
	fmt.Println(cfg.inputFile)

	pkgInfo, err := model.Load(cfg.inputFile)
	if err != nil {
		return fmt.Errorf("Error loading package info: %w", err)
	}

	// require-no-globals
	//
	// When doing model/repository work, the package `model` provides
	// the type declarations for `structs`, `var` and `const`'s.
	//
	// The repository package should provide interfaces, possibly
	// constructor functions for the repositories, but should not
	// expose symbols outside of that. Those belong in model.

	errs := NewLintError()
	errs.Combine(runLinter(cfg, NewLinter("require-no-globals", linterNoGlobals), pkgInfo))

	errs.Combine(runLinter(cfg, NewLinter("require-field-comment", linterFields), pkgInfo))
	errs.Combine(runLinter(cfg, NewLinter("require-dot-or-backtick", linterFields), pkgInfo))
	errs.Combine(runLinter(cfg, NewLinter("require-field-prefix", linterFields), pkgInfo))

	if errs.Empty() {
		return nil
	}
	return errs
}

func runLinter(cfg *options, linter Linter, pkgInfo *PackageInfo) *LintError {
	if !slices.Contains(cfg.rules, linter.GetName()) {
		return nil
	}

	return linter.Do(cfg, pkgInfo)
}
