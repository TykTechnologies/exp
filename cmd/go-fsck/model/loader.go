package model

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Load definitions from package located in sourcePath.
func Load(sourcePath string) (*Definition, error) {
	definition := &Definition{}
	fset := token.NewFileSet()

	packages, err := parser.ParseDir(fset, sourcePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, pkg := range packages {
		for filename, file := range pkg.Files {
			definition.Filename = filename

			// Collect imports
			for _, imported := range file.Imports {
				importLiteral := imported.Path.Value
				if imported.Name != nil {
					alias := imported.Name.Name
					importLiteral = alias + " " + importLiteral
				}

				definition.Imports = append(definition.Imports, importLiteral)
			}

			// Walk the AST and collect declarations
			ast.Walk(&declarationCollector{fset: fset, definition: definition}, file)
		}
	}

	return definition, nil
}
