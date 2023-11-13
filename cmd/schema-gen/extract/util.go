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

func jsonTag(tag string) string {
	return reflect.StructTag(tag).Get("json")
}

func write(cfg *options) error {
	opts := NewExtractOptions(cfg)

	sts, err := Extract(cfg.sourcePath, opts)
	if err != nil {
		return err
	}

	return dump(cfg, sts)
}

func dump(cfg *options, data interface{}) error {
	var (
		filename = cfg.outputFile
		pretty   = cfg.prettyJSON

		dataBytes []byte
		err       error
	)

	if pretty {
		dataBytes, err = json.MarshalIndent(data, "", "  ")
	} else {
		dataBytes, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(filename, dataBytes, 0644) //nolint:gosec
}
