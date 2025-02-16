package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/tools/go/packages"
)

// findFieldUsages searches for field accesses in Go files within the current module.
func findFieldUsages(root, structName, fieldName string, verbose bool) error {
	// Tokenizer set
	fset := token.NewFileSet()

	// Load the package with full type information
	cfg := &packages.Config{
		Mode:  packages.LoadSyntax | packages.LoadTypes | packages.NeedTypesInfo,
		Fset:  fset,
		Dir:   root,
		Tests: false,
	}

	// Load all Go packages in the project
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return fmt.Errorf("failed to load packages: %w", err)
	}

	// Iterate through packages
	for _, pkg := range pkgs {
		if pkg.TypesInfo == nil {
			continue // Skip packages without type info
		}

		for _, file := range pkg.Syntax {
			filename := fset.Position(file.Pos()).Filename
			src, err := os.ReadFile(filename)
			if err != nil {
				continue // Skip unreadable files
			}

			// Traverse AST and find field accesses
			ast.Inspect(file, func(n ast.Node) bool {
				// Look for field accesses (e.g., config.Cloud)
				if sel, ok := n.(*ast.SelectorExpr); ok {
					if ident, ok := sel.X.(*ast.Ident); ok && sel.Sel.Name == fieldName {
						// Check if the base identifier refers to the correct struct
						if obj, ok := pkg.TypesInfo.Uses[ident]; ok {
							if named, ok := obj.Type().(*types.Named); ok {
								currentName := named.Obj().Name()
								if verbose {
									fmt.Println(currentName)
								}
								if structName == "_" || currentName == structName {
									// Get position and line content
									pos := fset.Position(sel.Pos())
									line := strings.TrimSpace(getLineFromFile(src, pos.Line))

									// Print the result
									wd, _ := os.Getwd()
									filename = strings.TrimPrefix(filename, wd+"/")
									fmt.Printf("%s %d %s\n", filename, pos.Line, line)
								}
							}
						}
					}
				}
				return true
			})
		}
	}
	return nil
}

// getLineFromFile extracts a specific line from source code.
func getLineFromFile(src []byte, lineNumber int) string {
	lines := strings.Split(string(src), "\n")
	if lineNumber > 0 && lineNumber <= len(lines) {
		return strings.TrimSpace(lines[lineNumber-1])
	}
	return ""
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: discover <structName> <fieldName>")
		os.Exit(1)
	}

	structName := os.Args[1]
	fieldName := os.Args[2]
	verbose := slices.Contains(os.Args, "-v")
	root := "." // Start from the current directory

	err := findFieldUsages(root, structName, fieldName, verbose)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
