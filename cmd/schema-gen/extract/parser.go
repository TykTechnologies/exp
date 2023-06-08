package extract

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"sort"
	"strings"

	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
	"golang.org/x/exp/slices"
)

// Extract package structs
func Extract(filepath string, ignoreFiles ...string) (*PackageInfo, error) {
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

	p := NewParser(fileSet, pkgs[requiredPkgName])

	return p.GetDeclarations()
}

type objParser struct {
	pkg     *ast.Package
	fileset *token.FileSet

	visited map[string]bool // avoid re-visiting struct type
}

func NewParser(fileset *token.FileSet, pkg *ast.Package) *objParser {
	p := &objParser{
		pkg:     pkg,
		fileset: fileset,
		visited: map[string]bool{},
	}
	return p
}

// GetDeclaration returns a filled out PackageInfo{} and an error if any.
func (p *objParser) GetDeclarations() (*PackageInfo, error) {
	var err error
	result := &PackageInfo{
		Declarations: DeclarationList{},
		Imports:      []string{},
	}

	for _, fileObj := range p.pkg.Files {
		// https://pkg.go.dev/go/ast#File

		// Collect imports
		for _, imported := range fileObj.Imports {
			importLiteral := imported.Path.Value
			if imported.Name != nil {
				alias := getTypeDeclaration(imported.Name)
				fmt.Printf("WARN: package %s is aliased to %s\n", importLiteral, alias)
				importLiteral = alias + " " + importLiteral
			}

			if !strings.Contains(importLiteral, "/internal") {
				result.Imports = append(result.Imports, importLiteral)
			}
		}

		// Collect declarations
		for _, obj := range fileObj.Decls {
			genDecl, ok := obj.(*ast.GenDecl)
			if !ok {
				continue
			}

			info := &DeclarationInfo{
				Doc:     TrimSpace(genDecl.Doc),
				FileDoc: TrimSpace(fileObj.Doc),
				Types:   TypeList{},
			}

			for _, spec := range genDecl.Specs {
				switch obj := spec.(type) {
				case *ast.TypeSpec:
					typeInfo, err := NewTypeSpecInfo(obj)
					if err != nil {
						if !errors.Is(err, ErrUnexported) {
							fmt.Println("Declaration omitted due to error", err)
						}
						continue
					}
					p.parseStruct(typeInfo.Name, typeInfo.Name, typeInfo)

					info.Types.Append(typeInfo)
				}
			}

			sort.Stable(info.Types)

			if info.Valid() {
				result.Declarations.Append(info)
			}
		}
	}

	deduplicate(result.Imports)

	sort.Stable(result.Declarations)
	slices.Sort(result.Imports)

	return result, err
}

// NewTypeSpecInfo allocates a TypeInfo from a TypeSpec node.
// The function returns the struct info and ErrUnexported if
// the struct is not exported. This error is handled outside
// to skip documenting unexported types.
func NewTypeSpecInfo(from *ast.TypeSpec) (*TypeInfo, error) {
	info := &TypeInfo{
		Name:    getTypeDeclaration(from.Name),
		Doc:     TrimSpace(from.Doc),
		Comment: TrimSpace(from.Comment),
	}

	structObj, ok := from.Type.(*ast.StructType)
	if ok {
		info.StructObj = structObj
	} else {
		info.Type = TypeName(from)
	}

	if !ast.IsExported(info.Name) {
		return info, ErrUnexported
	}
	return info, nil
}

func (p *objParser) parseStruct(goPath, name string, structInfo *TypeInfo) {
	if structInfo.StructObj == nil {
		return
	}
	if visited, _ := p.visited[name]; visited {
		return
	}
	p.visited[name] = true

	for _, field := range structInfo.StructObj.Fields.List {
		pos := p.fileset.Position(field.Pos())
		filePos := path.Base(pos.String())

		var goName string
		if len(field.Names) > 0 {
			goName = field.Names[0].Name
		}

		if goName == "" {
			fmt.Printf("[%s] unidentified field in %s\n", filePos, goPath)
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

		fieldPath := goName
		if goPath != "" {
			fieldPath = goPath + "." + goName
		}

		fieldInfo := &FieldInfo{
			Doc:     TrimSpace(field.Doc),
			Comment: TrimSpace(field.Comment),

			Name: goName,
			Path: fieldPath,
			Type: getTypeDeclarationsForExpr(field.Type),
			Tag:  tagValue,

			JSONName: jsonName,
		}

		if fieldInfo.Valid() {
			structInfo.Fields = append(structInfo.Fields, fieldInfo)
		}
	}
	return
}
