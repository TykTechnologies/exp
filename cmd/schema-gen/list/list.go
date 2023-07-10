package list

import (
	"encoding/json"
	"fmt"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

func listStructures(cfg *options) error {
	fmt.Println(cfg.inputFile)

	pkgInfos, err := model.Load(cfg.inputFile)
	if err != nil {
		return fmt.Errorf("Error loading package info: %w", err)
	}

	for _, pkgInfo := range pkgInfos {
		pkg := listPackage(cfg, pkgInfo)
		printPackage(cfg, pkg)
	}

	return nil
}

type PackageFile struct {
	Name string `json:"name"`
	Path string `json:"path"`

	PackageInfo *PackageInfo `json:"-"`
}

// PackageFileMap key is symbol Path for the struct ordered into a file.
type PackageFileMap map[string]*PackageFile

func printPackage(cfg *options, files PackageFileMap) {
	for _, v := range files {
		s, _ := json.Marshal(v)
		fmt.Println(string(s))
	}
}

func listPackage(cfg *options, pkgInfo *PackageInfo) PackageFileMap {
	files := PackageFileMap{}

	for _, decls := range pkgInfo.Declarations {
		for _, typeDecl := range decls.Types {
			// Generate filename from type name
			name := format.SnakeCase(typeDecl.Name) + ".go"
			file, ok := files[name]
			if !ok {
				file = &PackageFile{
					Name: name,
					Path: typeDecl.Name,
				}
				files[name] = file
			}
		}
	}

	return files
}
