package model

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"sort"

	"golang.org/x/tools/go/ast/inspector"
)

type (
	Definition struct {
		Package string
		Doc     StringSet
		Imports StringSet
		Types   DeclarationList
		Consts  DeclarationList
		Vars    DeclarationList
		Funcs   DeclarationList
	}
)

// Load definitions from package located in sourcePath.
func Load(sourcePath string, verbose bool) ([]*Definition, error) {
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

			fmt.Fprintf(os.Stderr, "WARN: Skipping file %s with build tags: %v\n", filename, tags)
		}
	}

	collector := NewCollector(fset)

	insp := inspector.New(files)
	insp.WithStack(nil, collector.Visit)

	collector.Clean(verbose)

	results := make([]*Definition, 0, len(collector.definition))
	pkgNames := make([]string, 0, len(collector.definition))
	for _, pkg := range collector.definition {
		pkg.Sort()
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

	sort.Slice(results, func(i, j int) bool {
		return results[i].Package < results[j].Package
	})

	return results, nil
}

// ReadFile loads the definitions from a json file
func ReadFile(inputPath string) ([]*Definition, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	var result []*Definition
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	for _, decl := range result {
		decl.Fill()
	}

	return result, nil
}

func (d *Definition) Fill() {
	for _, decl := range d.Order() {
		decl.Imports = d.getImports(decl)
	}
}

func (d *Definition) Order() []*Declaration {
	count := len(d.Types) + len(d.Funcs) + len(d.Vars) + len(d.Consts)
	result := make([]*Declaration, 0, count)

	result = append(result, d.Types...)
	result = append(result, d.Funcs...)
	result = append(result, d.Vars...)
	result = append(result, d.Consts...)
	return result
}

func (d *Definition) Sort() {
	d.Types.Sort()
	d.Vars.Sort()
	d.Consts.Sort()
	d.Funcs.Sort()
}

func (d *Definition) getImports(decl *Declaration) []string {
	return d.Imports.Get(decl.File)
}
