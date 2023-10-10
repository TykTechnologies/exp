package internal

import (
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

// List returns a slice of local package paths in the specified root directory.
func List(rootPath string) ([]string, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, err
	}

	first := pkgs[0].PkgPath

	var localPackages []string

	for _, pkg := range pkgs {
		localPackages = append(localPackages, strings.ReplaceAll(pkg.PkgPath, first, "."))
	}

	return localPackages, nil
}

// ListCurrent uses the current path as the root directory for List().
func ListCurrent() ([]string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return List(currentPath)
}
