package model

import (
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/inspector"
)

// Load definitions from package located in sourcePath.
func Load(sourcePath string) ([]*Definition, error) {
	fset := token.NewFileSet()

	packages, err := parser.ParseDir(fset, sourcePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	files := []*ast.File{}
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			files = append(files, file)
		}
	}

	collector := NewCollector(fset)

	insp := inspector.New(files)
	insp.WithStack(nil, collector.Visit)

	results := make([]*Definition, 0, len(collector.definition))
	for _, pkg := range collector.definition {
		results = append(results, pkg)
	}
	return results, nil
}
