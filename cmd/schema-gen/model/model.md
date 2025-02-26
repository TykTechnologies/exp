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

Type defined as array of `DeclarationInfo` values, see [DeclarationInfo](#declarationinfo) definition.

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

Type defined as array of `TypeInfo` values, see [TypeInfo](#typeinfo) definition.

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

**Field: `enums` ([[]*EnumInfo](#enuminfo))**
Enums hold information for an enum value.

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

# EnumInfo

EnumInfo holds details about an enum definition.

**Field: `name` (`string`)**


**Field: `value` (``)**


**Field: `doc` (`string`)**


# ExtractOptions

ExtractOptions contains options for extraction

**Field: `IncludeFunctions` (`boolean`)**


**Field: `IncludeTests` (`boolean`)**


**Field: `IncludeUnexported` (`boolean`)**


**Field: `IgnoreFiles` (`[]string`)**


**Field: `IncludeInternal` (`boolean`)**


# JSONSchema

JSONSchema represents a JSON Schema document according to the draft-07 specification.
It includes standard fields used to define types, formats, validations.

**Field: `$schema` (`string`)**
Schema specifies the JSON Schema version URL.
Example: "http://json-schema.org/draft-07/schema#"

**Field: `$ref` (`string`)**
Ref is used to reference another schema definition.
Example: "#/definitions/SomeType"

**Field: `definitions` (`map[string]JSONSchema`)**
Definitions contains subSchema definitions that can be referenced by $ref.

**Field: `type` (`string`)**
Type indicates the JSON type of the instance (e.g., "object", "array", "string").

**Field: `format` (`string`)**
Format provides additional semantic validation for the instance.
Common formats include "date-time", "email", etc.

**Field: `pattern` (`string`)**
Pattern defines a regular expression that a string value must match

**Field: `properties` (`map[string]JSONSchema`)**
Properties defines the fields of an object and their corresponding schemas

**Field: `items` ([JSONSchema](#jsonschema))**
Items defines the schema for array elements

**Field: `enum` (`[]any`)**
Enum restricts a value to a fixed set of values

**Field: `required` (`[]string`)**
Required lists the properties that must be present in an object

**Field: `description` (`string`)**
Description provides a human-readable explanation of the schema.

**Field: `minimum` (`float64`)**
Minimum specifies the minimum numeric value allowed.

**Field: `maximum` (`float64`)**
Maximum specifies the maximum numeric value allowed.

**Field: `exclusiveMinimum` (`boolean`)**
ExclusiveMinimum, if true, requires the instance to be greater than (not equal to) Minimum.

**Field: `exclusiveMaximum` (`boolean`)**
ExclusiveMaximum, if true, requires the instance to be less than (not equal to) Maximum.

**Field: `multipleOf` (`float64`)**
MultipleOf indicates that the numeric instance must be a multiple of this value.

**Field: `additionalProperties` (`any`)**
AdditionalProperties controls whether an object can have properties beyond those defined
Can be a boolean or a schema that additional properties must conform to

