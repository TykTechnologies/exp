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
		// It's an identifier; check if it's referring to another type/package.
		if t.Obj != nil && t.Obj.Kind == ast.Typ {
			return true // It's referring to another type/package.
		}
	case *ast.SelectorExpr:
		// It's a selector expression; check if it's referring to another package.
		return true
		// Add cases for other types you want to handle.
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
					// It's a struct type, check if it references other types/packages.
					for _, f := range t.Fields.List {
						if containsOtherTypes(f.Type) {
							return false
						}
					}
				case *ast.InterfaceType:
					// It's an interface type, check if it references other types/packages.
					for _, f := range t.Methods.List {
						if containsOtherTypes(f.Type) {
							return false
						}
					}
				// Add other cases for any other self-contained types you want to support,
				// like arrays, slices, etc., or return false if the type is not supported.
				default:
					return false
				}
			} else {
				// There is a non-type spec in the GenDecl (e.g., a variable or constant declaration).
				return false
			}
		}
	case token.VAR, token.CONST:
		for _, spec := range genDecl.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {
				// Check if the variable/constant type references other types/packages.
				if containsOtherTypes(valueSpec.Type) {
					return false
				}
			} else {
				// There is a non-value spec in the GenDecl.
				return false
			}
		}
	default:
		// The GenDecl is not a type, variable, or constant declaration.
		return false
	}

	// All specs are self-contained types, variables, or constants.
	return true
}
