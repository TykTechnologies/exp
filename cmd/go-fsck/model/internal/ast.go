package internal

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

func CommentedNode(file *ast.File, node any) *printer.CommentedNode {
	return &printer.CommentedNode{
		Node:     node,
		Comments: file.Comments,
	}
}

func PrintSource(val *printer.CommentedNode, fset *token.FileSet, out io.Writer) error {
	return printer.Fprint(out, fset, val)
}

func IsSelfContainedType(genDecl *ast.GenDecl) bool {
	switch genDecl.Tok {
	case token.TYPE:
		for _, spec := range genDecl.Specs {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				switch t := typeSpec.Type.(type) {
				case *ast.StructType:
					// It's a struct type, check if it references other types/packages.
					for _, f := range t.Fields.List {
						if ContainsOtherTypes(f.Type) {
							return false
						}
					}
				case *ast.InterfaceType:
					// It's an interface type, check if it references other types/packages.
					for _, f := range t.Methods.List {
						if ContainsOtherTypes(f.Type) {
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
				if ContainsOtherTypes(valueSpec.Type) {
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

func ContainsOtherTypes(expr ast.Expr) bool {
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
