package lint

import (
	"fmt"
	"strings"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

func lint(cfg *options) error {
	fmt.Println(cfg.inputFile)

	pkgInfo, err := model.LoadPackageInfo(cfg.inputFile)
	if err != nil {
		return fmt.Errorf("Error loading package info: %w", err)
	}

	errs := NewFieldDocError()

	// Dump out declarations
	for _, decl := range pkgInfo.Declarations {
		for _, typeDecl := range decl.Types {
			for _, fieldDecl := range typeDecl.Fields {
				errs.Append(lintField(cfg, decl, typeDecl, fieldDecl))
			}
		}
	}

	if errs.Empty() {
		return nil
	}
	return errs
}

func lintField(cfg *options, decl *DeclarationInfo, typeDecl *TypeInfo, fieldDecl *FieldInfo) string {
	for _, rule := range cfg.rules {
		switch rule {
		case "godoc-fields":
			name, doc, field := fieldDecl.Name, fieldDecl.Doc, fieldDecl.Path

			if !strings.HasSuffix(doc, ".") {
				return fmt.Sprintf("[%s] Field comment must end with dot.", field)
			}

			findPrefix := fmt.Sprintf("%s ", name)
			if !strings.HasPrefix(doc, findPrefix) {
				return fmt.Sprintf("[%s] Comment must start with field name.\nGot:  %s,\nWant: %s", field, doc, findPrefix)
			}

			if cfg.verbose {
				fmt.Printf("// %s\n", fieldDecl.Doc)
				fmt.Printf("%s\n\n", fieldDecl.Name)
			}
			return ""
		}
	}
	return ""
}
