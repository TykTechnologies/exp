package model

import (
	"golang.org/x/tools/go/packages"
)

type Package struct {
	// ID is the ID of the package as x/tools packages loads it.
	ID string
	// Package is the name of the package.
	Package string
	// ImportPath contains the import path (github...).
	ImportPath string
	// Path is sanitized to contain the relative location (folder).
	Path string
	// TestPackage is true if this is a test package.
	TestPackage bool

	// Pkg serves to carry ast package information, preventing a double Load().
	// It's used during analysis and merging and discarded for the result.
	Pkg *packages.Package `json:"-"`
}

func (p Package) Name() string {
	return p.Package
}
