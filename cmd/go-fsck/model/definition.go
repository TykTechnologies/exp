package model

type Definition struct {
	Filename string
	Imports  []string
	Structs  []*Declaration
	Consts   []*Declaration
	Vars     []*Declaration
	Funcs    []*Declaration
}

type DeclarationKind string

const (
	StructKind  DeclarationKind = "struct"
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
