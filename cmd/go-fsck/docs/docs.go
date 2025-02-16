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
		if def.Package.TestPackage {
			continue
		}

		var (
			types  = def.Types.Exported()
			consts = def.Consts.Exported()
			vars   = def.Vars.Exported()
			funcs  = def.Funcs.Exported()
		)

		fmt.Println("# Package", def.Package.Package)
		fmt.Println()

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

		if len(funcs) > 0 {
			fmt.Println("## Function symbols\n")
			for _, fn := range funcs {
				fmt.Println("-", "`"+fn.Signature+"`")
			}
			fmt.Println()
		}
	}

	return nil
}
