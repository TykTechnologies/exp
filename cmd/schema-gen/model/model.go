package model

import (
	"encoding/json"
	"go/ast"
	"os"
)

// PackageInfo holds all the declarations.
type PackageInfo struct {
	// Imports holds a list of imported packages.
	Imports []string `json:"imports"`

	// Declarations within the package.
	Declarations DeclarationList `json:"declarations"`
}

func LoadPackageInfo(filename string) (*PackageInfo, error) {
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := &PackageInfo{}
	return result, json.Unmarshal(body, result)
}

// DeclarationInfo holds *ast.GenDecl docs and declarations.
type DeclarationInfo struct {
	// Doc is the declaration doc comment. It usually
	// occurs just before a *ast.TypeDecl, but may be
	// applied to multiple ones.
	Doc string `json:"doc,omitempty"`

	// FileDoc is the doc comment for a file which
	// contains the definitions here.
	FileDoc string `json:"file_doc,omitempty"`

	// Types are all the type declarations in the block.
	Types StructList `json:"types,omitempty"`
}

// DeclarationList implements list operations over a *DeclarationInfo slice.
type DeclarationList []*DeclarationInfo

func (x DeclarationList) Len() int           { return len(x) }
func (x DeclarationList) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x DeclarationList) Less(i, j int) bool { return x[i].Doc < x[j].Doc }

func (x *DeclarationList) Append(newInfo *DeclarationInfo) int {
	*x = append(*x, newInfo)
	return len(*x)
}

func (d *DeclarationInfo) Valid() bool {
	return len(d.Types) > 0
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

	// StructObj is the (optionally present) raw ast.StructType value
	StructObj *ast.StructType `json:"-"`
}

// StructList implements list operations over a *StructInfo slice.
type StructList []*StructInfo

func (x StructList) Len() int           { return len(x) }
func (x StructList) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x StructList) Less(i, j int) bool { return x[i].Name < x[j].Name }

func (x *StructList) Append(newInfo *StructInfo) int {
	*x = append(*x, newInfo)
	return len(*x)
}

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
}

func (f FieldInfo) Valid() bool {
	// Not encoded to json.
	if f.JSONName == "-" {
		return false
	}
	// Not exported.
	if !ast.IsExported(f.Name) {
		return false
	}
	// Ignored.
	if f.Name == "_" || f.Name == "" {
		return false
	}
	return true
}
