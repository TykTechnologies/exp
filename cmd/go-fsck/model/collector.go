package model

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model/internal"
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

func (v *collector) Clean() {
	for _, def := range v.definition {
		importMap := def.Imports.Map()

		// Change value to print debug output when cleaning imported
		// package references for function declarations.
		debugReferences := true

		if debugReferences {
			fmt.Printf("Imports: %s\n", spew.Sdump(importMap))
		}

		for _, fv := range def.Funcs {
			for k, v := range fv.References {
				if _, ok := importMap[k]; !ok {
					if debugReferences {
						fmt.Printf("Function %s reference doesn't exist in imports: %s: [%v]\n", fv.Name, k, v)
					}
					delete(fv.References, k)
				}
			}
		}
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

func (v *collector) collectImports(filename string, decl *ast.GenDecl, def *Definition) {
	for _, spec := range decl.Specs {
		imported, ok := spec.(*ast.ImportSpec)
		if !ok {
			continue
		}

		importLiteral := imported.Path.Value
		importClean := strings.Trim(importLiteral, `*`)
		if imported.Name != nil {
			alias := imported.Name.Name
			base := path.Base(importClean)
			switch alias {
			case base:
				fmt.Printf("WARN: removing %s alias for %s)\n", alias, importClean)
			case "_":
				// no warning
			default:
				fmt.Printf("WARN: package %s is aliased to %s\n", importLiteral, alias)
				importLiteral = alias + " " + importLiteral
			}
		}

		def.Imports.Add(filename, importLiteral)
	}
}

func collectFuncReferences(funcDecl *ast.FuncDecl) map[string][]string {
	imports := make(map[string][]string)

	// Traverse the function body and look for package identifiers.
	ast.Inspect(funcDecl.Body, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.SelectorExpr:
			// If it's a SelectorExpr, get the leftmost identifier which is the package name.
			if ident, ok := n.X.(*ast.Ident); ok {
				pkgName := ident.Name
				selName := n.Sel.Name
				imports[pkgName] = appendIfNotExists(imports[pkgName], selName)
			}
		case *ast.Ident:
			// If it's an identifier, it might be a package name.
			if obj := n.Obj; obj != nil && obj.Kind == ast.Pkg {
				pkgName := n.Name
				imports[pkgName] = nil // No specific symbol, just mark the package as imported.
			}
		}

		return true
	})

	return imports
}

func (v *collector) Visit(node ast.Node, push bool, stack []ast.Node) bool {
	file, ok := (stack[0]).(*ast.File)
	if !ok {
		return true
	}
	filename := path.Base(v.fset.Position(file.Pos()).Filename)

	packageName := file.Name.Name

	pkg, ok := v.definition[packageName]
	if !ok {
		pkg = &Definition{
			Package: packageName,
		}
		v.definition[packageName] = pkg
	}

	if file.Doc != nil {
		pkg.Doc.Add(filename, v.getSource(file, file.Doc.List))
	}

	switch node := node.(type) {
	case *ast.GenDecl:
		if node.Tok == token.IMPORT {
			v.collectImports(filename, node, pkg)
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
			Names:         names,
			File:          filename,
			SelfContained: internal.IsSelfContainedType(node),
			Source:        v.getSource(file, node),
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
		def := v.collectFuncDeclaration(file, node, filename, stack)
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

func (v *collector) collectFuncDeclaration(file *ast.File, decl *ast.FuncDecl, filename string, stack []ast.Node) *Declaration {
	args, returns := v.functionBindings(decl)

	declaration := &Declaration{
		Kind:       FuncKind,
		File:       filename,
		Name:       decl.Name.Name,
		Arguments:  args,
		Returns:    returns,
		Signature:  v.functionDef(decl),
		References: collectFuncReferences(decl),
		Source:     v.getSource(file, decl),
	}

	if decl.Recv != nil {
		declaration.Receiver = v.symbolType(decl.Recv.List[0].Type)
	}

	return declaration
}

func (p *collector) getSource(file *ast.File, node any) string {
	var buf strings.Builder
	err := internal.PrintSource(internal.CommentedNode(file, node), p.fset, &buf)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (p *collector) symbolType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + p.symbolType(t.X)
	case *ast.ArrayType:
		return "[]" + p.symbolType(t.Elt)
	case *ast.Ellipsis:
		return "..." + p.symbolType(t.Elt)
	case *ast.SelectorExpr:
		return p.symbolType(t.X) + "." + p.symbolType(t.Sel)
	case *ast.MapType:
		k, v := p.symbolType(t.Key), p.symbolType(t.Value)
		return fmt.Sprintf("map[%s]%s", k, v)
	case *ast.InterfaceType:
		return ""
	}
	return fmt.Sprintf("%T", expr)
}

func (p *collector) functionBindings(decl *ast.FuncDecl) (args []string, returns []string) {
	// Traverse arguments
	for _, field := range decl.Type.Params.List {
		argType := p.symbolType(field.Type)
		args = appendIfNotExists(args, argType)
	}

	// Traverse return values
	if decl.Type.Results != nil {
		for _, field := range decl.Type.Results.List {
			returnType := p.symbolType(field.Type)
			returns = appendIfNotExists(returns, returnType)
		}
	}
	return
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
		return fmt.Sprintf("%s (%s) %s", name, paramsString, returnString)
	}
	return fmt.Sprintf("%s (%s)", name, paramsString)
}

func appendIfNotExists(slice []string, element string) []string {
	for _, s := range slice {
		if s == element {
			return slice
		}
	}
	return append(slice, element)
}
