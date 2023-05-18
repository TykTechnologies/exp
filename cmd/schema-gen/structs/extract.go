package structs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
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
		info := &StructInfo{
			fileSet: p.fileset,
			Name:    rootName,
			Doc:     TrimSpace(g.Doc),
			Comment: TrimSpace(g.Comment),
		}

		switch obj := g.Type.(type) {
		case *ast.StructType:
			info.structObj = obj
			p.parse(rootName, rootName, info)
		default:
			info.Type = TypeName(g)
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
			Type: getTypeDeclarationsForExpr(field.Type),
			Tag:  tagValue,

			JSONName: jsonName,

			fileSet: structInfo.fileSet,
		}

		structInfo.Fields = append(structInfo.Fields, fieldInfo)
	}
}
