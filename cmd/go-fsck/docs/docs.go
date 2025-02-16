package docs

import (
	"fmt"
	"strings"

	"github.com/TykTechnologies/exp/cmd/go-fsck/internal"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
	"github.com/TykTechnologies/exp/cmd/go-fsck/model/loader"
)

func getDefinitions(cfg *options) ([]*model.Definition, error) {
	// Read the exported go-fsck.json data.
	defs, err := loader.ReadFile(cfg.inputFile)
	if err == nil {
		return defs, nil
	}

	// list current local packages
	packages, err := internal.ListPackages(".", "./...")
	if err != nil {
		return nil, err
	}

	defs = []*model.Definition{}

	for _, pkg := range packages {
		d, err := loader.Load(pkg, cfg.verbose)
		if err != nil {
			return nil, err
		}

		for _, v := range d {
			v.Package.ID = pkg.ID
			v.Package.ImportPath = pkg.ImportPath
			v.Package.Path = pkg.Path
			v.Package.Package = pkg.Package
			v.Package.TestPackage = pkg.TestPackage
		}

		defs = append(defs, d...)
	}

	return defs, nil
}

func docs(cfg *options) error {
	defs, err := getDefinitions(cfg)
	if err != nil {
		return err
	}

	// Loop through function definitions and collect referenced
	// symbols from imported packages. Globals may also reference
	// imported packages so this is incomplete at the moment.
	for _, def := range defs {
		if def.Package.TestPackage || strings.HasSuffix(def.Package.Package, "_test") {
			continue
		}

		var (
			types  = def.Types.Exported()
			consts = def.Consts.Exported()
			vars   = def.Vars.Exported()
			funcs  = def.Funcs.Exported()
		)

		var packageName = def.Package.ImportPath

		fmt.Println("# Package", packageName)
		fmt.Println()
		if def.Doc != "" {
			fmt.Println(strings.TrimSpace(def.Doc))
			fmt.Println()
		}

		if len(types) > 0 {
			fmt.Println("## Types\n")
			for _, v := range types {
				src := strings.TrimSpace(v.Source)
				fmt.Printf("```go\n%s\n```\n\n", src)
			}
		}

		if len(consts) > 0 {
			fmt.Println("## Consts\n")
			for _, v := range consts {
				src := strings.TrimSpace(v.Source)
				fmt.Printf("```go\n%s\n```\n\n", src)
			}
		}
		if len(consts) > 0 {
			fmt.Println("## Vars\n")
			for _, v := range vars {
				src := strings.TrimSpace(v.Source)
				fmt.Printf("```go\n%s\n```\n\n", src)
			}
		}

		symbol := func(fn *model.Declaration) string {
			if fn.Receiver != "" {
				return "func (" + fn.Receiver + ") " + fn.Signature
			}
			return "func " + fn.Signature
		}

		if len(funcs) > 0 {
			fmt.Println("## Function symbols\n")

			for _, fn := range funcs {
				fmt.Printf("- `%s`\n", symbol(fn))
			}
			fmt.Println()

			// Documented functions first.
			for _, fn := range funcs {
				if fn.Doc == "" {
					continue
				}

				fmt.Printf("### %s\n\n", fn.Name)
				fmt.Println(strings.TrimSpace(fn.Doc))
				fmt.Printf("```go\n%s\n```\n\n", symbol(fn))
				fmt.Println()
			}

			// List undocumented ones.
			for _, fn := range funcs {
				if fn.Doc != "" {
					continue
				}

				fmt.Printf("### %s\n\n", fn.Name)
				fmt.Printf("```go\n%s\n```\n\n", symbol(fn))
			}
			fmt.Println()
		}
	}

	return nil
}
