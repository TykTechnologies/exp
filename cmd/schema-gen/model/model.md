# PackageInfo

PackageInfo holds all the declarations for a package scope.

**Field: `name` (`string`)**
Name is the package name.

**Field: `imports` (`[]string`)**
Imports holds a list of imported packages.

**Field: `declarations` ([DeclarationList](#declarationlist))**
Declarations within the package.

**Field: `functions` ([[]*FuncInfo](#funcinfo))**
Functions within the package, enabled with `--include-functions`.

# DeclarationList

DeclarationList implements list operations over a `*DeclarationInfo` slice.

Type defined as array of `[]*DeclarationInfo` values, see [DeclarationInfo](DeclarationInfo) definition.

# FuncInfo

FuncInfo holds details about a function definition.

**Field: `name` (`string`)**
Name holds the name of the function.

**Field: `doc` (`string`)**
Doc holds the function doc comment.

**Field: `type` (`string`)**
Type holds the receiver if any.

**Field: `path` (`string`)**
Path is the path to the symbol (`Type.FuncName` or `FuncName` if global func).

**Field: `signature` (`string`)**
Signature is an interface compatible signature for the function.

**Field: `source` (`string`)**
Source is a 1-1 source code for the function.

# DeclarationInfo

DeclarationInfo holds the declarations block for an exposed value or type.

**Field: `doc` (`string`)**
Doc is the declaration doc comment. It usually
occurs just before a *ast.TypeDecl, but may be
applied to multiple ones.

**Field: `file_doc` (`string`)**
FileDoc is the doc comment for a file which
contains the definitions here.

**Field: `types` ([TypeList](#typelist))**
Types are all the type declarations in the block.

# TypeList

TypeList implements list operations over a *TypeInfo slice.

Type defined as array of `[]*TypeInfo` values, see [TypeInfo](TypeInfo) definition.

# TypeInfo

TypeInfo holds details about a type definition.

**Field: `name` (`string`)**
Name is struct go name.

**Field: `doc` (`string`)**
Doc is the struct doc.

**Field: `comment` (`string`)**
Comment is the struct comment.

**Field: `type` (`string`)**
Type is an optional type if the declaration is a type alias or similar.

**Field: `fields` ([[]*FieldInfo](#fieldinfo))**
Fields holds information of the fields, if this object is a struct.

**Field: `functions` ([[]*FuncInfo](#funcinfo))**


**Field: `` (`ast.StructType`)**
StructObj is the (optionally present) raw ast.StructType value

# FieldInfo

FieldInfo holds details about a field definition.

**Field: `name` (`string`)**
Name is the name of the field.

**Field: `type` (`string`)**
Type is the literal type of the Go field.

**Field: `path` (`string`)**
Path is the go path of this field starting from root object.

**Field: `doc` (`string`)**
Doc holds the field doc.

**Field: `comment` (`string`)**
Comment holds the field comment text.

**Field: `tag` (`string`)**
Tag is the go tag, unmodified.

**Field: `json_name` (`string`)**
JSONName is the corresponding json name of the field.
It's cleared if it's set to `-` (unexported).

**Field: `map_key` (`string`)**
MapKey is the map key type, if this field is a map.

