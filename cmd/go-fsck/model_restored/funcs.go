package model

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
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
			files = append(files, file)
		}
	}

	collector := NewCollector(fset)

	insp := inspector.New(files)
	insp.WithStack(nil, collector.Visit)

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

func NewCollector(fset *token.FileSet) *collector {
	return &collector{
		fset:       fset,
		definition: make(map[string]*Definition),
		seen:       make(map[string]*Declaration),
	}
}

func getBuildTags(file *ast.File) []string {

	re := regexp.MustCompile(`^\s*//\s*\+build\s+(.*)$`)

	var buildTags []string

	if file.Doc != nil {
		for _, comment := range file.Doc.List {
			match := re.FindStringSubmatch(comment.Text)
			if len(match) > 1 {
				buildTags = append(buildTags, match[1])
			}
		}
	}

	return buildTags
}
