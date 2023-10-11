package model

import (
	"go/ast"
)

func collectFuncReferences(funcDecl *ast.FuncDecl) map[string][]string {
	imports := make(map[string][]string)

	// Traverse the function body and look for package identifiers.
	ast.Inspect(funcDecl.Body, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.SelectorExpr:
			// If it's a SelectorExpr, get the leftmost identifier which is the package name.
			if ident, ok := n.X.(*ast.Ident); ok {
				pkgName := ident.Name

				if ident.Obj != nil {
					if ident.Obj.Kind != ast.Pkg {
						// pkgName is not a package
						return true
					}
				}

				selName := n.Sel.Name
				if pkgName != "internal" && ast.IsExported(selName) {
					imports[pkgName] = appendIfNotExists(imports[pkgName], selName)
				}
			}
		case *ast.Ident:
			// If it's an identifier, it might be a package name.
			if obj := n.Obj; obj != nil && obj.Kind == ast.Pkg {
				pkgName := n.Name
				imports[pkgName] = nil // No specific symbol, just mark the package as imported.
			}
		}

		return true
	})

	return imports
}
