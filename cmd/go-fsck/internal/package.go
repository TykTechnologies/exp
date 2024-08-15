package internal

import (
	"os"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

// ListPackages returns a slice of local packages in the specified root directory.
// The second argument is either `.` or `./...` (recursive).
func ListPackages(rootPath string, pattern string) ([]model.Package, error) {
	if err := os.Chdir(rootPath); err != nil {
		return nil, err
	}

	packages, err := listPackages(pattern)
	if err != nil || len(packages) == 0 {
		return nil, err
	}

	result := cleanPackages(packages)

	return result, nil
}

func cleanPackages(pkgs []*packages.Package) []model.Package {
	results := make([]model.Package, 0, len(pkgs))
	first := pkgs[0].PkgPath
	for _, pkg := range pkgs {
		cleanPath := "." + strings.TrimPrefix(pkg.PkgPath, first)
		var testPackage bool
		if strings.Contains(cleanPath, ".test") {
			continue
		}
		if strings.HasSuffix(cleanPath, "_test") {
			cleanPath = cleanPath[:len(cleanPath)-5]
			testPackage = true
		}
		result := model.Package{
			Package:     pkg.Name, //filepath.Base(pkg.PkgPath),
			ImportPath:  pkg.PkgPath,
			Path:        cleanPath,
			TestPackage: testPackage,
		}
		results = append(results, result)
	}
	return results
}

func listPackages(pattern string) ([]*packages.Package, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedModule,
		Tests: true,
	}

	return packages.Load(cfg, pattern)
}
