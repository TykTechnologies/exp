package restore

import (
	"strings"

	"github.com/stoewer/go-strcase"
)

func toFilename(s string) string {
	s = strings.ReplaceAll(s, "OAuth", "Oauth")
	s = strings.ReplaceAll(s, "CoProcess", "Coprocess")
	s = strcase.SnakeCase(s)
	return s + ".go"
}

func isConflicting(names []string) bool {
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
		if ok, _ := conflicting[name]; ok {
			return ok
		}
	}
	return false
}

var builtInTypes = map[string]bool{
	"string":     true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uintptr":    true,
	"float32":    true,
	"float64":    true,
	"complex64":  true,
	"complex128": true,
	"byte":       true,
	"rune":       true,
	"bool":       true,
	"error":      true,
}
