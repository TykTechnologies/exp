package converter

import (
	"fmt"
	"strings"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

// ConvertToJSONSchema converts PackageInfo to JSON Schema with only root type and its dependencies
func ConvertToJSONSchema(pkgInfo *model.PackageInfo, rootType string, config *RequiredFieldsConfig) (map[string]interface{}, error) {
	schema := map[string]interface{}{
		"$schema":     "http://json-schema.org/draft-07/schema#",
		"definitions": make(map[string]interface{}),
	}

	definitions := schema["definitions"].(map[string]interface{})
	dependencies := make(map[string]bool)

	// Find root type and collect its dependencies
	var rootTypeInfo *model.TypeInfo
	for _, decl := range pkgInfo.Declarations {
		for _, typeInfo := range decl.Types {
			if typeInfo.Name == rootType {
				rootTypeInfo = typeInfo
				collectDependencies(typeInfo, pkgInfo, dependencies)
				break
			}
		}
	}

	if rootTypeInfo == nil {
		return nil, fmt.Errorf("root type %q not found in package", rootType)
	}

	// Process only root type and its dependencies
	for _, decl := range pkgInfo.Declarations {
		for _, typeInfo := range decl.Types {
			if typeInfo.Name == rootType || dependencies[typeInfo.Name] {
				switch {
				case len(typeInfo.EnumValues) > 0:
					enumSchema := generateEnumSchema(typeInfo)
					definitions[typeInfo.Name] = enumSchema

				case len(typeInfo.Fields) > 0:
					structSchema := generateStructSchema(typeInfo, config)
					definitions[typeInfo.Name] = structSchema
				}
			}
		}
	}

	schema["$ref"] = "#/definitions/" + rootType

	return schema, nil
}

// generateEnumSchema creates a JSON Schema definition for an enum type.
// It handles both string and integer enums.
func generateEnumSchema(typeInfo *model.TypeInfo) map[string]interface{} {
	enumValues := make([]interface{}, 0)

	for _, enum := range typeInfo.EnumValues {
		enumValues = append(enumValues, enum.Value)
	}

	// Convert Go types to proper JSON Schema types
	jsonType := "string"
	if typeInfo.Type == "int" {
		jsonType = "integer"
	}

	schema := map[string]interface{}{
		"type": jsonType,
		"enum": enumValues,
	}

	return schema
}

// generateStructSchema creates a JSON Schema definition for a struct type.
// It processes all fields and creates property definitions.
func generateStructSchema(typeInfo *model.TypeInfo, config *RequiredFieldsConfig) map[string]interface{} {
	properties := make(map[string]interface{})
	required := make([]string, 0)

	requiredFields := config.Fields[typeInfo.Name]
	requiredMap := make(map[string]bool)
	for _, field := range requiredFields {
		requiredMap[field] = true
	}

	for _, field := range typeInfo.Fields {
		// Get base type without array prefix
		baseType := strings.TrimPrefix(field.Type, "[]")
		isArray := strings.HasPrefix(field.Type, "[]")

		var fieldSchema map[string]interface{}

		if isCustomType(baseType) {
			if isArray {
				fieldSchema = map[string]interface{}{
					"type": "array",
					"items": map[string]interface{}{
						"$ref": "#/definitions/" + baseType,
					},
				}
			} else {
				fieldSchema = map[string]interface{}{
					"$ref": "#/definitions/" + field.Type,
				}
			}
		} else {
			fieldSchema = getJSONType(field.Type)
		}

		if field.Doc != "" {
			fieldSchema["description"] = field.Doc
		}

		properties[field.JSONName] = fieldSchema

		if requiredMap[field.Name] {
			required = append(required, field.JSONName)
		}
	}

	schema := map[string]interface{}{
		"type":                 "object",
		"properties":           properties,
		"additionalProperties": false,
	}

	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

func getJSONType(goType string) map[string]interface{} {
	if goType == "[]byte" {
		return map[string]interface{}{
			"type":   "string",
			"format": "byte",
		}
	}

	// Handle arrays
	if strings.HasPrefix(goType, "[]") {
		elementType := strings.TrimPrefix(goType, "[]")
		return map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": getBaseJSONType(elementType),
			},
		}
	}

	// Handle maps
	if strings.HasPrefix(goType, "map[") {
		valueType := strings.Split(strings.TrimPrefix(goType, "map["), "]")[1]

		if valueType == "interface{}" || valueType == "any" {
			return map[string]interface{}{
				"type":                 "object",
				"additionalProperties": true,
			}
		}

		return map[string]interface{}{
			"type": "object",
			"additionalProperties": map[string]interface{}{
				"type": getBaseJSONType(valueType),
			},
		}
	}

	schema := map[string]interface{}{
		"type": getBaseJSONType(goType),
	}

	// Add constraints for numeric types
	switch goType {
	case "uint8", "byte":
		schema["minimum"] = 0
		schema["maximum"] = 255
	case "uint16":
		schema["minimum"] = 0
		schema["maximum"] = 65535
	case "uint32":
		schema["minimum"] = 0
		schema["maximum"] = 4294967295
	case "uint64", "uint":
		schema["minimum"] = 0
	case "int8":
		schema["minimum"] = -128
		schema["maximum"] = 127
	case "int16":
		schema["minimum"] = -32768
		schema["maximum"] = 32767
	case "int32", "rune":
		schema["minimum"] = -2147483648
		schema["maximum"] = 2147483647
	case "time.Time":
		schema["type"] = "string"
		schema["format"] = "date-time"
	case "time.Duration":
		schema["type"] = "string"
		schema["pattern"] = "^[-+]?([0-9]*(\\.[0-9]*)?[a-z]+)+$"
	case "complex64", "complex128":
		// Represent complex numbers as an object with real and imaginary parts
		return map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"real": map[string]interface{}{
					"type": "number",
				},
				"imag": map[string]interface{}{
					"type": "number",
				},
			},
			"required": []string{"real", "imag"},
		}
	}

	return schema
}

// getJSONType converts Go types to JSON Schema types.
// For example, 'int' becomes 'integer', 'float64' becomes 'number'.
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

// collectDependencies recursively collects all type dependencies
func collectDependencies(typeInfo *model.TypeInfo, pkgInfo *model.PackageInfo, dependencies map[string]bool) {
	for _, field := range typeInfo.Fields {
		baseType := getBaseType(field.Type)

		if isCustomType(baseType) {
			if !dependencies[baseType] {
				dependencies[baseType] = true

				for _, decl := range pkgInfo.Declarations {
					for _, depType := range decl.Types {
						if depType.Name == baseType {
							collectDependencies(depType, pkgInfo, dependencies)
						}
					}
				}
			}
		}
	}
}

// getBaseType Removes array and pointer from a filed type
func getBaseType(fieldType string) string {
	baseType := strings.TrimPrefix(fieldType, "[]")
	baseType = strings.TrimPrefix(baseType, "*")
	// Handle case where pointer comes before array
	if strings.HasPrefix(fieldType, "[]*") {
		baseType = strings.TrimPrefix(fieldType, "[]*")
	}
	return baseType
}

// isCustomType determines if a type is a built-in type or a custom type.
// Returns true for custom types that need to be referenced in definitions.
func isCustomType(typeName string) bool {
	// Check for maps
	if strings.HasPrefix(typeName, "map[") {
		return false
	}

	if typeName == "[]byte" {
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
	default:
		return true
	}
}

// RequiredFieldsConfig defines which fields are required for each type
type RequiredFieldsConfig struct {
	Fields map[string][]string // map[TypeName][]FieldName
}
