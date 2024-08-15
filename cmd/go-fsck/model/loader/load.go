package loader

import (
	"encoding/json"
	"go/ast"
	"go/token"
	"os"

	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"

	"github.com/TykTechnologies/exp/cmd/go-fsck/internal/collector"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

// Load definitions from package located in sourcePath.
func Load(sourcePath string, verbose bool) ([]*model.Definition, error) {
	fset := token.NewFileSet()

	cfg := &packages.Config{
		Mode:  packages.LoadSyntax,
		Tests: true,
		Fset:  fset,
	}

	pkgs, err := packages.Load(cfg, sourcePath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) > 2 {
		pkgs = pkgs[:2]
	}

	files := []*ast.File{}
	for _, pkg := range pkgs {
		files = append(files, pkg.Syntax...)
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
