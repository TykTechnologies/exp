package restore

import (
	"fmt"
	"os"
	"strings"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

func saveLayout(cfg *options, files map[string][]*model.Declaration, filenames []string, imports []string) error {
	for _, filename := range filenames {
		if cfg.removeTests && strings.HasSuffix(filename, "_test.go") {
			continue
		}
		decls := files[filename]

		// Collect sources
		lines := []string{"package " + cfg.packageName}
		if len(imports) != 0 {
			lines = append(lines, "", "import (")
			for _, i := range imports {
				lines = append(lines, strings.TrimSpace(i))
			}
			lines = append(lines, ")")
		}

		for _, decl := range decls {
			if decl.Source != "" {
				lines = append(lines, "", decl.Source)
				continue
			}
			return fmt.Errorf("Missing source for object, %#v", decl)
		}

		if len(lines) > 0 {
			body := []byte(strings.Join(lines, "\n"))
			if err := os.WriteFile(filename, body, 0644); err != nil {
				return fmt.Errorf("error saving: %w", err)
			}
		}
	}

	return nil
}
