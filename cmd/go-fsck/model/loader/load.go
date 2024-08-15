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
		/*
			if verbose {
				spew.Dump(pkg)
			}
			for _, file := range pkg.Syntax {
				filename := path.Base(fset.Position(file.Pos()).Filename)
				if verbose {
					log.Printf("  Filename:   %s", filename)
				}

				if !strings.HasSuffix(filename, ".go") {
					log.Printf("skipping %s.%s", pkg.Name, filename)
					continue
				}

				src, err := os.ReadFile(path.Join(sourcePath, filename))
				if err != nil {
					return nil, fmt.Errorf("Error reading in source file: %s", filename)
				}

				tags := BuildTags(src)
				if len(tags) == 0 {
					files = append(files, file)
					continue
				}

				if verbose {
					log.Printf("WARN: Skipping file %s with build tags: %v", filename, tags)
				}
			} */
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
