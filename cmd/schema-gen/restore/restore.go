package restore

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

func restore(cfg *options) error {
	// unpack options into scope
	var (
		inputFile  = cfg.inputFile
		outputFile = cfg.outputFile
	)

	pkgInfo, err := model.LoadPackageInfo(inputFile)
	if err != nil {
		return fmt.Errorf("Error loading package info: %w", err)
	}

	body, err := restorePackageInfo(pkgInfo, cfg)
	if err != nil {
		return err
	}

	fmt.Println(outputFile)
	return os.WriteFile(outputFile, body, 0644)
}

func restorePackageInfo(pkgInfo *model.PackageInfo, cfg *options) ([]byte, error) {
	var output bytes.Buffer

	output.WriteString("package " + cfg.packageName + "\n\n")

	if len(pkgInfo.Imports) > 0 {
		output.WriteString("import (")
		for _, importLiteral := range pkgInfo.Imports {
			output.WriteString("\n\t" + importLiteral)
		}
		output.WriteString("\n)\n\n")
	}

	// Dump out filedocs
	for _, decl := range pkgInfo.Declarations {
		if decl.FileDoc != "" {
			output.WriteString("/*\n" + decl.FileDoc + "\n*/")
		}
	}

	// Dump out declarations
	for _, decl := range pkgInfo.Declarations {
		printDoc(&output, decl.Doc)
		output.WriteString("\ntype (")
		for _, typeDecl := range decl.Types {
			printDoc(&output, typeDecl.Doc)

			// Generic type declaration
			if typeDecl.Type != "" {
				// Type declaration
				output.WriteString(fmt.Sprintf("\n%s = %s\n", typeDecl.Name, typeDecl.Type))
				continue
			}

			// Struct type declaration
			output.WriteString(fmt.Sprintf("\n%s struct {", typeDecl.Name))
			for _, fieldDecl := range typeDecl.Fields {
				printDoc(&output, fieldDecl.Doc)
				// Field declaration
				output.WriteString(fmt.Sprintf("\n%s %s `%s`", fieldDecl.Name, fieldDecl.Type, fieldDecl.Tag))
			}
			output.WriteString("\n}\n")
		}
		output.WriteString(")\n")
	}

	contents, err := format.Source(output.Bytes())
	if err != nil {
		fmt.Println("Error formatting source:", err)
		return output.Bytes(), nil
	}
	return contents, nil
}

func printDoc(output *bytes.Buffer, comment string) {
	if comment == "" {
		return
	}

	lines := strings.Split(comment, "\n")
	for _, line := range lines {
		line = " " + strings.TrimSpace(line)
		output.WriteString("\n//" + line)
	}
}
