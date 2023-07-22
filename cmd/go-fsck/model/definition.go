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

	Imports map[string][]string
)

type (
	DeclarationKind string

	Declaration struct {
		Kind      DeclarationKind
		File      string
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

func (i Imports) Get(key string) ([]string, bool) {
	val, ok := i[key]
	sort.Strings(val)
	return val, ok
}

type DeclarationList []*Declaration

func (p *DeclarationList) Append(in *Declaration) {
	*p = append(*p, in)
}
