# PackageInfo

PackageInfo holds all the declarations.

**Imports** (JSON: `imports`)

Imports holds a list of imported packages.

**Declarations** (JSON: `declarations`)

Declarations within the package.

# DeclarationList

DeclarationList implements list operations over a *DeclarationInfo slice.

No exposed fields available.

# DeclarationInfo

DeclarationInfo holds *ast.GenDecl docs and declarations.

**Doc** (JSON: `doc`)

Doc is the declaration doc comment. It usually
occurs just before a *ast.TypeDecl, but may be
applied to multiple ones.

**FileDoc** (JSON: `file_doc`)

FileDoc is the doc comment for a file which
contains the definitions here.

**Types** (JSON: `types`)

Types are all the type declarations in the block.

# FieldInfo

FieldInfo holds details about a field.

**Name** (JSON: `name`)

Name is the name of the field.

**Type** (JSON: `type`)

Type is the literal type of the Go field.

**Path** (JSON: `path`)

Path is the go path of this field starting from root object.

**Doc** (JSON: `doc`)

Doc holds the field doc.

**Comment** (JSON: `comment`)

Comment holds the field comment text.

**Tag** (JSON: `tag`)

Tag is the go tag, unmodified.

**JSONName** (JSON: `json_name`)

JSONName is the corresponding json name of the field.

**MapKey** (JSON: `map_key`)

MapKey is the map key type, if this field is a map.

# TypeInfo

TypeInfo holds ast field information for the docs generator.

**Name** (JSON: `name`)

Name is struct go name.

**Doc** (JSON: `doc`)

Doc is the struct doc.

**Comment** (JSON: `comment`)

Comment is the struct comment.

**Type** (JSON: `type`)

Type is an optional type if the declaration is a type alias or similar.

**Fields** (JSON: `fields`)

Fields holds information of the fields, if this object is a struct.

# TypeList

TypeList implements list operations over a *TypeInfo slice.

No exposed fields available.

