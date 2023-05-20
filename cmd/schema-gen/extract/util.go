package extract

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"os"
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
	case *ast.SelectorExpr:
		declaration = getTypeDeclarationsForExpr(obj.X) + "." + getTypeDeclaration(obj.Sel)
	default:
		declaration = fmt.Sprintf("%#v", arrayType.Elt)
	}

	return fmt.Sprintf("[]%s", declaration)
}

func getTypeDeclarationsForMapType(mapType *ast.MapType) string {
	keyIdent, ok := mapType.Key.(*ast.Ident)
	if !ok {
		fmt.Println("WARN: unsupported key type in map")
		return "any"
	}
	keyType := getTypeDeclaration(keyIdent)

	valueIdent, ok := mapType.Value.(*ast.Ident)
	if !ok {
		return fmt.Sprintf("map[%s]interface{}", keyType)
	}
	valueType := getTypeDeclaration(valueIdent)

	return fmt.Sprintf("map[%s]%s", keyType, valueType)
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
	case *ast.InterfaceType:
		return ""
	case *ast.StructType:
		return "struct{}"
	case *ast.StarExpr:
		return getTypeDeclarationsForExpr(expr.X)
	default:
		return fmt.Sprintf("%#v", expr)
	}
}

func deduplicate(input []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(input))

	// Iterate over the input slice
	for _, str := range input {
		// Check if the element has been seen before
		if _, ok := seen[str]; !ok {
			// Add the element to the result slice
			result = append(result, str)
			// Mark the element as seen
			seen[str] = true
		}
	}

	return result
}

func jsonTag(tag string) string {
	return reflect.StructTag(tag).Get("json")
}

func write(filename string, inputPackage string) error {
	sts, err := Extract(inputPackage)
	if err != nil {
		return err
	}

	return dump(filename, sts)
}

func dump(filename string, data interface{}) error {
	println(filename)

	dataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, dataBytes, 0644) //nolint:gosec
}
