package restore

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/stoewer/go-strcase"
	"golang.org/x/exp/maps"

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
		name := t.Name

		// Group receivers next to type declaration:
		//
		// Receiver can be *T or T or unset; If it's set, that function belongs
		// into $T.go; the function behaviour is explicitly bound to T.
		if t.Receiver != "" {
			filename := strcase.SnakeCase(strings.TrimLeft(t.Receiver, "*")) + ".go"
			return filename
		}

		// Constructor naming conventions:
		//
		// Match New$T functions into $T scope.

		if filename, ok := strings.CutPrefix(name, "New"); ok {
			filename = strcase.SnakeCase(filename) + ".go"
			if _, exists := files[filename]; exists {
				return filename
			}
		}

		// Test naming conventions:
		//
		// - name test `TestStructName` if it tests StructName{},
		// - `TestStructName_subTest` for individual scoped tests,
		//
		// If a test depends on multiple objects and cannot be scoped
		// into a $T_test.go file, then it will live in funcs_test.go.
		//
		// This is expected on some level. When coupling types together,
		// a func($T, $V) is expected to be global. Structs may provide
		// conversions, but code that needs multiple types doesn't have
		// a definition of local behaviour as it needs both $T and $V
		// to work. This means that the code will work in the package,
		// but won't work with a reduced scope, until both $T and $V are
		// moved into importable packages.

		if filename, ok := strings.CutPrefix(name, "Test"); ok {
			cleanName := strings.SplitN(filename, "_", 2)
			filename = strcase.SnakeCase(cleanName[0]) + ".go"
			if _, exists := files[filename]; exists {
				return filename
			}
			return "funcs_test.go"
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

	filenames := maps.Keys(files)
	sort.Slice(filenames, func(i int, j int) bool {
		clean := func(s string) string {
			cut := ".go"
			s = s[0 : len(s)-len(cut)]
			cut = "_test"
			if strings.HasSuffix(s, cut) {
				return s[:len(s)-len(cut)]
			}
			return s
		}
		p1, p2 := filenames[i], filenames[j]
		c1, c2 := clean(p2), clean(p2)
		if c1 == c2 {
			return strings.Compare(p1, p2) < 0
		}
		return strings.Compare(c1, c2) < 0
	})

	for _, filename := range filenames {
		decls := files[filename]

		fmt.Println("#", filename)
		fmt.Println()
		for _, decl := range decls {
			if decl.Kind == model.FuncKind {
				if decl.Receiver != "" {
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
