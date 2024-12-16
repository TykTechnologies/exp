package model

import (
	"fmt"
)

type Package struct {
	// Package is the name of the package.
	Package string
	// ImportPath contains the import path (github...).
	ImportPath string
	// Path is sanitized to contain the relative location (folder).
	Path string
	// TestPackage is true if this is a test package.
	TestPackage bool
}

func (p Package) Name() string {
	return p.Package
}

func (p Package) String() string {
	return fmt.Sprintf("package=%s import_path=%s path=%s test_package=%v", p.Package, p.ImportPath, p.Path, p.TestPackage)
}

func (p Package) Equal(in Package) bool {
	return p.ImportPath == in.ImportPath
}
