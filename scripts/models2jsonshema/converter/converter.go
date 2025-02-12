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

// ParseAndConvertStruct parses the given repo directory for Go structs and
// converts the specified rootType to JSON Schema, writing the result to "schema.json".
func ParseAndConvertStruct(repoDir, rootType, outFile string) error {
	// Default the output file to "schema.json" if none is provided
	if outFile == "" {
		outFile = "schema.json"
	}

	absDir, err := filepath.Abs(repoDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for %q: %w", repoDir, err)
	}

	// Extract type information from the user's code
	pkgInfos, err := extract.Extract(absDir+"/", &extract.ExtractOptions{})
	if err != nil {
		return fmt.Errorf("failed to extract types from %q: %w", absDir, err)
	}
	if len(pkgInfos) == 0 {
		return fmt.Errorf("no package info extracted from %q", absDir)
	}

	// Convert to JSON Schema
	schema, err := ConvertToJSONSchema(pkgInfos[0], absDir, rootType, NewDefaultConfig())
	if err != nil {
		return fmt.Errorf("failed to convert to JSON Schema: %w", err)
	}

	// Marshal the JSON Schema and write to file
	jsonBytes, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal schema to JSON: %w", err)
	}
	if err := os.WriteFile(outFile, jsonBytes, 0o644); err != nil {
		return fmt.Errorf("failed to write %q: %w", outFile, err)
	}

	fmt.Printf("Successfully generated JSON Schema for type %q in %s\n", rootType, outFile)
	return nil
}

// ConvertToJSONSchema converts PackageInfo to JSON Schema with only the root type and its (internal and external) dependencies.
func ConvertToJSONSchema(pkgInfo *model.PackageInfo, repoDir, rootType string, config *RequiredFieldsConfig) (map[string]interface{}, error) {
	schema := map[string]interface{}{
		"$schema":     "http://json-schema.org/draft-07/schema#",
		"definitions": make(map[string]interface{}),
	}
	definitions := schema["definitions"].(map[string]interface{})

	// We'll store discovered dependencies in this map
	dependencies := make(map[string]bool)

	// Build an alias mapping from the root package's imports
	aliasMap := buildAliasMap(pkgInfo.Imports)

	// Find the root type and collect its dependencies
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

	// Process internal types (no dot in their name) to generate JSON Schema definitions
	for _, decl := range pkgInfo.Declarations {
		for _, typ := range decl.Types {
			// If the type is either the rootType or a discovered dependency
			if typ.Name == rootType || dependencies[typ.Name] {
				// Only handle if it's an internal type (no dot in the name)
				if !strings.Contains(typ.Name, ".") {
					switch {
					case len(typ.Enums) > 0:
						// Enum
						definitions[typ.Name] = generateEnumSchema(typ)

					case len(typ.Fields) > 0:
						// Struct
						definitions[typ.Name] = generateStructSchema(typ, config)

					case strings.HasPrefix(typ.Type, "map["):
						// Named map type
						definitions[typ.Name] = generateMapDefinition(typ.Type)

					case strings.HasPrefix(typ.Type, "[]"):
						// Named slice/array type
						definitions[typ.Name] = generateSliceDefinition(typ.Type)

					default:
						log.Printf("Skipping type %q with underlying type %q\n", typ.Name, typ.Type)
					}
				}
			}
		}
	}

	// Process external dependencies recursively
	visited := make(map[string]bool)
	for dep := range dependencies {
		if strings.Contains(dep, ".") {
			if err := processExternalType(dep, repoDir, aliasMap, definitions, visited); err != nil {
				// Log a warning but keep going
				fmt.Fprintf(os.Stderr, "warning: %v\n", err)
			}
		}
	}

	// Point the top-level $ref to our root type
	schema["$ref"] = "#/definitions/" + rootType
	return schema, nil
}

