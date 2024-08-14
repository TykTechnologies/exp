package model

import (
	"go/ast"
	"sort"
	"strings"
)

type (
	Package struct {
		// Package is the name of the package.
		Package string
		// ImportPath contains the import path (github...).
		ImportPath string
		// Path is sanitized to contain the relative location (folder).
		Path string
		// TestPackage is true if this is a test package.
		TestPackage bool
	}

	Definition struct {
		Package

		Doc     StringSet
		Imports StringSet
		Types   DeclarationList
		Consts  DeclarationList
		Vars    DeclarationList
		Funcs   DeclarationList
	}
)

type Declaration struct {
	Kind DeclarationKind
	File string

	SelfContained bool

	Imports []string `json:",omitempty"`

	References map[string][]string `json:",omitempty"`

	Name     string   `json:",omitempty"`
	Names    []string `json:",omitempty"`
	Receiver string   `json:",omitempty"`

	Arguments []string `json:",omitempty"`
	Returns   []string `json:",omitempty"`

	Signature string `json:",omitempty"`
	Source    string
}

func (p Package) Name() string {
	return p.Package
}

func (d DeclarationKind) String() string {
	return string(d)
}

func (d *Definition) getImports(decl *Declaration) []string {
	return d.Imports.Get(decl.File)
}

func (d *Definition) Order() []*Declaration {
	count := len(d.Types) + len(d.Funcs) + len(d.Vars) + len(d.Consts)
	result := make([]*Declaration, 0, count)

	result = append(result, d.Types...)
	result = append(result, d.Funcs...)
	result = append(result, d.Vars...)
	result = append(result, d.Consts...)
	return result
}

func (d *Definition) Sort() {
	d.Types.Sort()
	d.Vars.Sort()
	d.Consts.Sort()
	d.Funcs.Sort()
}

func (d *Definition) ClearSource() {
	d.Types.ClearSource()
	d.Vars.ClearSource()
	d.Consts.ClearSource()
	d.Funcs.ClearSource()
}

func (d *Definition) Fill() {
	for _, decl := range d.Order() {
		decl.Imports = d.getImports(decl)
	}
}

func (d *Declaration) Keys() []string {
	trimPath := "*."
	if d.Name != "" {
		return []string{
			strings.Trim(d.Receiver+"."+d.Name, trimPath),
		}
	}
	if len(d.Names) != 0 {
		result := make([]string, len(d.Names))
		for k, v := range d.Names {
			result[k] = strings.Trim(d.Receiver+"."+v, trimPath)
		}
	}
	return nil
}

type DeclarationList []*Declaration

func (p *DeclarationList) Append(in ...*Declaration) {
	*p = append(*p, in...)
}

func (p DeclarationList) FindKind(kind DeclarationKind) (result []*Declaration) {
	for _, decl := range p {
		if decl.Kind == kind {
			result = append(result, decl)
		}
	}
	return
}

func (p *DeclarationList) ClearSource() {
	for _, decl := range *p {
		decl.Source = ""
	}
}

func (p *DeclarationList) Sort() {
	sort.Slice(*p, func(i, j int) bool {
		a, b := (*p)[i], (*p)[j]
		if a.Kind != b.Kind {
			indexOf := map[DeclarationKind]int{
				CommentKind: 0,
				ImportKind:  1,
				ConstKind:   2,
				StructKind:  3,
				TypeKind:    4,
				VarKind:     5,
				FuncKind:    6,
			}
			return indexOf[a.Kind] < indexOf[b.Kind]
		}
		ae, be := ast.IsExported(a.Name), ast.IsExported(b.Name)
		if ae != be {
			return ae
		}

		if a.Receiver != b.Receiver {
			if a.Receiver == "" {
				return true
			}
			return a.Receiver < b.Receiver
		}

		if a.Signature != b.Signature {
			return a.Signature < b.Signature
		}

		return a.Name < b.Name
	})
}
