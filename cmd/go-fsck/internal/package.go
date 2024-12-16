package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

// ListPackages returns a slice of local packages in the specified root directory.
// The second argument is either `.` or `./...` (recursive).
func ListPackages(rootPath string, pattern string) ([]*model.Package, error) {
	if err := os.Chdir(rootPath); err != nil {
		return nil, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	packages, err := listPackages(pattern)
	if err != nil || len(packages) == 0 {
		return nil, err
	}

	result := cleanPackages(packages, cwd)

	return result, nil
}

func cleanPackages(pkgs []*packages.Package, workDir string) []*model.Package {
	results := make([]*model.Package, 0, len(pkgs))
	seen := make(map[string]*packages.Package)

	for _, pkg := range pkgs {
		// Filters out tests packages. We only care about the Dir's
		// so we don't want to duplicate folders for tests.
		if strings.Contains(pkg.ID, ".test") {
			continue
		}

		testPackage := strings.HasSuffix(pkg.PkgPath, "_test")

		cleanPath := "." + strings.TrimPrefix(pkg.Dir, workDir)

		pkgKey := cleanPath + "-" + fmt.Sprint(testPackage)
		fmt.Println(pkgKey)

		if val, ok := seen[pkgKey]; ok {

			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")

			enc.Encode(val)
			enc.Encode(pkg)

			//panic("seen package twice: " + pkgKey)
		}

		result := &model.Package{
			Package:     pkg.Name, //filepath.Base(pkg.PkgPath),
			ImportPath:  pkg.PkgPath,
			Path:        cleanPath,
			TestPackage: testPackage,
		}

		seen[pkgKey] = pkg

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
