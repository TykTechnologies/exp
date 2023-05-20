package extract

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

func TrimSpace(in *ast.CommentGroup) string {
	if in != nil {
		return strings.TrimSpace(in.Text())
	}
	return ""
}

var TypeName = getTypeDeclarations

func getTypeDeclaration(ident *ast.Ident) string {
	return ident.Name
}

func getTypeDeclarations(expr *ast.TypeSpec) string {
	return getTypeDeclarationsForExpr(expr.Type)
}

func getTypeDeclarationsForPointerType(starExpr *ast.StarExpr) string {
	return "*" + getTypeDeclarationsForExpr(starExpr.X)
}

func getTypeDeclarationsForArrayType(arrayType *ast.ArrayType) string {
	var declaration string
	switch obj := arrayType.Elt.(type) {
	case *ast.Ident:
		declaration = getTypeDeclaration(obj)
	case *ast.ArrayType:
		declaration = getTypeDeclarationsForArrayType(obj)
	case *ast.MapType:
		declaration = getTypeDeclarationsForMapType(obj)
	case *ast.StarExpr:
		declaration = getTypeDeclarationsForPointerType(obj)
	default:
		declaration = fmt.Sprintf("%#v", arrayType.Elt)
	}

	return fmt.Sprintf("[]%s", declaration)
}

func getTypeDeclarationsForMapType(mapType *ast.MapType) string {
	var sb strings.Builder
	sb.WriteString("map[")
	sb.WriteString(getTypeDeclaration(mapType.Key.(*ast.Ident)))
	sb.WriteString("]")
	sb.WriteString(getTypeDeclaration(mapType.Value.(*ast.Ident)))
	return sb.String()
}

func getTypeDeclarationsForExpr(expr ast.Expr) string {
	switch expr := expr.(type) {
	case *ast.SelectorExpr:
		return getTypeDeclarationsForExpr(expr.X) + "." + getTypeDeclaration(expr.Sel)
	case *ast.Ident:
		return getTypeDeclaration(expr)
	case *ast.ArrayType:
		return getTypeDeclarationsForArrayType(expr)
	case *ast.MapType:
		return getTypeDeclarationsForMapType(expr)
	default:
		return fmt.Sprintf("%#v", expr)
	}
}

func jsonTag(tag string) string {
	return reflect.StructTag(tag).Get("json")
}
