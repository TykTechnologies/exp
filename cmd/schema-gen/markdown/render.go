package markdown

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

func render(kind, inputFile, outputFile, rootElement string) error {
	pkgInfo, err := model.Load(inputFile)
	if err != nil {
		return fmt.Errorf("Error loading package info for %s: %w", inputFile, err)
	}

	order := pkgInfo.Declarations.GetOrder(rootElement)

	switch kind {
	case "markdown":
		body, err := renderMarkdown(sanitize(pkgInfo.Declarations), order)
		if err != nil {
			return err
		}
		return os.WriteFile(outputFile, body, 0644)
	}
	return fmt.Errorf("Renderer %q not implemented", kind)
}

func renderMarkdown(schema model.DeclarationList, order []string) ([]byte, error) {
	output := new(bytes.Buffer)
	decls := schema.Find(order)
	for _, decl := range decls {
		if err := renderMarkdownType(output, decl); err != nil {
			return nil, err
		}
	}
	return output.Bytes(), nil
}

func renderMarkdownType(w io.Writer, decl *model.TypeInfo) error {
	fmt.Fprintf(w, "# %s\n\n", decl.Name)
	if decl.Doc != "" {
		fmt.Fprintf(w, "%s\n\n", decl.Doc)
	}
	renderMarkdownFields(w, decl)
	return nil
}

func renderMarkdownFields(w io.Writer, decl *model.TypeInfo) {
	if len(decl.Fields) == 0 {
		fmt.Fprintf(w, "No exposed fields available.\n\n")
		return
	}

	for _, field := range decl.Fields {
		fmt.Fprintf(w, "**%s** (JSON: `%s`)\n\n", field.Name, strings.Split(field.JSONName, ",")[0])
		fmt.Fprintf(w, "%s\n\n", field.Doc)
		if field.Comment != "" {
			fmt.Fprintf(w, "> %s\n\n", field.Comment)
		}
	}
}

func sanitize(x model.DeclarationList) model.DeclarationList {
	// everything is a pointer, so no weird handling for modifiying
	// the slice values
	for _, decl := range x {
		// Move declaration doc comment into type decl if there
		// is only one and the comment is empty. Weird.
		if len(decl.Types) == 1 && decl.Types[0].Doc == "" {
			decl.Types[0].Doc = decl.Doc
			decl.Doc = ""
		}
	}
	return x
}
