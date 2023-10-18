package model

import (
	"fmt"
	"os"
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
func (i StringSet) Map() map[string]string {
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
			fmt.Fprintf(os.Stderr, "WARN: Import path conflict: %s, %s (prev) != %s (new)\n", short, val, long)
		}

		result[short] = long
	}

	return result
}
