package model

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
