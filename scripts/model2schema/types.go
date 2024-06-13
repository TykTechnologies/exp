package main

import "strings"

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
