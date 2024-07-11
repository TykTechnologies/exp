package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

// TypeInfo represents information about a type
type TypeInfo struct {
	Name       string   `yaml:"name"`
	Position   string   `yaml:"position"`
	References []string `yaml:"references"`
}

// Scope represents a code scope
type Scope struct {
	Filename string      `yaml:"filename"`
	Types    []*TypeInfo `yaml:"types"`
}

// Scopes represents a collection of Scope
type Scopes struct {
	Scopes []*Scope `yaml:"scopes"`
}

// NewScopes initializes and returns a new Scopes instance
func NewScopes() *Scopes {
	return &Scopes{
		Scopes: make([]*Scope, 0),
	}
}

// addType adds a new type to the scope
func (s *Scopes) addType(filename string, typeInfo *TypeInfo) {
	var scope *Scope
	for _, sc := range s.Scopes {
		if sc.Filename == filename {
			scope = sc
			break
		}
	}
	if scope == nil {
		scope = &Scope{Filename: filename}
		s.Scopes = append(s.Scopes, scope)
	}
	scope.Types = append(scope.Types, typeInfo)
}

// extendScopes extends the scopes to include referenced types
func (s *Scopes) extendScopes() {
	visited := make(map[string]bool)
	for _, scope := range s.Scopes {
		for _, typeInfo := range scope.Types {
			s.extendTypeScope(scope, typeInfo, visited)
		}
	}

	for _, scope := range s.Scopes {
		for _, typeInfo := range scope.Types {
			var refs []string
			skipSymbol := ""
			for _, symbol := range typeInfo.References {
				if skipSymbol == symbol {
					skipSymbol = ""
					continue
				}

				if strings.Contains(symbol, ".") {
					sParts := strings.Split(symbol, ".")
					skipSymbol = sParts[1]
				}

				refs = append(refs, symbol)
			}
			typeInfo.References = refs
		}
	}

}

// extendTypeScope extends the scope for a given type
func (s *Scopes) extendTypeScope(scope *Scope, typeInfo *TypeInfo, visited map[string]bool) {
	if visited[typeInfo.Name] {
		return
	}
	visited[typeInfo.Name] = true

	for _, ref := range typeInfo.References {
		for _, sc := range s.Scopes {
			for _, ti := range sc.Types {
				if ti.Name == ref {
					if sc.Filename != scope.Filename {
						// Move the referenced type to the current scope
						scope.Types = append(scope.Types, ti)
						// Remove the type from its original scope
						sc.Types = removeType(sc.Types, ti.Name)
					}
					s.extendTypeScope(scope, ti, visited)
				}
			}
		}
	}
}

// removeType removes a type from a slice of types
func removeType(types []*TypeInfo, name string) []*TypeInfo {
	for i, t := range types {
		if t.Name == name {
			return append(types[:i], types[i+1:]...)
		}
	}
	return types
}

// Visitor struct to implement the ast.Visitor interface
type Visitor struct {
	fileSet     *token.FileSet
	scopes      *Scopes
	currentFile string
}

var builtInTypes = map[string]struct{}{
	"bool": {}, "byte": {}, "complex64": {}, "complex128": {}, "error": {}, "float32": {}, "float64": {},
	"int": {}, "int8": {}, "int16": {}, "int32": {}, "int64": {}, "rune": {}, "string": {}, "uint": {},
	"uint8": {}, "uint16": {}, "uint32": {}, "uint64": {}, "uintptr": {},
}

// Visit method to traverse the AST nodes
func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.TypeSpec:
		// Process type declaration
		typeInfo := &TypeInfo{
			Name:     n.Name.Name,
			Position: v.fileSet.Position(n.Pos()).String(),
		}
		// Find references within the type declaration
		seen := make(map[string]struct{})
		ast.Inspect(n.Type, func(node ast.Node) bool {
			if ident, ok := node.(*ast.Ident); ok {
				if _, isBuiltIn := builtInTypes[ident.Name]; !isBuiltIn {
					if _, exists := seen[ident.Name]; !exists {
						seen[ident.Name] = struct{}{}
						// Only add fully qualified names or built-in package types
						if ident.Obj == nil {
							typeInfo.References = append(typeInfo.References, ident.Name)
						}
					}
				}
			}

			if selector, ok := node.(*ast.SelectorExpr); ok {
				if pkgIdent, ok := selector.X.(*ast.Ident); ok {
					ref := fmt.Sprintf("%s.%s", pkgIdent.Name, selector.Sel.Name)
					if _, exists := seen[ref]; !exists {
						seen[ref] = struct{}{}
						seen[pkgIdent.Name] = struct{}{}
						typeInfo.References = append(typeInfo.References, ref)
					}
				}
				return true
			}

			return true
		})
		v.scopes.addType(v.currentFile, typeInfo)
	}
	return v
}

func main() {
	// Use the current directory as the source directory
	srcDir := "."

	scopes := NewScopes()

	// Walk through the directory and parse Go files
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			parseFile(path, scopes)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking through source directory: %v\n", err)
	}

	// Extend scopes to include referenced types
	scopes.extendScopes()

	// Print scopes as YAML
	output, err := yaml.Marshal(scopes)
	if err != nil {
		fmt.Printf("Error marshaling scopes to YAML: %v\n", err)
		return
	}
	fmt.Println(string(output))
}

// parseFile parses a Go source file and identifies scopes
func parseFile(filePath string, scopes *Scopes) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file %s: %v\n", filePath, err)
		return
	}

	visitor := &Visitor{
		fileSet:     fset,
		scopes:      scopes,
		currentFile: filePath,
	}
	ast.Walk(visitor, node)
}
