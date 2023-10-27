package model

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type StringSet map[string][]string

func (i *StringSet) Add(key, lit string) {
	data := *i
	if data == nil {
		data = make(StringSet)
	}
	if set, ok := data[key]; ok {
		if slices.Contains(set, lit) {
			return
		}
		data[key] = append(set, lit)
		return
	}
	data[key] = []string{lit}
	*i = data
}

func (i StringSet) All() []string {
	result := []string{}
	for _, set := range i {
		result = append(result, set...)
	}
	return result
}

func (i StringSet) Get(key string) []string {
	val, _ := i[key]
	if val != nil {
		sort.Strings(val)
	}
	return val
}

// Map returns a map with the short package name as the key
// and the full import path as the value.
func (i StringSet) Map() (map[string]string, []error) {
	warnings := []error{}
	warningSeen := map[string]bool{}

	result := map[string]string{}
	imports := i.All()

	for _, imported := range imports {
		var short, long string

		// aliased package
		if strings.Contains(imported, " ") {
			line := strings.Split(imported, " ")
			short, long = line[0], strings.Trim(line[1], `"`)
		} else {
			long = strings.Trim(imported, `"`)
			short = path.Base(long)
		}

		val, ok := result[short]

		if ok && val != long {
			warning := "Import conflict for %s, "
			// Sort val/long so shorter is left hand side
			if len(val) < len(long) {
				warning += val + " != " + long
			} else {
				warning += long + " != " + val
			}
			if _, seen := warningSeen[warning]; !seen {
				warningSeen[warning] = true
				warnings = append(warnings, fmt.Errorf(warning, short))
			}
		}

		result[short] = long
	}

	return result, warnings
}
