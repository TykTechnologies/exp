package restore

import (
	"fmt"
	"strings"

	"github.com/stoewer/go-strcase"
)

func toFilename(s string) string {
	s = strings.ReplaceAll(s, "OAuth", "Oauth")
	s = strings.ReplaceAll(s, "CoProcess", "Coprocess")
	s = strcase.SnakeCase(s)
	// hack
	if s == "" {
		return "funcs.go"
	}
	return s + ".go"
}

func toType(s string) (string, bool) {
	// variadic ...
	s = strings.TrimLeft(s, ".")
	// functions
	if strings.HasPrefix(s, "func") {
		return "<func>", false
	}
	// channels
	// implementation gap: anonymous structs
	if strings.Contains(s, " ") {
		ss := strings.SplitN(s, " ", 2)
		if strings.Contains(ss[0], "chan") {
			s = ss[1]
		}
	}
	// maps, slices
	for strings.Contains(s, "]") {
		// Handle:
		// - `[]` V
		// - `map[...]` V
		// - `.*]` V
		ss := strings.SplitN(s, "]", 2)
		s = ss[len(ss)-1]
	}
	// *T to T
	s = strings.TrimLeft(s, "*")
	if yes, _ := builtInTypes[s]; yes {
		return s, false
	}
	return s, true
}

func IsConflicting(names []string) error {
	// The problem with unexported functions is that their imports,
	// when merged, would conflict with another function. For example,
	// when using text/template or html/template, math/rand, crypto/rand,
	// or an internal package matching stdlib (internal/crypto).
	conflicting := map[string]bool{
		"html/template":            true,
		"text/template":            true,
		"math/rand":                true,
		"crypto":                   true,
		"crypto/rand":              true,
		"context":                  true,
		"golang.org/x/net/context": true,
	}
	for _, name := range names {
		clean := name
		if strings.Contains(name, " ") {
			clean = strings.Split(name, " ")[1]
		}
		clean = strings.Trim(clean, `"`)
		if ok, _ := conflicting[clean]; ok {
			return fmt.Errorf("Imports conflict over %s", clean)
		}
	}
	return nil
}

var builtInTypes = map[string]bool{
	"string":      true,
	"int":         true,
	"int8":        true,
	"int16":       true,
	"int32":       true,
	"int64":       true,
	"uint":        true,
	"uint8":       true,
	"uint16":      true,
	"uint32":      true,
	"uint64":      true,
	"uintptr":     true,
	"float32":     true,
	"float64":     true,
	"complex64":   true,
	"complex128":  true,
	"byte":        true,
	"rune":        true,
	"bool":        true,
	"error":       true,
	"interface{}": true,
}
