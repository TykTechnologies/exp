package restore

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/stoewer/go-strcase"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

func restore(cfg *options) error {
	defs, err := model.ReadFile(cfg.inputFile)
	if err != nil {
		return err
	}

	files := make(map[string][]*model.Declaration, 0)
	add := func(filename string, decl ...*model.Declaration) {
		s, ok := files[filename]
		if !ok {
			files[filename] = decl
			return
		}

		s = append(s, decl...)
		files[filename] = s
	}
	classFunc := func(t *model.Declaration) string {
		// Receiver can be *T or T or unset;
		if t.Receiver != "" {
			filename := strcase.SnakeCase(strings.TrimLeft(t.Receiver, "*")) + ".go"
			return filename
		}

		// Match New$T functions into $T scope.
		name := t.Name
		if strings.HasPrefix(name, "New") {
			filename := strcase.SnakeCase(name[3:]) + ".go"
			if _, exists := files[filename]; exists {
				return filename
			}
		}

		return "funcs.go"
	}

	for _, def := range defs {
		for _, t := range def.Types {
			name := findShortest(t.Names, t.Name)
			filename := strcase.SnakeCase(name) + ".go"

			add(filename, t)
		}

		for _, t := range def.Funcs {
			filename := classFunc(t)
			add(filename, t)
		}

		for _, t := range def.Vars {
			add("vars.go", t)
		}

		for _, t := range def.Consts {
			add("const.go", t)
		}
	}

	b, _ := json.MarshalIndent(files, "", "  ")
	fmt.Println(string(b))

	for filename, decls := range files {
		fmt.Println("#", filename)
		fmt.Println()
		for _, decl := range decls {
			if decl.Kind == model.FuncKind {
				if decl.Signature != "" {
					fmt.Printf("- func (%s) %s\n", decl.Receiver, decl.Signature)
					continue
				}
				fmt.Printf("- func %s\n", decl.Signature)
				continue
			}
			if len(decl.Names) > 0 {
				for _, name := range decl.Names {
					fmt.Println("-", decl.Kind, name)
				}
				continue
			}
			fmt.Println("-", decl.Kind, decl.Name)
		}
		fmt.Println()
	}

	return nil
}

func findShortest(s []string, def string) string {
	if len(s) > 0 {
		r := s[0]
		for _, v := range s {
			if len(v) < len(r) {
				r = v
			}
		}
		return r
	}
	return def
}
