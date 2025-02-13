package jsonschema

import (
	"strings"
)

// getJSONType converts a Go type into its JSON Schema representation (for fields).
func getJSONType(goType string) *JSONSchema {
	if goType == "[]byte" {
		return &JSONSchema{
			Type:   "string",
			Format: "byte",
		}
	}

	// handle slices
	if strings.HasPrefix(goType, "[]") {
		elementType := strings.TrimPrefix(goType, "[]")
		if isCustomType(elementType) {
			return &JSONSchema{
				Type: "array",
				Items: &JSONSchema{
					Ref: "#/definitions/" + elementType,
				},
			}
		}
		return &JSONSchema{
			Type: "array",
			Items: &JSONSchema{
				Type: getBaseJSONType(elementType),
			},
		}
	}

	// handle maps
	if strings.HasPrefix(goType, "map[") {
		inside := goType[len("map["):]
		parts := strings.SplitN(inside, "]", 2)
		if len(parts) != 2 {
			return &JSONSchema{
				Type:                 "object",
				AdditionalProperties: true,
			}
		}
		keyType := strings.TrimSpace(parts[0])
		valueType := strings.TrimSpace(parts[1])

		if keyType != "string" {
			return &JSONSchema{
				Type:                 "object",
				AdditionalProperties: true,
			}
		}
		if valueType == "interface{}" || valueType == "any" {
			return &JSONSchema{
				Type:                 "object",
				AdditionalProperties: true,
			}
		}
		if !isCustomType(valueType) {
			return &JSONSchema{
				Type: "object",
				AdditionalProperties: &JSONSchema{
					Type: getBaseJSONType(valueType),
				},
			}
		}
		// custom type => $ref
		return &JSONSchema{
			Type: "object",
			AdditionalProperties: &JSONSchema{
				Ref: "#/definitions/" + valueType,
			},
		}
	}

	// fallback for non-array, non-map
	schema := &JSONSchema{
		Type: getBaseJSONType(goType),
	}

	switch goType {
	case "uint8", "byte":
		schema.Minimum = ToPtr(0.0)
		schema.Maximum = ToPtr(255.0)
	case "uint16":
		schema.Minimum = ToPtr(0.0)
		schema.Maximum = ToPtr(65535.0)
	case "uint32":
		schema.Minimum = ToPtr(0.0)
		schema.Maximum = ToPtr(4294967295.0)
	case "uint64", "uint":
		schema.Minimum = ToPtr(0.0)
	case "int8":
		schema.Minimum = ToPtr(-128.0)
		schema.Maximum = ToPtr(127.0)
	case "int16":
		schema.Minimum = ToPtr(-32768.0)
		schema.Maximum = ToPtr(32767.0)
	case "int32", "rune":
		schema.Minimum = ToPtr(-2147483648.0)
		schema.Maximum = ToPtr(2147483647.0)
	case "time.Time":
		schema.Type = "string"
		schema.Format = "date-time"
	case "time.Duration":
		schema.Type = "string"
		schema.Pattern = "^[-+]?([0-9]*(\\.[0-9]*)?[a-z]+)+$"
	case "complex64", "complex128":
		return &JSONSchema{
			Type: "object",
			Properties: map[string]*JSONSchema{
				"real": {Type: "number"},
				"imag": {Type: "number"},
			},
			Required: []string{"real", "imag"},
		}
	}
	return schema
}

// getBaseJSONType maps Go base types to JSON Schema types.
func getBaseJSONType(goType string) string {
	switch goType {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"byte", "rune", "uintptr":
		return "integer"
	case "float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	case "string", "time.Time", "time.Duration":
		return "string"
	case "interface{}", "any":
		return "object"
	case "error":
		return "string"
	default:
		return "string"
	}
}

// getBaseType removes array and pointer markers from a field type.
func getBaseType(fieldType string) string {
	baseType := strings.TrimPrefix(fieldType, "[]")
	baseType = strings.TrimPrefix(baseType, "*")
	// Handle the special case of "[]*SomeType"
	if strings.HasPrefix(fieldType, "[]*") {
		baseType = strings.TrimPrefix(fieldType, "[]*")
	}
	return baseType
}

// isCustomType determines if a type is not one of the built-in or immediate recognized types.
func isCustomType(typeName string) bool {
	if strings.HasPrefix(typeName, "map[") || typeName == "[]byte" {
		return false
	}
	switch typeName {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"complex64", "complex128",
		"string", "bool", "interface{}", "any",
		"byte", "rune",
		"uintptr", "error",
		"time.Time", "time.Duration":
		return false
	}
	return true
}

// buildAliasMap creates a mapping from import alias to import path.
func buildAliasMap(imports []string) map[string]string {
	aliasMap := make(map[string]string)
	for _, imp := range imports {
		parts := strings.Split(imp, " ")
		if len(parts) == 2 {
			alias := parts[0]
			path := strings.Trim(parts[1], "\"")
			aliasMap[alias] = path
		} else {
			// No alias; deduce one from the import path
			impPath := strings.Trim(imp, "\"")
			segs := strings.Split(impPath, "/")
			aliasMap[segs[len(segs)-1]] = impPath
		}
	}
	return aliasMap
}

func ToPtr[T any](v T) *T {
	return &v
}

func parseJSONTag(tagValue string) string {
	parts := strings.SplitN(tagValue, ",", 2)
	if len(parts) == 0 {
		return ""
	}
	name := parts[0]
	if name == "-" {
		return ""
	}
	return name
}


func qualifyTypeName(baseType, pkgAlias string) string {
	if strings.Contains(baseType, ".") {
		return baseType
	}
	if pkgAlias == "" {
		return baseType
	}
	return pkgAlias + "." + baseType
}