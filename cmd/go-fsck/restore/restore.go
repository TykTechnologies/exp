package restore

import (
	"fmt"
	"go/ast"
	"os"
	"sort"
	"strings"

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
		filename, explicit := strings.CutPrefix(filename, "!")

		filename, _ = cutSuffix(filename, "_test.go")
		filename, _ = cutSuffix(filename, ".go")
		fileTest := filename + "_test.go"
		filename = filename + ".go"

		for _, t := range decls {
			dest := filename
			switch {
			case explicit:
			default:
				if isTest := strings.HasSuffix(t.File, "_test.go"); isTest {
					dest = fileTest
				}
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
			return filename
		}

		// Functions can return (T, error), T, error...
		// bind those to the type. This catches constructors
		// based on the return type.

		firstArg := map[string]string{
			"http.ResponseWriter": "http_handlers.go",
			"http.Handler":        "http_handlers.go",
			"http.":               "http_util.go",
			"jwt.":                "jwt.go",
			"tls.":                "tls.go",
			"user.":               "user.go",
			"testing.":            "testing.go",
		}

		if len(t.Arguments) > 0 {
			first := strings.TrimLeft(t.Arguments[0], `*`)
			for arg, filename := range firstArg {
				if strings.HasPrefix(first, arg) {
					return filename
				}
			}

			// Group by first argument type
			if !strings.Contains(first, ".") {
				if filename, ok := findFile(strings.TrimLeft(first, `*`)); ok {
					return filename
				}
			}
		}

		if len(t.Returns) > 0 {
			first := strings.TrimLeft(t.Returns[0], `*`)
			for arg, filename := range firstArg {
				if strings.HasPrefix(first, arg) {
					return filename
				}
			}

			// Group by first return type
			if !strings.Contains(first, ".") {
				if filename, ok := findFile(strings.TrimLeft(first, `*`)); ok {
					return filename
				}
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

		if ast.IsExported(name) {
			filename := toFilename(name)
			return filename
		}

		if len(t.Imports) > 0 && isConflicting(t.Imports) {
			fmt.Println("WARN: possible conflict over", name)
			return toFilename(name)
		}

		fmt.Println(t.Signature)

		return "funcs.go"
	}

	var found bool
	var def *model.Definition

	for _, def = range defs {
		if cfg.packageName == def.Package {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no such package: %s", cfg.packageName)
	}

	for _, t := range def.Types {
		name := findShortest(t.Names, t.Name)
		add(toFilename(name), t)
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

	m := model.DeclarationList{}
	for filename, fd := range files {
		count, total := 0, 0
		for _, f := range fd {
			if f.Kind == model.TypeKind {
				count++
			}
			total++
		}
		if count == total {
			m.Append(fd...)
			delete(files, filename)
		}
	}
	for _, t := range m {
		if isConflicting(t.Imports) {
			add("model_rich.go", t)
		}
		add("model.go", t)
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
