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
	// Handle array types
	if strings.HasPrefix(goType, "[]") {
		elementType := strings.TrimPrefix(goType, "[]")
		return map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": getBaseJSONType(elementType),
			},
		}
	}

	// Handle regular types
	return map[string]interface{}{
		"type": getBaseJSONType(goType),
	}
}

// getJSONType converts Go types to JSON Schema types.
// For example, 'int' becomes 'integer', 'float64' becomes 'number'.
func getBaseJSONType(goType string) string {
	switch goType {
	case "int", "int32", "int64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	case "string":
		return "string"
	default:
		return "string" // Default to string for custom types
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
	switch typeName {
	case "int", "int32", "int64", "float32", "float64", "string", "bool":
		return false
	default:
		return true
	}
}

// RequiredFieldsConfig defines which fields are required for each type
type RequiredFieldsConfig struct {
	Fields map[string][]string // map[TypeName][]FieldName
}
