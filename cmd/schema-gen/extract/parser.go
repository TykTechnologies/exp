package extract

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"golang.org/x/exp/slices"

	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

// ExtractOptions contains options for extraction
type ExtractOptions struct {
	includeFunctions  bool
	includeTests      bool
	includeUnexported bool
	ignoreFiles       []string
}

func NewExtractOptions(cfg *options) *ExtractOptions {
	return &ExtractOptions{
		includeFunctions:  cfg.includeFunctions,
		includeTests:      cfg.includeTests,
		includeUnexported: cfg.includeUnexported,
		ignoreFiles:       cfg.ignoreFiles,
	}
}

// Extract package structs
func Extract(filepath string, options *ExtractOptions) ([]*PackageInfo, error) {
	var (
		ignoreFiles = options.ignoreFiles
	)

	ignoreList := make(map[string]bool)
	for _, file := range ignoreFiles {
		ignoreList[file] = true
	}

	// filter skips explicitly ignored files, and tests files
	filter := func(fInfo os.FileInfo) bool {
		if ignored := ignoreList[fInfo.Name()]; ignored {
			return false
		}

		if strings.HasSuffix(fInfo.Name(), "_test.go") {
			return options.includeTests
		}

		return true
	}

	fileSet := token.NewFileSet()
	pkgs, err := parser.ParseDir(fileSet, path.Dir(filepath), filter, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	result := make([]*PackageInfo, 0, len(pkgs))
	for pkgName, pkg := range pkgs {
		p := NewParser(fileSet, pkg)
		pkgInfo, err := p.GetDeclarations(options)
		pkgInfo.Name = pkgName
		if err != nil {
			return nil, err
		}
		result = append(result, pkgInfo)
	}
	return result, nil
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

func (p *objParser) getSourceCode(node ast.Node) string {
	var buf strings.Builder
	err := printer.Fprint(&buf, p.fileset, node)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (p *objParser) functionDef(fun *ast.FuncDecl) string {
	var fset = p.fileset
	name := fun.Name.Name
	params := make([]string, 0)
	for _, p := range fun.Type.Params.List {
		var typeNameBuf bytes.Buffer
		err := printer.Fprint(&typeNameBuf, fset, p.Type)
		if err != nil {
			log.Fatalf("failed printing %s", err)
		}
		names := make([]string, 0)
		for _, name := range p.Names {
			names = append(names, name.Name)
		}
		params = append(params, fmt.Sprintf("%s %s", strings.Join(names, ","), typeNameBuf.String()))
	}
	returns := make([]string, 0)
	if fun.Type.Results != nil {
		for _, r := range fun.Type.Results.List {
			var typeNameBuf bytes.Buffer
			err := printer.Fprint(&typeNameBuf, fset, r.Type)
			if err != nil {
				log.Fatalf("failed printing %s", err)
			}

			returns = append(returns, typeNameBuf.String())
		}
	}
	returnString := ""
	if len(returns) == 1 {
		returnString = returns[0]
	} else if len(returns) > 1 {
		returnString = fmt.Sprintf("(%s)", strings.Join(returns, ", "))
	}

	paramsString := strings.Join(params, ", ")
	if returnString != "" {
		return fmt.Sprintf("%s (%s) %v", name, paramsString, returnString)
	}
	return fmt.Sprintf("%s (%s)", name, paramsString)
}

// GetDeclaration returns a filled out PackageInfo{} and an error if any.
func (p *objParser) GetDeclarations(options *ExtractOptions) (*PackageInfo, error) {
	var err error
	result := &PackageInfo{
		Declarations: DeclarationList{},
		Imports:      []string{},
	}

	var funcs []*FuncInfo
	var globalFuncs []*FuncInfo

	for _, fileObj := range p.pkg.Files {
		// https://pkg.go.dev/go/ast#File

		if options.includeFunctions {
			ast.Inspect(fileObj, func(n ast.Node) (res bool) {
				res = true
				if fun, ok := n.(*ast.FuncDecl); ok {
					if !fun.Name.IsExported() && !options.includeUnexported {
						return
					}
					if fun.Recv == nil || len(fun.Recv.List) == 0 {
						name := getTypeDeclaration(fun.Name)
						funcinfo := &FuncInfo{
							Name:      name,
							Doc:       TrimSpace(fun.Doc),
							Path:      name,
							Signature: p.functionDef(fun),
							Source:    p.getSourceCode(n),
						}
						globalFuncs = append(globalFuncs, funcinfo)
						return
					}
					if r, ok := fun.Recv.List[0].Type.(*ast.StarExpr); ok {
						if len(fun.Recv.List[0].Names) == 0 {
							return
						}
						recvType := getTypeDeclaration(fun.Recv.List[0].Names[0]) + " " + getTypeDeclarationsForPointerType(r)
						signature := p.functionDef(fun)
						goPath := r.X.(*ast.Ident).Name
						funcinfo := &FuncInfo{
							Name:      getTypeDeclaration(fun.Name),
							Doc:       TrimSpace(fun.Doc),
							Type:      recvType,
							Path:      goPath,
							Signature: signature,
							Source:    p.getSourceCode(n),
						}
						funcs = append(funcs, funcinfo)
					}

				}

				return
			})
		}

		// Collect imports
		for _, imported := range fileObj.Imports {
			importLiteral := imported.Path.Value
			if imported.Name != nil {
				alias := getTypeDeclaration(imported.Name)
				//fmt.Fprintf(os.Stderr, "WARN: package %s is aliased to %s\n", importLiteral, alias)
				importLiteral = alias + " " + importLiteral
			}

			if !strings.Contains(importLiteral, "/internal") {
				if slices.Contains(result.Imports, importLiteral) {
					continue
				}
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
						isUnexported := errors.Is(err, ErrUnexported)
						if isUnexported && !options.includeUnexported {
							continue
						}
					}

					p.parseStruct(typeInfo.Name, typeInfo.Name, typeInfo, options)

					info.Types.Append(typeInfo)
				}
			}

			sort.Stable(info.Types)

			if info.Valid() {
				result.Declarations.Append(info)
			}
		}
	}

	if options.includeFunctions {
		for _, funcInfo := range funcs {
			for _, decl := range result.Declarations {
				for _, typeDecl := range decl.Types {
					if typeDecl.Name == funcInfo.Path {
						typeDecl.Functions = append(typeDecl.Functions, funcInfo)
					}
				}
			}
		}

		result.Functions = globalFuncs
	}

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

func (p *objParser) parseStruct(goPath, name string, structInfo *TypeInfo, options *ExtractOptions) {
	if structInfo.StructObj == nil {
		return
	}
	if visited, _ := p.visited[name]; visited {
		return
	}
	p.visited[name] = true

	for _, field := range structInfo.StructObj.Fields.List {
		//pos := p.fileset.Position(field.Pos())
		//filePos := path.Base(pos.String())

		var goName string
		if len(field.Names) > 0 {
			goName = field.Names[0].Name

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
		if jsonName == "-" {
			// fields with json `-` don't get encoded
			jsonName = ""
		}

		fieldPath := goName
		if goPath != "" {
			fieldPath = goPath
			if goName != "" {
				fieldPath += "." + goName
			}
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

		isExported := ast.IsExported(fieldInfo.Name)
		if !isExported && !options.includeUnexported {
			continue
		}

		structInfo.Fields = append(structInfo.Fields, fieldInfo)
	}
	return
}