// processExternalType loads an external package for a qualified type (e.g. "model.Inner"),
// generates its JSON Schema definition, and then recursively processes its custom fields.
func processExternalType(qualifiedType, repoDir string, aliasMap map[string]string, definitions map[string]interface{}, visited map[string]bool) error {
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

	// Lookup the import path from our alias map
	pkgPath, ok := aliasMap[pkgAlias]
	if !ok {
		return fmt.Errorf("alias %q not found in alias map", pkgAlias)
	}
	log.Println("Loading external package:", pkgPath)

	// Load the external package
	extPkgInfo, err := loadExternalPackage(pkgPath, repoDir)
	if err != nil {
		return fmt.Errorf("failed to load external package %q: %w", pkgPath, err)
	}

	// Build an alias map for the external package
	extAliasMap := buildAliasMap(extPkgInfo.Imports)
	// Also add an entry for the external package’s own name
	extAliasMap[extPkgInfo.Name] = pkgPath

	// Find the type in the external package
	var extType *model.TypeInfo
	for _, decl := range extPkgInfo.Declarations {
		for _, t := range decl.Types {
			if t.Name == typeName {
				extType = t
				break
			}
		}
	}
	if extType == nil {
		return fmt.Errorf("type %q not found in external package %q", typeName, pkgPath)
	}

	// Generate the JSON Schema for this external type
	var extSchema map[string]interface{}
	switch {
	case len(extType.Enums) > 0:
		// External enum
		extSchema = generateEnumSchema(extType)

	case len(extType.Fields) > 0:
		// External struct
		extSchema = generateStructSchema(extType, &RequiredFieldsConfig{Fields: map[string][]string{}})

	case strings.HasPrefix(extType.Type, "map["):
		extSchema = generateMapDefinition(extType.Type)

	case strings.HasPrefix(extType.Type, "[]"):
		extSchema = generateSliceDefinition(extType.Type)

	default:
		log.Printf("Skipping external type %q (type: %q)\n", typeName, extType.Type)
	}

	// If we generated a schema, store it in definitions
	if extSchema != nil {
		// Use the fully qualified name as the key
		definitions[qualifiedType] = extSchema
	}

	// Recursively process each custom field in the external type
	for _, field := range extType.Fields {
		baseType := getBaseType(field.Type)
		if isCustomType(baseType) {
			var depQualified string
			// If the field's type is already qualified, use it
			// otherwise qualify it with pkgAlias
			if strings.Contains(baseType, ".") {
				depQualified = baseType
			} else {
				depQualified = pkgAlias + "." + baseType
			}
			if err := processExternalType(depQualified, repoDir, extAliasMap, definitions, visited); err != nil {
				return err
			}
		}
	}

	return nil
}

// loadExternalPackage uses golang.org/x/tools/go/packages to load a package from its import path
// and then runs your extraction process on it.
func loadExternalPackage(pkgPath, repoDir string) (*model.PackageInfo, error) {
	absDir, err := filepath.Abs(repoDir)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps,
		Dir:  absDir,
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

	// Use the directory of the first Go file in the package
	pkgDir := filepath.Dir(pkgs[0].GoFiles[0])
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
			// No alias; deduce one from the import path
			impPath := strings.Trim(imp, "\"")
			segs := strings.Split(impPath, "/")
			aliasMap[segs[len(segs)-1]] = impPath
		}
	}
	return aliasMap
}

// collectDependencies recursively collects type dependencies from a given struct type's fields.
// If a field references a named custom type, we also parse that named type's definition
// to discover deeper dependencies (like "CertData" inside "CertsData").
func collectDependencies(typeInfo *model.TypeInfo, pkgInfo *model.PackageInfo, dependencies map[string]bool) {
	for _, field := range typeInfo.Fields {
		baseType := getBaseType(field.Type)
		if isCustomType(baseType) {
			if !dependencies[baseType] {
				dependencies[baseType] = true
				// If it's an internal type, we find its definition and parse deeper
				if !strings.Contains(baseType, ".") {
					// e.g. baseType == "CertsData"
					for _, decl := range pkgInfo.Declarations {
						for _, depType := range decl.Types {
							if depType.Name == baseType {
								// This named type might be "[]CertData", "map[string]PortWhiteList", etc.
								collectTypeDefinitionDeps(depType, pkgInfo, dependencies)
							}
						}
					}
				}
			}
		}
	}
}

// collectTypeDefinitionDeps inspects a named type's underlying type
// (e.g. "CertsData" -> "[]CertData") to find further dependencies.
func collectTypeDefinitionDeps(typeInfo *model.TypeInfo, pkgInfo *model.PackageInfo, dependencies map[string]bool) {
	underlying := typeInfo.Type

	// 1) If it's a slice: e.g. "[]CertData"
	if strings.HasPrefix(underlying, "[]") {
		elemType := strings.TrimPrefix(underlying, "[]")
		elemType = strings.TrimPrefix(elemType, "*")
		if isCustomType(elemType) {
			if !dependencies[elemType] {
				dependencies[elemType] = true
				if !strings.Contains(elemType, ".") {
					for _, decl := range pkgInfo.Declarations {
						for _, depType := range decl.Types {
							if depType.Name == elemType {
								collectTypeDefinitionDeps(depType, pkgInfo, dependencies)
							}
						}
					}
				}
			}
		}
		return
	}

	// 2) If it's a map: e.g. "map[string]PortWhiteList"
	if strings.HasPrefix(underlying, "map[") {
		inside := underlying[len("map["):]
		parts := strings.SplitN(inside, "]", 2)
		if len(parts) == 2 {
			valueType := strings.TrimSpace(parts[1])
			valueType = strings.TrimPrefix(valueType, "*")
			if isCustomType(valueType) {
				if !dependencies[valueType] {
					dependencies[valueType] = true
					if !strings.Contains(valueType, ".") {
						for _, decl := range pkgInfo.Declarations {
							for _, depType := range decl.Types {
								if depType.Name == valueType {
									collectTypeDefinitionDeps(depType, pkgInfo, dependencies)
								}
							}
						}
					}
				}
			}
		}
		return
	}

	// 3) If it's a struct (fields > 0), we do the usual field-based check.
	if len(typeInfo.Fields) > 0 {
		for _, field := range typeInfo.Fields {
			baseType := getBaseType(field.Type)
			if isCustomType(baseType) {
				if !dependencies[baseType] {
					dependencies[baseType] = true
					if !strings.Contains(baseType, ".") {
						for _, decl := range pkgInfo.Declarations {
							for _, depType := range decl.Types {
								if depType.Name == baseType {
									collectTypeDefinitionDeps(depType, pkgInfo, dependencies)
								}
							}
						}
					}
				}
			}
		}
		return
	}

	// 4) If it's an enum, built-in alias, or something else, do nothing.
}

