package converter

import (
	"fmt"
	"strings"

	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

// ConvertToJSONSchema converts PackageInfo to JSON Schema with only root type and its dependencies
func ConvertToJSONSchema(pkgInfo *model.PackageInfo, rootType string) (map[string]interface{}, error) {
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
					structSchema := generateStructSchema(typeInfo)
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
func generateStructSchema(typeInfo *model.TypeInfo) map[string]interface{} {
	properties := make(map[string]interface{})
	required := make([]string, 0)

	// Process each field in the struct
	for _, field := range typeInfo.Fields {
		// Initialize field schema with its type
		fieldSchema := map[string]interface{}{
			"type": getJSONType(field.Type),
		}

		// Add documentation as description if available
		if field.Doc != "" {
			fieldSchema["description"] = field.Doc
		}

		// If field type is a custom type (not a built-in type),
		// replace type with a reference to its definition
		if isCustomType(field.Type) {
			fieldSchema["$ref"] = "#/definitions/" + field.Type
			delete(fieldSchema, "type")
		}

		// Add field to properties and mark as required
		properties[field.JSONName] = fieldSchema
		required = append(required, field.JSONName)
	}

	// Create the complete struct schema
	return map[string]interface{}{
		"type":                 "object",
		"properties":           properties,
		"required":             required,
		"additionalProperties": false,
	}
}

// getJSONType converts Go types to JSON Schema types.
// For example, 'int' becomes 'integer', 'float64' becomes 'number'.
func getJSONType(goType string) string {
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
	// For structs, check field types
	for _, field := range typeInfo.Fields {
		// Remove any pointer symbols
		fieldType := strings.TrimPrefix(field.Type, "*")

		// If it's a custom type (not built-in), add it as dependency
		if isCustomType(fieldType) {
			if !dependencies[fieldType] {
				dependencies[fieldType] = true

				// Find the dependent type and collect its dependencies
				for _, decl := range pkgInfo.Declarations {
					for _, depType := range decl.Types {
						if depType.Name == fieldType {
							collectDependencies(depType, pkgInfo, dependencies)
						}
					}
				}
			}
		}
	}
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
