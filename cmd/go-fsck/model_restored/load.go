package model

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"sort"

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
			filename := path.Base(fset.Position(file.Pos()).Filename)

			src, err := os.ReadFile(path.Join(sourcePath, filename))
			if err != nil {
				return nil, fmt.Errorf("Error reading in source file: %s", filename)
			}

			tags := BuildTags(src)
			if len(tags) == 0 {
				files = append(files, file)
				continue
			}

			fmt.Printf("WARN: Skipping file %s with build tags: %v\n", filename, tags)
		}
	}

	collector := NewCollector(fset)

	insp := inspector.New(files)
	insp.WithStack(nil, collector.Visit)

	collector.Clean()

	results := make([]*Definition, 0, len(collector.definition))
	pkgNames := make([]string, 0, len(collector.definition))
	for _, pkg := range collector.definition {
		pkgNames = append(pkgNames, pkg.Package)
	}
	sort.Strings(pkgNames)

	for _, pkg := range collector.definition {
		for _, name := range pkgNames {
			if pkg.Package == name {
				results = append(results, pkg)
			}
		}
	}

	return results, nil
}