// generateEnumSchema creates a JSON Schema definition for an enum type.
func generateEnumSchema(typeInfo *model.TypeInfo) map[string]interface{} {
	enumValues := make([]interface{}, 0)
	for _, enum := range typeInfo.Enums {
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

// generateMapDefinition creates a top-level JSON Schema definition for a map type (e.g. map[string]Something).
func generateMapDefinition(goType string) map[string]interface{} {
	// Example: "map[string]interface{}" or "map[string]PortWhiteList"
	inside := goType[len("map["):]
	parts := strings.SplitN(inside, "]", 2)
	if len(parts) != 2 {
		// fallback to a generic object
		return map[string]interface{}{
			"type":                 "object",
			"additionalProperties": true,
		}
	}
	keyType := strings.TrimSpace(parts[0])   // e.g. "string"
	valueType := strings.TrimSpace(parts[1]) // e.g. "interface{}" or "PortWhiteList"

	// JSON only supports string keys as standard objects.
	if keyType != "string" {
		return map[string]interface{}{
			"type":                 "object",
			"additionalProperties": true,
		}
	}

	// If the value is interface{} or any => no constraints
	if valueType == "interface{}" || valueType == "any" {
		return map[string]interface{}{
			"type":                 "object",
			"additionalProperties": true,
		}
	}

	// If it's built-in, produce a simple type
	if !isCustomType(valueType) {
		return map[string]interface{}{
			"type": "object",
			"additionalProperties": map[string]interface{}{
				"type": getBaseJSONType(valueType),
			},
		}
	}

	// Otherwise it's a custom type => reference
	return map[string]interface{}{
		"type": "object",
		"additionalProperties": map[string]interface{}{
			"$ref": "#/definitions/" + valueType,
		},
	}
}

// generateSliceDefinition creates a top-level JSON Schema definition for a slice type (e.g. []CertData).
func generateSliceDefinition(goType string) map[string]interface{} {
	// e.g. "[]CertData"
	elemType := strings.TrimPrefix(goType, "[]")
	elemType = strings.TrimSpace(elemType)

	// If it's built-in:
	if !isCustomType(elemType) {
		return map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": getBaseJSONType(elemType),
			},
		}
	}

	// Otherwise custom
	return map[string]interface{}{
		"type": "array",
		"items": map[string]interface{}{
			"$ref": "#/definitions/" + elemType,
		},
	}
}

// getJSONType converts a Go type into its JSON Schema representation (for fields).
func getJSONType(goType string) map[string]interface{} {
	// handle []byte specially
	if goType == "[]byte" {
		return map[string]interface{}{
			"type":   "string",
			"format": "byte",
		}
	}

	// handle slices
	if strings.HasPrefix(goType, "[]") {
		elementType := strings.TrimPrefix(goType, "[]")
		if isCustomType(elementType) {
			return map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"$ref": "#/definitions/" + elementType,
				},
			}
		}
		return map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": getBaseJSONType(elementType),
			},
		}
	}

	// handle maps
	if strings.HasPrefix(goType, "map[") {
		inside := goType[len("map["):]
		parts := strings.SplitN(inside, "]", 2)
		if len(parts) != 2 {
			return map[string]interface{}{
				"type":                 "object",
				"additionalProperties": true,
			}
		}
		keyType := strings.TrimSpace(parts[0])
		valueType := strings.TrimSpace(parts[1])

		if keyType != "string" {
			return map[string]interface{}{
				"type":                 "object",
				"additionalProperties": true,
			}
		}
		if valueType == "interface{}" || valueType == "any" {
			return map[string]interface{}{
				"type":                 "object",
				"additionalProperties": true,
			}
		}
		if !isCustomType(valueType) {
			return map[string]interface{}{
				"type": "object",
				"additionalProperties": map[string]interface{}{
					"type": getBaseJSONType(valueType),
				},
			}
		}
		// custom type => $ref
		return map[string]interface{}{
			"type": "object",
			"additionalProperties": map[string]interface{}{
				"$ref": "#/definitions/" + valueType,
			},
		}
	}

	// fallback for non-array, non-map
	schema := map[string]interface{}{
		"type": getBaseJSONType(goType),
	}

	// numeric constraints
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
	// For direct map[...] or []byte, we skip
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
	// Otherwise assume it's custom
	return true
}

// RequiredFieldsConfig defines which fields are required for each type.
type RequiredFieldsConfig struct {
	Fields map[string][]string // map[TypeName][]FieldName
}

// NewDefaultConfig just returns a sample required-fields config
func NewDefaultConfig() *RequiredFieldsConfig {
	return &RequiredFieldsConfig{
		Fields: map[string][]string{
			"User":  {"ID", "Name"}, // Only ID and Name are required for User
			"Inner": {"Name"},
		},
	}
}
