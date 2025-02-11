package converter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/TykTechnologies/exp/cmd/schema-gen/extract"
	"github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

// ConvertToJSONSchema converts PackageInfo to JSON Schema with only the root type and its (internal and external) dependencies.
func ConvertToJSONSchema(pkgInfo *model.PackageInfo, rootType string, config *RequiredFieldsConfig) (map[string]interface{}, error) {
	schema := map[string]interface{}{
		"$schema":     "http://json-schema.org/draft-07/schema#",
		"definitions": make(map[string]interface{}),
	}
	definitions := schema["definitions"].(map[string]interface{})
	dependencies := make(map[string]bool)

	// Build an alias mapping from the root package's imports.
	aliasMap := buildAliasMap(pkgInfo.Imports)

	// Find the root type and collect its dependencies.
	var rootTypeInfo *model.TypeInfo
	for _, decl := range pkgInfo.Declarations {
		for _, typ := range decl.Types {
			if typ.Name == rootType {
				rootTypeInfo = typ
				collectDependencies(typ, pkgInfo, dependencies)
				break
			}
		}
	}
	if rootTypeInfo == nil {
		return nil, fmt.Errorf("root type %q not found in package", rootType)
	}

	// Process internal types (those without a dot in their name).
	for _, decl := range pkgInfo.Declarations {
		for _, typ := range decl.Types {
			if typ.Name == rootType || dependencies[typ.Name] {
				if !strings.Contains(typ.Name, ".") { // internal type
					if len(typ.EnumValues) > 0 {
						definitions[typ.Name] = generateEnumSchema(typ)
					} else if len(typ.Fields) > 0 {
						definitions[typ.Name] = generateStructSchema(typ, config)
					}
				}
			}
		}
	}

	// Process external dependencies recursively.
	visited := make(map[string]bool)
	for dep := range dependencies {
		if strings.Contains(dep, ".") {
			log.Println("log here i am me", dep)
			if err := processExternalType(dep, aliasMap, definitions, visited); err != nil {
				// Log a warning but continue processing.
				fmt.Fprintf(os.Stderr, "warning: %v\n", err)
			}
		}
	}

	schema["$ref"] = "#/definitions/" + rootType
	return schema, nil
}

// processExternalType loads an external package for a qualified type (e.g. "model.Inner"),
// generates its JSON Schema definition, and then recursively processes its custom fields.
func processExternalType(qualifiedType string, aliasMap map[string]string, definitions map[string]interface{}, visited map[string]bool) error {
	log.Println("here i am quantifies", qualifiedType)
	if visited[qualifiedType] {
		return nil
	}
	visited[qualifiedType] = true

	parts := strings.SplitN(qualifiedType, ".", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid qualified type: %s", qualifiedType)
	}
	pkgAlias := parts[0]
	typeName := parts[1]
	pkgPath, ok := aliasMap[pkgAlias]
	if !ok {
		return fmt.Errorf("alias %q not found in alias map", pkgAlias)
	}
	log.Println("here is me", pkgPath)
	extPkgInfo, err := loadExternalPackage(pkgPath)
	if err != nil {
		return fmt.Errorf("failed to load external package %q: %w", pkgPath, err)
	}
	jsonBytes, err := json.MarshalIndent(extPkgInfo, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal PackageInfo to JSON: %v", err)
	}

	fmt.Println(string(jsonBytes))

	// Build an alias map for the external package.
	extAliasMap := buildAliasMap(extPkgInfo.Imports)
	// Also add an entry for the external package’s own name.
	extAliasMap[extPkgInfo.Name] = pkgPath

	// Look for the type in the external package.
	var extType *model.TypeInfo
	for _, decl := range extPkgInfo.Declarations {
		for _, typ := range decl.Types {
			if typ.Name == typeName {
				extType = typ
				break
			}
		}
	}
	if extType == nil {
		return fmt.Errorf("type %q not found in external package %q", typeName, pkgPath)
	}

	// Generate the JSON Schema for the external type.
	var extSchema map[string]interface{}
	if len(extType.EnumValues) > 0 {
		extSchema = generateEnumSchema(extType)
	} else if len(extType.Fields) > 0 {
		// For external types, we use an empty required config (or you could supply defaults).
		extSchema = generateStructSchema(extType, &RequiredFieldsConfig{Fields: map[string][]string{}})
	}
	if extSchema != nil {
		// Save the definition using the fully qualified name.
		definitions[qualifiedType] = extSchema
	}

	// Recursively process each custom field in the external type.
	for _, field := range extType.Fields {
		baseType := getBaseType(field.Type)
		if isCustomType(baseType) {
			var depQualified string
			// If the field's type is already qualified, use it; otherwise, qualify it with the external package alias.
			if strings.Contains(baseType, ".") {
				depQualified = baseType
			} else {
				depQualified = pkgAlias + "." + baseType
			}
			if err := processExternalType(depQualified, extAliasMap, definitions, visited); err != nil {
				return err
			}
		}
	}

	return nil
}

