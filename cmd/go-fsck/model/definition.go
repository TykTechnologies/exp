package model

import "strings"

type Definition struct {
	Package  string
	Filename string

	Imports []string
	Types   DeclarationList
	Consts  DeclarationList
	Vars    DeclarationList
	Funcs   DeclarationList
}

type DeclarationKind string

const (
	StructKind  DeclarationKind = "struct"
	ImportKind                  = "import"
	ConstKind                   = "const"
	TypeKind                    = "type"
	FuncKind                    = "func"
	VarKind                     = "var"
	CommentKind                 = "comment"
)

type Declaration struct {
	Kind      DeclarationKind
	Name      string   `json:",omitempty"`
	Names     []string `json:",omitempty"`
	Receiver  string   `json:",omitempty"`
	Signature string   `json:",omitempty"`
	Source    string
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

func (p *DeclarationList) Append(in *Declaration) {
	*p = append(*p, in)
}
