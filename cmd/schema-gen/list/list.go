package list

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
	"golang.org/x/exp/slices"
)

func listStructures(cfg *options) error {
	matches, err := filepath.Glob(cfg.inputFile)
	if err != nil {
		return err
	}

	sort.Strings(matches)

	all := []*TypeDeclaration{}

	find := func(ts []*TypeDeclaration, p string) *TypeDeclaration {
		for _, t := range ts {
			if t.Path == p {
				return t
			}
		}
		return nil
	}

	for _, filename := range matches {
		pkgInfos, err := model.Load(filename)
		if err != nil {
			return fmt.Errorf("Error loading package info for %s: %w", filename, err)
		}

		for _, pkgInfo := range pkgInfos {
			pkg := listSymbols(cfg, pkgInfo)
			for _, sym := range pkg {
				got := find(all, sym.Path)
				if got != nil {
					got.Doc = sym.Doc
					got.Seen = append(got.Seen, filename)
					continue
				}
				sym.Seen = []string{filename}
				all = append(all, sym)
			}
		}
	}

	for _, sym := range all {
		sym.Versions = sym.GetVersions(sym.Seen)
	}

	return printSymbols(cfg, all)
}

type TypeDeclaration struct {
	Seen []string `json:"-"`

	Versions []string `json:"versions"`

	Path string `json:"path"`
	Type string `json:"type"`
	Tag  string `json:"tag"`
	Doc  string `json:"doc"`
}

var (
	majorVer = regexp.MustCompile(`v([0-9]+\.0\.0)`)
	minorVer = regexp.MustCompile(`v([0-9]+\.[0-9]+)\.`)
	patchVer = regexp.MustCompile(`v([0-9]+\.[0-9]+\.[0-9]+)`)
)

// GetVersions extracts version info from Seen. It reports the first
// minor versions where a type declaration has been seen.
func (t *TypeDeclaration) GetVersions(from []string) []string {
	majors := []string{}
	minors := []string{}
	patches := []string{}

	for _, seen := range from {
		patch := patchVer.FindString(seen)
		minor := minorVer.FindString(seen)

		major := majorVer.FindString(seen)
		if major != "" {
			parts := strings.Split(major, ".")
			major = parts[0]
			majors = append(majors, major)
			minors = append(minors, minor)
			patches = append(patches, patch)
		} else {
			parts := strings.Split(patch, ".")
			major = parts[0]
		}

		// If a field exists in vX.0.0 don't list more
		// versions after that version.
		if slices.Contains(majors, major) {
			continue
		}

		if !slices.Contains(minors, minor) {
			minors = append(minors, minor)
			patches = append(patches, patch)
		}
	}

	if len(patches) > 0 && strings.HasSuffix(patches[0], ".0") {
		return []string{patches[0]}
	}

	return patches
}

func (t *TypeDeclaration) String() string {
	if t.Tag != "" {
		return t.Path + " " + t.Type + " " + t.Tag + " " + fmt.Sprint(t.Versions)
	}
	return t.Path + " " + t.Type + " " + fmt.Sprint(t.Versions)
}

// PackageFileMap key is symbol Path for the struct ordered into a file.
type PackageFileMap map[string]*TypeDeclaration

func printSymbols(cfg *options, symbols []*TypeDeclaration) error {
	if cfg.json || cfg.prettyJSON {
		var (
			out []byte
			err error
		)
		if cfg.prettyJSON {
			out, err = json.MarshalIndent(symbols, "", "  ")
		} else {
			out, err = json.Marshal(symbols)
		}
		if err != nil {
			return err
		}

		fmt.Println(string(out))
		return nil
	}

	for _, symbol := range symbols {
		fmt.Println(symbol)
	}
	return nil
}

func listSymbols(cfg *options, pkgInfo *PackageInfo) []*TypeDeclaration {
	files := []*TypeDeclaration{}

	for _, decls := range pkgInfo.Declarations {
		for _, typeDecl := range decls.Types {
			for _, field := range typeDecl.Fields {
				files = append(files, &TypeDeclaration{
					Path: field.Path,
					Type: field.Type,
					Tag:  field.Tag,
					Doc:  field.Doc,
				})
			}
		}
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Path < files[j].Path
	})

	return files
}
