package modfile

import (
	"fmt"
	"strings"

	"golang.org/x/mod/modfile"
)

// Parse parses the contents of a go.mod file and returns the filtered modules
// that match the given match.
//
// Arguments:
// - contents: The contents of the go.mod file as a []byte.
// - match: The match to filter the module paths.
func Parse(contents []byte, match string) ([]*modfile.Require, error) {
	f, err := modfile.Parse("go.mod", contents, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.mod: %w", err)
	}

	var matchingModules []*modfile.Require
	for _, req := range f.Require {
		if req.Indirect {
			continue
		}

		if match == "" || strings.Contains(req.Mod.Path, match) {
			matchingModules = append(matchingModules, req)
		}
	}

	return matchingModules, nil
}
