package model

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type collector struct {
	fset *token.FileSet

	definition map[string]*Definition
	seen       map[string]*Declaration
}

func NewCollector(fset *token.FileSet) *collector {
	return &collector{
		fset:       fset,
		definition: make(map[string]*Definition),
		seen:       make(map[string]*Declaration),
	}
}

func (v *collector) appendSeen(key string, value *Declaration) {
	if len(value.Names) == 1 {
		value.Name = value.Names[0]
		value.Names = nil
	}
	v.seen[key] = value
}

func (v *collector) isSeen(key string) bool {
	decl, ok := v.seen[key]
	return ok && decl != nil
}

func (v *collector) collectImports(decl *ast.GenDecl, def *Definition) {
	for _, spec := range decl.Specs {
		imported, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}

		importLiteral := imported.Path.Value
		if imported.Name != nil {
			alias := imported.Name.Name
			fmt.Printf("WARN: package %s is aliased to %s\n", importLiteral, alias)
			importLiteral = alias + " " + importLiteral
		}

		if !strings.Contains(importLiteral, "/internal") {
			if slices.Contains(def.Imports, importLiteral) {
				continue
			}
			def.Imports = append(def.Imports, importLiteral)
		}
	}

	sort.Strings(def.Imports)
}

func (v *collector) Visit(node ast.Node, push bool, stack []ast.Node) bool {
	file, ok := (stack[0]).(*ast.File)
	if !ok {
		return true
	}

	packageName := file.Name.Name
	pkg, ok := v.definition[packageName]
	if !ok {
		pkg = &Definition{
			Package: packageName,
		}
		v.definition[packageName] = pkg
	}

	switch node := node.(type) {
	case *ast.GenDecl:
		if node.Tok == token.IMPORT {
			v.collectImports(node, pkg)
			return true
		}

		// If there's a function declaration in the stack,
		// the var/const/struct is internal to a function.
		for _, k := range stack {
			_, ok := k.(*ast.FuncDecl)
			if ok {
				return true
			}
		}

		names := v.Names(node)
		for _, name := range names {
			if v.isSeen(packageName + "." + name) {
				return true
			}
		}

		def := &Declaration{
			Names:  names,
			Source: v.getNodeSource(node),
		}

		for _, name := range names {
			v.appendSeen(packageName+"."+name, def)
		}

		switch node.Tok {
		case token.CONST:
			def.Kind = ConstKind
			pkg.Consts.Append(def)
		case token.VAR:
			def.Kind = VarKind
			pkg.Vars.Append(def)
		case token.TYPE:
			def.Kind = TypeKind
			pkg.Types.Append(def)
		}

	case *ast.FuncDecl:
		def := v.collectFuncDeclaration(node)
		if def != nil {
			key := strings.Trim(packageName+"."+def.Receiver+"."+def.Name, "*.")
			if v.isSeen(key) {
				return true
			}
			defer v.appendSeen(key, def)

			pkg.Funcs.Append(def)
		}

	}

	return true
}

func (v *collector) Names(decl *ast.GenDecl) []string {
	names := make([]string, 0, len(decl.Specs))
	for _, spec := range decl.Specs {
		if val, ok := spec.(*ast.ValueSpec); ok {
			names = append(names, v.identNames(val.Names)...)
			continue
		}

		if val, ok := spec.(*ast.TypeSpec); ok {
			names = append(names, val.Name.Name)
			continue
		}

		v.error("warning getting names: unhandled %T", spec)
	}
	if len(names) == 0 {
		return nil
	}
	return names
}

func (v *collector) error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func (v *collector) identNames(decl []*ast.Ident) []string {
	if len(decl) == 0 {
		return nil
	}

	result := make([]string, 0, len(decl))
	for _, t := range decl {
		result = append(result, t.Name)
	}
	return result
}

func (v *collector) collectFuncDeclaration(decl *ast.FuncDecl) *Declaration {
	declaration := &Declaration{
		Kind:      FuncKind,
		Name:      decl.Name.Name,
		Signature: v.functionDef(decl),
		Source:    v.getNodeSource(decl),
	}

	if decl.Recv != nil {
		var recvType string
		switch t := decl.Recv.List[0].Type.(type) {
		case *ast.StarExpr:
			recvType = "*" + t.X.(*ast.Ident).Name
		case *ast.Ident:
			recvType = t.Name
		}
		declaration.Receiver = recvType
	}

	return declaration
}

func (p *collector) getNodeSource(node ast.Node) string {
	var buf strings.Builder
	err := printer.Fprint(&buf, p.fset, node)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (p *collector) functionDef(fun *ast.FuncDecl) string {
	var fset = p.fset
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
