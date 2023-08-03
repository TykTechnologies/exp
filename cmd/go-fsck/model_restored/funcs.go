package model

import (
	"go/ast"
	"go/token"
)

func appendIfNotExists(slice []string, element string) []string {
	for _, s := range slice {
		if s == element {
			return slice
		}
	}
	return append(slice, element)
}

func containsOtherTypes(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.Ident:

		if t.Obj != nil && t.Obj.Kind == ast.Typ {
			return true
		}
	case *ast.SelectorExpr:

		return true

	}

	return false
}

func isSelfContainedType(genDecl *ast.GenDecl) bool {
	switch genDecl.Tok {
	case token.TYPE:
		for _, spec := range genDecl.Specs {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				switch t := typeSpec.Type.(type) {
				case *ast.StructType:

					for _, f := range t.Fields.List {
						if containsOtherTypes(f.Type) {
							return false
						}
					}
				case *ast.InterfaceType:

					for _, f := range t.Methods.List {
						if containsOtherTypes(f.Type) {
							return false
						}
					}

				default:
					return false
				}
			} else {

				return false
			}
		}
	case token.VAR, token.CONST:
		for _, spec := range genDecl.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {

				if containsOtherTypes(valueSpec.Type) {
					return false
				}
			} else {

				return false
			}
		}
	default:

		return false
	}

	return true
}
