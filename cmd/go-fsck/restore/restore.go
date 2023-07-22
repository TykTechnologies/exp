package restore

import (
	"fmt"
	"go/ast"
	"os"
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

	files := make(map[string]model.DeclarationList, 0)
	add := func(filename string, decls ...*model.Declaration) {
		fileTest := filename[:len(filename)-3] + "_test.go"

		for _, t := range decls {
			dest := filename
			if isTest := strings.HasSuffix(t.File, "_test.go"); isTest {
				dest = fileTest
			}

			s, ok := files[dest]
			if !ok {
				files[dest] = decls
				return
			}
			files[dest] = append(s, t)
		}
	}
	classifyFunc := func(t *model.Declaration) string {
		name := t.Name
		isTest := strings.HasSuffix(t.File, "_test.go")

		findFile := func(find string) (string, bool) {
			for filename, f := range files {
				for _, v := range f {
					if v.Name == find {
						return filename, true
					}
					for _, name := range v.Names {
						if name == find {
							return filename, true
						}
					}
				}
			}
			return "", false
		}

		// Group receivers next to type declaration:
		//
		// Receiver can be *T or T or unset; If it's set, that function belongs
		// into $T.go; the function behaviour is explicitly bound to T.

		receiver := strings.TrimLeft(t.Receiver, "*")
		if receiver != "" {
			filename, ok := findFile(receiver)
			if !ok {
				fmt.Println("Couldn't find receiver for %q", receiver)
				os.Exit(1)
			}

			if isTest {
				if strings.HasSuffix(filename, "_test.go") {
					return filename
				}
				return filename[:len(filename)-3] + "_test.go"
			}

			if filename, changed := cutSuffix(filename, "_test.go"); changed {
				return filename + ".go"
			}
			return filename
		}

		// Constructor naming conventions:
		//
		// Match New$T functions into $T scope.

		if filename, ok := strings.CutPrefix(name, "New"); ok {
			filename = strcase.SnakeCase(filename)
			if isTest {
				return filename + "_test.go"
			}
			return filename + ".go"
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

		if ast.IsExported(name) {
			filename := strcase.SnakeCase(name)
			if isTest {
				return filename + "_test.go"
			}
			return filename + ".go"
		}

		// The problem with unexported functions is that their imports,
		// when merged, would conflict with another function. For example,
		// when using text/template or html/template, math/rand, crypto/rand,
		// or an internal package matching stdlib (internal/crypto).

		isConflicting := func(names []string) bool {
			conflicting := map[string]bool{
				"html/template": true,
				"text/template": true,
				"math/rand":     true,
				"crypto/rand":   true,
			}
			for _, name := range names {
				if ok, _ := conflicting[name]; ok {
					return true
				}
			}
			return false
		}

		if len(t.Imports) > 0 && isConflicting(t.Imports) {
			filename := strcase.SnakeCase(name)
			if isTest {
				return filename + "_test.go"
			}
			return filename + ".go"
		}

		if isTest {
			return "funcs_test.go"
		}

		return "funcs.go"
	}

	var found bool
	for _, def := range defs {
		if cfg.packageName == def.Package {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no such package: %s", cfg.packageName)
	}

	for _, def := range defs {
		if cfg.packageName != def.Package {
			continue
		}

		for _, t := range def.Types {
			name := findShortest(t.Names, t.Name)
			filename := strcase.SnakeCase(name)
			if strings.HasSuffix(t.File, "_test.go") {
				filename = filename + "_test"
			}
			filename = filename + ".go"

			add(filename, t)
		}

		for _, t := range def.Funcs {
			filename := classifyFunc(t)
			add(filename, t)
		}

		for _, t := range def.Vars {
			add("vars.go", t)
		}

		for _, t := range def.Consts {
			add("const.go", t)
		}
	}

	filenames := maps.Keys(files)
	sort.Slice(filenames, func(i int, j int) bool {
		clean := func(s string) string {
			cut := ".go"
			s = s[0 : len(s)-len(cut)]
			cut = "_test"
			s, _ = cutSuffix(s, cut)
			return s
		}
		p1, p2 := filenames[i], filenames[j]
		c1, c2 := clean(p2), clean(p2)
		if c1 == c2 {
			return strings.Compare(p1, p2) < 0
		}
		return strings.Compare(c1, c2) < 0
	})

	if cfg.save {
		return saveLayout(cfg, files, filenames)
	}
	return printLayout(cfg, files, filenames)
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

func cutSuffix(s string, suffix string) (after string, changed bool) {
	if !strings.HasSuffix(s, suffix) {
		return s, false
	}
	return s[:len(s)-len(suffix)], true
}
