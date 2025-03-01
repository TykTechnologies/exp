package model

import (
	"encoding/json"
	"go/ast"
	"io"
	"os"
	"sort"
	"strings"
)

// PackageInfo holds all the declarations for a package scope.
type PackageInfo struct {
	// Name is the package name.
	Name string `json:"name"`

	// Imports holds a list of imported packages.
	Imports []string `json:"imports"`

	// Declarations within the package.
	Declarations DeclarationList `json:"declarations"`

	// Functions within the package, enabled with `--include-functions`.
	Functions []*FuncInfo `json:"functions,omitempty"`
}

// Load reads and decodes a json file to produce a `*PackageInfo`.
func Load(filename string) ([]*PackageInfo, error) {
	var body []byte
	var err error

	switch filename {
	case "-":
		body, err = io.ReadAll(os.Stdin)
	default:
		body, err = os.ReadFile(filename)
	}

	if err != nil {
		return nil, err
	}

	result := []*PackageInfo{}
	err = json.Unmarshal(body, &result)
	return result, err
}

// DeclarationInfo holds the declarations block for an exposed value or type.
type DeclarationInfo struct {
	// Doc is the declaration doc comment. It usually
	// occurs just before a *ast.TypeDecl, but may be
	// applied to multiple ones.
	Doc string `json:"doc,omitempty"`

	// FileDoc is the doc comment for a file which
	// contains the definitions here.
	FileDoc string `json:"file_doc,omitempty"`

	// Types are all the type declarations in the block.
	Types TypeList `json:"types,omitempty"`
}

// DeclarationList implements list operations over a `*DeclarationInfo` slice.
type DeclarationList []*DeclarationInfo

// TypeMap returns a map of type references, the key being the type name.
func (x DeclarationList) TypeMap() map[string]*TypeInfo {
	result := make(map[string]*TypeInfo)

	for _, decl := range x {
		for _, t := range decl.Types {
			if t.Doc == "" && strings.HasPrefix(decl.Doc, t.Name) {
				t.Doc = decl.Doc
			}
			result[t.Name] = t
		}
	}

	return result
}

// Find returns a TypeList containing TypeInfo objects from the DeclarationList in the specified order.
func (x DeclarationList) Find(order []string) TypeList {
	typeInfoMap := x.TypeMap()
	result := make(TypeList, 0, len(typeInfoMap))

	if len(order) == 0 {
		order = make([]string, 0, len(typeInfoMap))
		for key := range typeInfoMap {
			order = append(order, key)
		}
		sort.Strings(order)
	}

	for _, typeName := range order {
		if typeInfo, ok := typeInfoMap[typeName]; ok {
			result = append(result, typeInfo)
		}
	}

	return result
}

func (x DeclarationList) Len() int           { return len(x) }
func (x DeclarationList) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x DeclarationList) Less(i, j int) bool { return x[i].Doc < x[j].Doc }

func (x *DeclarationList) Append(newInfo *DeclarationInfo) int {
	*x = append(*x, newInfo)
	return len(*x)
}

func (x *DeclarationInfo) Valid() bool {
	return len(x.Types) > 0
}

// EnumInfo holds details about an enum definition.
type EnumInfo struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Doc   string      `json:"doc,omitempty"`
}

// TypeInfo holds details about a type definition.
type TypeInfo struct {
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

	Functions []*FuncInfo `json:"functions,omitempty"`

	// StructObj is the (optionally present) raw ast.StructType value
	StructObj *ast.StructType `json:"-"`

	// Enums hold information for an enum value.
	Enums []*EnumInfo `json:"enums,omitempty"`
}

// TypeRef trims array and pointer info from a type reference.
func (f *TypeInfo) TypeRef() string {
	return strings.TrimLeft(f.Type, "[]*")
}

// TypeList implements list operations over a *TypeInfo slice.
type TypeList []*TypeInfo

// TypeMap returns a map of type references, the key being the type name.
func (x TypeList) TypeMap() map[string]*TypeInfo {
	result := make(map[string]*TypeInfo)

	for _, t := range x {
		result[t.Name] = t
	}

	return result
}

func (x TypeList) Len() int           { return len(x) }
func (x TypeList) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x TypeList) Less(i, j int) bool { return x[i].Name < x[j].Name }

func (x *TypeList) Append(newInfo *TypeInfo) int {
	*x = append(*x, newInfo)
	return len(*x)
}

// FieldInfo holds details about a field definition.
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
	// It's cleared if it's set to `-` (unexported).
	JSONName string `json:"json_name"`

	// MapKey is the map key type, if this field is a map.
	MapKey string `json:"map_key,omitempty"`
}

func (f *FieldInfo) TypeRef() string {
	return strings.TrimLeft(f.Type, "[]*")
}

// FuncInfo holds details about a function definition.
type FuncInfo struct {
	// Name holds the name of the function.
	Name string `json:"name"`

	// Doc holds the function doc comment.
	Doc string `json:"doc,omitempty"`

	// Type holds the receiver if any.
	Type string `json:"type,omitempty"`

	// Path is the path to the symbol (`Type.FuncName` or `FuncName` if global func).
	Path string `json:"path"`

	// Signature is an interface compatible signature for the function.
	Signature string `json:"signature"`

	// Source is a 1-1 source code for the function.
	Source string `json:"source"`
}
