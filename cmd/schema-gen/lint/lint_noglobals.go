package lint

import "github.com/TykTechnologies/exp/cmd/schema-gen/model"

func linterNoGlobals(cfg *options, pkgInfo *model.PackageInfo) *LintError {
	return nil
}