// loadExternalPackage uses golang.org/x/tools/go/packages to load a package from its import path
// and then runs your extraction process on it.
func loadExternalPackage(pkgPath string) (*model.PackageInfo, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports,
	}
	pkgs, err := packages.Load(cfg, pkgPath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no package found for %q", pkgPath)
	}
	if len(pkgs[0].GoFiles) == 0 {
		return nil, fmt.Errorf("external package %q has no Go files", pkgPath)
	}
	log.Println("log.ppepe", pkgs[0].GoFiles)
	// Use the directory of the first Go file in the package.
	pkgDir := filepath.Dir(pkgs[0].GoFiles[0])
	log.Println(pkgDir)
	pkgInfos, err := extract.Extract(pkgDir+"/", &extract.ExtractOptions{})
	if err != nil {
		return nil, err
	}
	if len(pkgInfos) == 0 {
		return nil, fmt.Errorf("no package info extracted for %q", pkgPath)
	}
	return pkgInfos[0], nil
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
			// No alias; deduce one from the import path.
			impPath := strings.Trim(imp, "\"")
			segs := strings.Split(impPath, "/")
			aliasMap[segs[len(segs)-1]] = impPath
		}
	}
	return aliasMap
}

// generateEnumSchema creates a JSON Schema definition for an enum type.
func generateEnumSchema(typeInfo *model.TypeInfo) map[string]interface{} {
	enumValues := make([]interface{}, 0)
	for _, enum := range typeInfo.EnumValues {
		enumValues = append(enumValues, enum.Value)
	}
	jsonType := "string"
	if typeInfo.Type == "int" {
		jsonType = "integer"
	}
	return map[string]interface{}{
		"type": jsonType,
		"enum": enumValues,
	}
}

// generateStructSchema creates a JSON Schema definition for a struct type.
func generateStructSchema(typeInfo *model.TypeInfo, config *RequiredFieldsConfig) map[string]interface{} {
	properties := make(map[string]interface{})
	required := make([]string, 0)
	requiredFields := config.Fields[typeInfo.Name]
	requiredMap := make(map[string]bool)
	for _, field := range requiredFields {
		requiredMap[field] = true
	}
	for _, field := range typeInfo.Fields {
		isArray := strings.HasPrefix(field.Type, "[]")
		// Use the field type as-is (this might be a qualified type like "model.Inner").
		baseType := strings.TrimPrefix(field.Type, "[]")
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
					"$ref": "#/definitions/" + baseType,
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

// getJSONType converts a Go type into its JSON Schema representation.
func getJSONType(goType string) map[string]interface{} {
	if goType == "[]byte" {
		return map[string]interface{}{
			"type":   "string",
			"format": "byte",
		}
	}
	if strings.HasPrefix(goType, "[]") {
		elementType := strings.TrimPrefix(goType, "[]")
		return map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": getBaseJSONType(elementType),
			},
		}
	}
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
	// Add constraints for some numeric types.
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

// collectDependencies recursively collects type dependencies from a given type.
func collectDependencies(typeInfo *model.TypeInfo, pkgInfo *model.PackageInfo, dependencies map[string]bool) {
	for _, field := range typeInfo.Fields {
		baseType := getBaseType(field.Type)
		if isCustomType(baseType) {
			if !dependencies[baseType] {
				dependencies[baseType] = true
				// For internal types, collect dependencies from the same package.
				if !strings.Contains(baseType, ".") {
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
}

// getBaseType removes array and pointer markers from a field type.
func getBaseType(fieldType string) string {
	baseType := strings.TrimPrefix(fieldType, "[]")
	baseType = strings.TrimPrefix(baseType, "*")
	if strings.HasPrefix(fieldType, "[]*") {
		baseType = strings.TrimPrefix(fieldType, "[]*")
	}
	return baseType
}

// isCustomType determines if a type is not one of the built-in types.
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
	default:
		return true
	}
}

// RequiredFieldsConfig defines which fields are required for each type.
type RequiredFieldsConfig struct {
	Fields map[string][]string // map[TypeName][]FieldName
}

func test() {
}
