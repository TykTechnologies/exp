# PackageInfo

PackageInfo holds all the declarations for a package scope.

**Field: `imports`** (Imports, `[]string`)

Imports holds a list of imported packages.

**Field: `declarations`** (Declarations, [DeclarationList](#DeclarationList))

Declarations within the package.

# DeclarationList

DeclarationList implements list operations over a `*DeclarationInfo` slice.

Type defined as `[]*DeclarationInfo`, see [DeclarationInfo](DeclarationInfo) definition.

# DeclarationInfo

DeclarationInfo holds the declarations block for an exposed value or type.

**Field: `doc`** (Doc, `string`)

Doc is the declaration doc comment. It usually
occurs just before a *ast.TypeDecl, but may be
applied to multiple ones.

**Field: `file_doc`** (FileDoc, `string`)

FileDoc is the doc comment for a file which
contains the definitions here.

**Field: `types`** (Types, [TypeList](#TypeList))

Types are all the type declarations in the block.

# TypeList

TypeList implements list operations over a *TypeInfo slice.

Type defined as `[]*TypeInfo`, see [TypeInfo](TypeInfo) definition.

# TypeInfo

TypeInfo holds details about a type definition.

**Field: `name`** (Name, `string`)

Name is struct go name.

**Field: `doc`** (Doc, `string`)

Doc is the struct doc.

**Field: `comment`** (Comment, `string`)

Comment is the struct comment.

**Field: `type`** (Type, `string`)

Type is an optional type if the declaration is a type alias or similar.

**Field: `fields`** (Fields, [[]*FieldInfo](#FieldInfo))

Fields holds information of the fields, if this object is a struct.

**Field: `functions`** (Functions, [[]*FuncInfo](#FuncInfo))



# FieldInfo

FieldInfo holds details about a field definition.

**Field: `name`** (Name, `string`)

Name is the name of the field.

**Field: `type`** (Type, `string`)

Type is the literal type of the Go field.

**Field: `path`** (Path, `string`)

Path is the go path of this field starting from root object.

**Field: `doc`** (Doc, `string`)

Doc holds the field doc.

**Field: `comment`** (Comment, `string`)

Comment holds the field comment text.

**Field: `tag`** (Tag, `string`)

Tag is the go tag, unmodified.

**Field: `json_name`** (JSONName, `string`)

JSONName is the corresponding json name of the field.

**Field: `map_key`** (MapKey, `string`)

MapKey is the map key type, if this field is a map.

# FuncInfo

FuncInfo holds details about a function definition.

**Field: `name`** (Name, `string`)



**Field: `doc`** (Doc, `string`)



**Field: `type`** (Type, `string`)



**Field: `path`** (Path, `string`)



**Field: `signature`** (Signature, `string`)



**Field: `source`** (Source, `string`)



