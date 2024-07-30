package loader

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"

	"golang.org/x/tools/go/ast/inspector"

	"github.com/TykTechnologies/exp/cmd/go-fsck/internal/collector"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

// Load definitions from package located in sourcePath.
func Load(sourcePath string, verbose bool) ([]*model.Definition, error) {
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

	sink := collector.NewCollector(fset)

	insp := inspector.New(files)
	insp.WithStack(nil, sink.Visit)

	results := sink.Clean(verbose)

	return results, nil
}

// ReadFile loads the definitions from a json file
func ReadFile(inputPath string) ([]*model.Definition, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	var result []*model.Definition
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	for _, decl := range result {
		decl.Fill()
	}

	return result, nil
}
