package model

import (
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type (
	Definition struct {
		Package string

		Imports Imports
		Types   DeclarationList
		Consts  DeclarationList
		Vars    DeclarationList
		Funcs   DeclarationList
	}
)

type Imports map[string][]string

type (
	DeclarationKind string

	Declaration struct {
		Kind      DeclarationKind
		File      string
		Imports   []string `json:",omitempty"`
		Name      string   `json:",omitempty"`
		Names     []string `json:",omitempty"`
		Receiver  string   `json:",omitempty"`
		Signature string   `json:",omitempty"`
		Source    string
	}
)

const (
	StructKind  DeclarationKind = "struct"
	ImportKind                  = "import"
	ConstKind                   = "const"
	TypeKind                    = "type"
	FuncKind                    = "func"
	VarKind                     = "var"
	CommentKind                 = "comment"
)

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

func (i *Imports) Add(key, lit string) {
	data := *i
	if data == nil {
		data = make(Imports)
	}
	if set, ok := data[key]; ok {
		if slices.Contains(set, lit) {
			return
		}
		data[key] = append(set, lit)
		return
	}
	data[key] = []string{lit}
	*i = data
}

func (i Imports) Get(key string) []string {
	val, _ := i[key]
	if val != nil {
		sort.Strings(val)
	}
	return val
}

type DeclarationList []*Declaration

func (p *DeclarationList) Append(in *Declaration) {
	*p = append(*p, in)
}
