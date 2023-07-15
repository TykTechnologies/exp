package model

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
)

type declarationCollector struct {
	fset       *token.FileSet
	definition *Definition
}

func (v *declarationCollector) error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func (v *declarationCollector) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.GenDecl:
		switch node.Tok {
		case token.TYPE:
			v.collectTypeDeclarations(node)
		case token.CONST:
			v.collectConstDeclarations(node)
		case token.VAR:
			v.collectVarDeclarations(node)
		}
	case *ast.FuncDecl:
		v.collectFuncDeclaration(node)
	}

	return v
}

func (v *declarationCollector) declNames(decl *ast.GenDecl) []string {
	names := make([]string, 0, len(decl.Specs))
	for _, spec := range decl.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			v.error("skipped %T, expected ConstSpec", spec)
			continue
		}
		for _, name := range valueSpec.Names {
			names = append(names, name.Name)
		}
	}
	return names
}

func (v *declarationCollector) collectTypeDeclarations(decl *ast.GenDecl) {
	declaration := &Declaration{
		Kind:   TypeKind,
		Names:  v.declNames(decl),
		Source: v.getNodeSource(decl),
	}

	v.definition.Structs = append(v.definition.Structs, declaration)
}

func (v *declarationCollector) collectConstDeclarations(decl *ast.GenDecl) {
	declaration := &Declaration{
		Kind:   ConstKind,
		Names:  v.declNames(decl),
		Source: v.getNodeSource(decl),
	}

	v.definition.Consts = append(v.definition.Consts, declaration)
}

func (v *declarationCollector) collectVarDeclarations(decl *ast.GenDecl) {
	declaration := &Declaration{
		Kind:   VarKind,
		Names:  v.declNames(decl),
		Source: v.getNodeSource(decl),
	}

	v.definition.Vars = append(v.definition.Vars, declaration)
}

func (v *declarationCollector) collectFuncDeclaration(decl *ast.FuncDecl) {
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

	v.definition.Funcs = append(v.definition.Funcs, declaration)
}

func (p *declarationCollector) getNodeSource(node ast.Node) string {
	var buf strings.Builder
	err := printer.Fprint(&buf, p.fset, node)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (p *declarationCollector) functionDef(fun *ast.FuncDecl) string {
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
