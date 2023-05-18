package structs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
)

// Extract package structs
func Extract(filepath string, ignoreFiles ...string) (StructList, error) {
	ignoreList := make(map[string]bool)
	for _, file := range ignoreFiles {
		ignoreList[file] = true
	}

	// filter skips explicitly ignored files, and tests files
	filter := func(fInfo os.FileInfo) bool {
		return !(ignoreList[fInfo.Name()] || strings.HasSuffix(fInfo.Name(), "_test.go"))
	}

	fileSet := token.NewFileSet()
	pkgs, err := parser.ParseDir(fileSet, path.Dir(filepath), filter, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if len(pkgs) != 1 {
		return nil, fmt.Errorf("expecting single go package, got %d", len(pkgs))
	}

	requiredPkgName := func() string {
		// Get first package name
		var pkgName string
		for pkgName, _ = range pkgs {
			break
		}
		return pkgName
	}()

	if _, ok := pkgs[requiredPkgName]; !ok {
		return nil, fmt.Errorf("required package %q", requiredPkgName)
	}

	p := newObjParser(fileSet, pkgs[requiredPkgName])
	p.parseGlobalStructs()

	sort.Stable(p.info)

	if p.errList.Empty() {
		return p.info, nil
	}
	return p.info, p.errList
}

type objParser struct {
	fileset *token.FileSet

	info    StructList
	globals map[string]*ast.TypeSpec // exported global types
	visited map[string]*StructInfo   // avoid re-visiting struct type
	pkg     *ast.Package
	errList *FieldDocError
}

func newObjParser(fileset *token.FileSet, pkg *ast.Package) *objParser {
	p := &objParser{
		fileset: fileset,
		info:    StructList{},
		visited: map[string]*StructInfo{},
		globals: map[string]*ast.TypeSpec{},
		errList: &FieldDocError{},
		pkg:     pkg,
	}
	p.fillGlobalTypes()
	return p
}

func (p *objParser) fillGlobalTypes() {
	for _, fileObj := range p.pkg.Files {
		if fileObj.Scope == nil {
			continue
		}
		for objName, obj := range fileObj.Scope.Objects {
			if decl, ok := obj.Decl.(*ast.TypeSpec); ok && ast.IsExported(objName) {
				p.globals[objName] = decl
			}
		}
	}
}

func (p *objParser) parseGlobalStructs() {
	// Every global is an *ast.TypeSpec
	for rootName, g := range p.globals {
		switch obj := g.Type.(type) {
		case *ast.StructType:
			rootStructInfo := &StructInfo{
				Name:      rootName,
				Doc:       TrimSpace(g.Doc),
				Comment:   TrimSpace(g.Comment),
				fileSet:   p.fileset,
				structObj: obj,
			}
			p.parse(rootName, rootName, rootStructInfo)
		default:
			ident := extractIdentFromExpr(g.Type)
			info := &StructInfo{
				Name:    rootName,
				Doc:     TrimSpace(g.Doc),
				Comment: TrimSpace(g.Comment),
				Type:    ident.String(),
				fileSet: p.fileset,
			}
			p.visited[rootName] = info
			p.info.append(info)
		}
	}
}

func (p *objParser) parse(goPath, name string, structInfo *StructInfo) {
	if p.visited[name] != nil {
		return
	}

	p.visited[name] = structInfo
	p.info.append(structInfo)

	for _, field := range structInfo.structObj.Fields.List {
		pos := structInfo.fileSet.Position(field.Pos())
		filePos := path.Base(pos.String())

		var goName string
		if len(field.Names) > 0 {
			goName = field.Names[0].Name
		}

		// ignored field.
		if goName == "_" {
			continue
		}
		// unexposed field.
		if !ast.IsExported(goName) {
			continue
		}

		if goName == "" {
			p.errList.WriteError(fmt.Sprintf("[%s] unidentified field in %s", filePos, goPath))
			continue
		}

		// fmt.Println("goName", goName)

		ident := extractIdentFromExpr(field.Type)
		if ident == nil {
			if len(field.Names) > 0 {
				// inline fields
				ident = extractIdentFromExpr(p.globals[goName])
			}
		}
		if ident == nil {
			ident = ast.NewIdent("any")
		}

		tagValue := ""
		if field.Tag != nil {
			tagValue = string(field.Tag.Value)
			tagValue = strings.Trim(tagValue, "`")
		}

		jsonName := jsonTag(tagValue)
		if jsonName == "" {
			// fields without json tag encode to field name
			jsonName = goName
		}

		fieldInfo := &FieldInfo{
			Doc:     TrimSpace(field.Doc),
			Comment: TrimSpace(field.Comment),

			Name: goName,
			Path: goPath + "." + goName,
			Type: ident.String(),
			Tag:  tagValue,

			JSONName: jsonName,

			IsArray: isExprArray(field.Type),

			fileSet: structInfo.fileSet,
		}
		// p.parseNestedObj(ident.Name, fieldInfo)

		structInfo.Fields = append(structInfo.Fields, fieldInfo)
	}
}

func (p *objParser) parseNestedObj(name string, field *FieldInfo) {
	if p.globals[name] != nil {
		switch obj := p.globals[name].Type.(type) {
		case *ast.StructType:
			newInfo := &StructInfo{
				structObj: obj,
				fileSet:   field.fileSet,
				Name:      name,
			}
			p.parse(name, name, newInfo)

		case *ast.ArrayType:
			typeName := extractIdentFromExpr(obj).Name
			field.Type = "[]" + typeName
			field.IsArray = true
			if structObj, ok := p.globals[typeName].Type.(*ast.StructType); ok {
				newInfo := &StructInfo{
					structObj: structObj,
					fileSet:   field.fileSet,
					Name:      typeName,
				}
				p.parse(typeName, typeName, newInfo)
			}

		case *ast.MapType:
			typeName := extractIdentFromExpr(obj).Name
			field.MapKey = extractIdentFromExpr(obj.Key).Name
			field.Type = fmt.Sprintf("map[%s]%s", typeName, field.MapKey)

			if structObj, ok := p.globals[typeName].Type.(*ast.StructType); ok {
				newInfo := &StructInfo{
					structObj: structObj,
					fileSet:   field.fileSet,
					Name:      typeName,
				}
				p.parse(typeName, typeName, newInfo)
			}
		}
	}
}

func extractIdentFromExpr(expr any) *ast.Ident {
	switch objType := expr.(type) {
	case *ast.StarExpr:
		return extractIdentFromExpr(objType.X)

	case *ast.Ident:
		return objType

	case *ast.MapType:
		return extractIdentFromExpr(objType.Value)

	case *ast.ArrayType:
		return extractIdentFromExpr(objType.Elt)

	case *ast.InterfaceType:
		return ast.NewIdent("any")

	case *ast.SelectorExpr:
		return ast.NewIdent("object")
	}
	return nil
}

func isExprArray(expr ast.Expr) bool {
	_, ok := expr.(*ast.ArrayType)
	return ok
}

func jsonTag(tag string) string {
	return reflect.StructTag(tag).Get("json")
}
