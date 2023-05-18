package structs

import (
	"go/ast"
	"go/token"
)

// StructList is a list of information for exported struct type info,
// starting from the root struct declaration(XTykGateway).
type StructList []*StructInfo

func (x StructList) Len() int           { return len(x) }
func (x StructList) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x StructList) Less(i, j int) bool { return x[i].Name < x[j].Name }

func (x *StructList) append(newInfo *StructInfo) int {
	*x = append(*x, newInfo)
	return len(*x)
}

// StructInfo holds ast field information for the docs generator.
type StructInfo struct {
	// Name is struct go name.
	Name string `json:"name"`

	// Doc is the struct doc.
	Doc string `json:"doc,omitempty"`

	// Comment is the struct comment.
	Comment string `json:"comment,omitempty"`

	// Type is an optional type if the declaration is a type alias or similar.
	Type string `json:"type,omitempty"`

	// Fields holds information of the fields, if this object is a struct.
	Fields []*FieldInfo `json:"fields,omitempty"`

	// fileSet holds a token.FileSet, used to resolve symbols to file:line
	fileSet *token.FileSet

	// structObj is the raw ast.StructType value, private.
	structObj *ast.StructType
}

func (*StructInfo) Anchor() {}

// FieldInfo holds details about a field.
type FieldInfo struct {
	// Name is the name of the field.
	Name string `json:"name"`

	// Type is the literal type of the Go field.
	Type string `json:"type"`

	// Path is the go path of this field starting from root object.
	Path string `json:"path"`

	// Doc holds the field doc.
	Doc string `json:"doc,omitempty"`

	// Comment holds the field comment text.
	Comment string `json:"comment,omitempty"`

	// Tag is the go tag, unmodified.
	Tag string `json:"tag"`

	// JSONName is the corresponding json name of the field.
	JSONName string `json:"json_name"`

	// MapKey is the map key type, if this field is a map.
	MapKey string `json:"map_key,omitempty"`

	// IsArray reports if this field is an array.
	IsArray bool `json:"is_array"`

	// fileSet holds a token.FileSet, used to resolve symbols to file:line.
	fileSet *token.FileSet
}

func (*FieldInfo) Anchor() {}
