package testusage

import (
	"fmt"
	"log"
	"strings"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

type FuncRef struct {
	FuncDecl          *model.Declaration
	Definition        *model.Definition // Reference to the Definition to access imports and other data
	ReferencedPackage string
	ReferencedSymbol  string
	ImportPath        string
	IsExternal        bool
}

func (f *FuncRef) String() string {
	return fmt.Sprint(f.FuncDecl.Signature)
}

type Report struct {
	Refs []*FuncRef
}

func (r *Report) String() string {
	var b strings.Builder
	for _, v := range r.Refs {
		b.WriteString(v.String() + "\n")
	}
	return b.String()
}

// NewReport analyzes the provided definitions, resolves external dependencies, and returns a list of *FuncRef.
func NewReport(definitions []*model.Definition) ([]*FuncRef, error) {
	var allFuncRefs []*FuncRef

	for _, def := range definitions {
		log.Printf("Processing package: %s\n", def.Package.Package)
		funcRefs, err := getFunctionNames(def)
		if err != nil {
			return nil, err
		}
		allFuncRefs = append(allFuncRefs, funcRefs...)
	}

	// Resolve external dependencies
	resolvedRefs, err := resolveExternalDependencies(allFuncRefs, definitions)
	if err != nil {
		return nil, err
	}

	log.Printf("Total test function references found: %d\n", len(resolvedRefs))
	return resolvedRefs, nil
}

// isTestFunction checks if a function is a test function based on its name or signature.
func isTestFunction(funcDecl *model.Declaration) bool {
	switch {
	case strings.HasPrefix(funcDecl.Name, "Test"):
		return true
	case strings.HasPrefix(funcDecl.Name, "Benchmark"):
		return true
	case strings.HasPrefix(funcDecl.Name, "Example"):
		return true
	}

	// Simplified argument type check
	for _, arg := range funcDecl.Arguments {
		if strings.HasPrefix(arg, "testing.") {
			return true
		}
	}

	return false
}

// getFunctionNames processes a definition and returns a slice of *FuncRef.
func getFunctionNames(def *model.Definition) ([]*FuncRef, error) {
	var funcRefs []*FuncRef

	for _, funcDecl := range def.Funcs {
		if isTestFunction(funcDecl) || strings.HasSuffix(funcDecl.File, "_test.go") {
			for prodPkg, symbols := range funcDecl.References {
				for _, symbol := range symbols {
					isExternal := prodPkg != def.Package.ImportPath
					ref := &FuncRef{
						FuncDecl:          funcDecl,
						Definition:        def, // Store the definition for later use
						ReferencedPackage: prodPkg,
						ReferencedSymbol:  symbol,
						ImportPath:        "", // To be resolved
						IsExternal:        isExternal,
					}
					funcRefs = append(funcRefs, ref)
				}
			}
		}
	}

	return funcRefs, nil
}

// resolveExternalDependencies processes funcRefs and resolves the correct import paths for external symbols.
func resolveExternalDependencies(funcRefs []*FuncRef, definitions []*model.Definition) ([]*FuncRef, error) {
	for _, ref := range funcRefs {
		symbolPackage := ref.ReferencedPackage

		// First, try to resolve within the same definition
		imported, ok := ref.Definition.Imports[ref.FuncDecl.File]
		if !ok {
			log.Printf("No imports found for file: %s in definition: %s", ref.FuncDecl.File, ref.Definition.Package.Package)
		}

		var matchedImport string
		for _, p := range imported {
			if strings.HasPrefix(p, symbolPackage+" ") || strings.HasSuffix(p, symbolPackage+"\"") {
				matchedImport = p
				break
			}
		}

		// If not found, try resolving in other definitions
		if matchedImport == "" {
			for _, def := range definitions {
				if def == ref.Definition {
					continue // Skip the current definition as we've already checked it
				}
				otherImports, ok := def.Imports[ref.FuncDecl.File]
				if ok {
					for _, p := range otherImports {
						if strings.HasPrefix(p, symbolPackage+" ") || strings.HasSuffix(p, symbolPackage+"\"") {
							matchedImport = p
							break
						}
					}
				}
				if matchedImport != "" {
					break
				}
			}
		}

		if matchedImport == "" {
			return nil, fmt.Errorf("no matching import found for symbol package: %s in function: %s", symbolPackage, ref.FuncDecl.Name)
		}

		ref.ImportPath = matchedImport
	}

	return funcRefs, nil
}
