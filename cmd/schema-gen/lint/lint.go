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
				errs.Append(lintField(cfg, decl, typeDecl, fieldDecl)...)
			}
		}
	}

	if errs.Empty() {
		return nil
	}
	return errs
}

func lintField(cfg *options, decl *DeclarationInfo, typeDecl *TypeInfo, fieldDecl *FieldInfo) []string {
	result := make([]string, 0, len(cfg.rules))

	var (
		name  = fieldDecl.Name
		doc   = fieldDecl.Doc
		field = fieldDecl.Path
	)

	if cfg.verbose {
		fmt.Printf("// %s\n", doc)
		fmt.Printf("%s\n\n", field)
	}

	for _, rule := range cfg.rules {
		result = append(result, validateRule(rule, field, name, doc))
	}
	return result
}

func validateRule(rule, field, name, doc string) string {
	switch rule {
	case "require-field-comment":
		if doc == "" {
			return fmt.Sprintf("[%s] No comment on exposed field.", field)
		}
	case "require-dot-or-backtick":
		if doc == "" {
			return ""
		}
		if strings.HasSuffix(doc, ".") || strings.HasSuffix(doc, "`") {
			return ""
		}

		return fmt.Sprintf("[%s] Field comment must end with dot or `.\nGot:  %s\n", field, doc)
	case "require-field-prefix":
		if doc == "" {
			return ""
		}
		prefixes := []string{
			fmt.Sprintf("%s ", name),
			fmt.Sprintf("%s: ", name),
		}
		hasPrefix := func(doc string, prefixes []string) bool {
			for _, prefix := range prefixes {
				if strings.HasPrefix(doc, prefix) {
					return true
				}
			}
			return false
		}(doc, prefixes)

		if !hasPrefix {
			return fmt.Sprintf("[%s] Comment must start with field name.\nGot:  %s\n", field, doc)
		}
	default:
	}
	return ""
}
