package jsonschema

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
	if outFile == "" {
		outFile = "schema.json"
	}

	absDir, err := filepath.Abs(repoDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for %q: %w", repoDir, err)
	}

	pkgInfos, err := extract.Extract(absDir+"/", &extract.ExtractOptions{})
	if err != nil {
		return fmt.Errorf("failed to extract types from %q: %w", absDir, err)
	}
	if len(pkgInfos) == 0 {
		return fmt.Errorf("no package info extracted from %q", absDir)
	}

	schema, err := ConvertToJSONSchema(pkgInfos[0], absDir, rootType, NewDefaultConfig())
	if err != nil {
		return fmt.Errorf("failed to convert to JSON Schema: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal schema to JSON: %w", err)
	}
	if err := os.WriteFile(outFile, jsonBytes, 0o644); err != nil {
		return fmt.Errorf("failed to write %q: %w", outFile, err)
	}

	return nil
}

// ConvertToJSONSchema converts PackageInfo to JSON Schema with only the root type and its (internal and external) dependencies.
func ConvertToJSONSchema(pkgInfo *model.PackageInfo, repoDir, rootType string, config *RequiredFieldsConfig) (*JSONSchema, error) {
	rootSchema := &JSONSchema{
		Schema:      "http://json-schema.org/draft-07/schema#",
		Definitions: make(map[string]*JSONSchema),
	}
	definitions := rootSchema.Definitions

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
				CollectDependencies(typ, pkgInfo, dependencies)
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
						definitions[typ.Name] = GenerateEnumSchema(typ)

					case len(typ.Fields) > 0:
						// Struct
						definitions[typ.Name] = GenerateStructSchema(typ, config)

					case strings.HasPrefix(typ.Type, "map["):

						definitions[typ.Name] = GenerateMapDefinition(typ.Type)

					case strings.HasPrefix(typ.Type, "[]"):
						definitions[typ.Name] = GenerateSliceDefinition(typ.Type)

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
			if err := ProcessExternalType(dep, repoDir, aliasMap, definitions, visited); err != nil {
				fmt.Fprintf(os.Stderr, "warning: %v\n", err)
			}
		}
	}

	rootSchema.Ref = "#/definitions/" + rootType
	return rootSchema, nil
}

// ProcessExternalType loads an external package for a qualified type (e.g. "model.Inner"),
// generates its JSON Schema definition, and then recursively processes its custom fields.
func ProcessExternalType(qualifiedType, repoDir string, aliasMap map[string]string, definitions map[string]*JSONSchema, visited map[string]bool) error {
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
	// Load the external package
	extPkgInfo, err := LoadExternalPackage(pkgPath, repoDir)
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

	var extSchema *JSONSchema
	switch {
	case len(extType.Enums) > 0:
		extSchema = GenerateEnumSchema(extType)

	case len(extType.Fields) > 0:
		extSchema = GenerateStructSchema(extType, &RequiredFieldsConfig{Fields: map[string][]string{}})

	case strings.HasPrefix(extType.Type, "map["):
		extSchema = GenerateMapDefinition(extType.Type)

	case strings.HasPrefix(extType.Type, "[]"):
		extSchema = GenerateSliceDefinition(extType.Type)

	default:
		log.Printf("Skipping external type %q (type: %q)\n", typeName, extType.Type)
	}

	if extSchema != nil {
		definitions[qualifiedType] = extSchema
	}
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
			if err := ProcessExternalType(depQualified, repoDir, extAliasMap, definitions, visited); err != nil {
				return err
			}
		}
	}

	return nil
}

// LoadExternalPackage uses golang.org/x/tools/go/packages to load a package from its import path
// and then runs the extraction process on it.
func LoadExternalPackage(pkgPath, repoDir string) (*model.PackageInfo, error) {
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

// CollectDependencies recursively collects type dependencies from a given struct type's fields.
// If a field references a named custom type, we also parse that named type's definition
// to discover deeper dependencies (like "CertData" inside "CertsData").
func CollectDependencies(typeInfo *model.TypeInfo, pkgInfo *model.PackageInfo, dependencies map[string]bool) {
	for _, field := range typeInfo.Fields {
		if strings.HasPrefix(field.Type, "map[") {
			handleMapField(field.Type, pkgInfo, dependencies)
			continue
		}
		baseType := getBaseType(field.Type)
		if isCustomType(baseType) {
			if !dependencies[baseType] {
				dependencies[baseType] = true
				if !strings.Contains(baseType, ".") {
					// e.g. baseType == "CertsData"
					for _, decl := range pkgInfo.Declarations {
						for _, depType := range decl.Types {
							if depType.Name == baseType {
								// This named type might be "[]CertData", "map[string]PortWhiteList", etc.
								CollectTypeDefinitionDeps(depType, pkgInfo, dependencies)
							}
						}
					}
				}
			}
		}
	}
}

// CollectTypeDefinitionDeps inspects a named type's underlying type
// (e.g. "CertsData" -> "[]CertData") to find further dependencies.
func CollectTypeDefinitionDeps(typeInfo *model.TypeInfo, pkgInfo *model.PackageInfo, dependencies map[string]bool) {
	underlying := typeInfo.Type
	// If it's a slice: e.g. "[]CertData"
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
								CollectTypeDefinitionDeps(depType, pkgInfo, dependencies)
							}
						}
					}
				}
			}
		}
		return
	}

	// If it's a map: e.g. "map[string]PortWhiteList"
	if strings.HasPrefix(underlying, "map[") {
		handleMapField(underlying, pkgInfo, dependencies)
		return
	}

	// If it's a struct (fields > 0), we do the usual field-based check.
	if len(typeInfo.Fields) > 0 {
		for _, field := range typeInfo.Fields {
			if strings.HasPrefix(field.Type, "map[") {
				handleMapField(field.Type, pkgInfo, dependencies)
				continue
			}
			baseType := getBaseType(field.Type)
			if isCustomType(baseType) {
				if !dependencies[baseType] {
					dependencies[baseType] = true
					if !strings.Contains(baseType, ".") {
						for _, decl := range pkgInfo.Declarations {
							for _, depType := range decl.Types {
								if depType.Name == baseType {
									CollectTypeDefinitionDeps(depType, pkgInfo, dependencies)
								}
							}
						}
					}
				}
			} else {
				fmt.Println(baseType)
			}
		}
		return
	}

	// If it's an enum, built-in alias, or something else, do nothing.
}

// GenerateEnumSchema creates a JSON Schema definition for an enum type.
func GenerateEnumSchema(typeInfo *model.TypeInfo) *JSONSchema {
	enumValues := make([]any, 0, len(typeInfo.Enums))
	for _, enum := range typeInfo.Enums {
		enumValues = append(enumValues, enum.Value)
	}
	jsonType := "string"
	if typeInfo.Type == "int" {
		jsonType = "integer"
	}
	return &JSONSchema{
		Type: jsonType,
		Enum: enumValues,
	}
}

// GenerateStructSchema creates a JSON Schema definition for a struct type.
func GenerateStructSchema(typeInfo *model.TypeInfo, config *RequiredFieldsConfig) *JSONSchema {
	schema := &JSONSchema{
		Type:       "object",
		Properties: make(map[string]*JSONSchema),
		// Typically, "additionalProperties" is either `false` or another schema
		AdditionalProperties: false,
	}

	requiredFields := config.Fields[typeInfo.Name]
	requiredMap := make(map[string]bool)
	for _, field := range requiredFields {
		requiredMap[field] = true
	}
	var required []string

	for _, field := range typeInfo.Fields {

		if field.JSONName == "-" {
			continue
		}

		isArray := strings.HasPrefix(field.Type, "[]")
		baseType := strings.TrimPrefix(field.Type, "[]")
		var fieldSchema *JSONSchema
		if isCustomType(baseType) {
			if isArray {
				fieldSchema = &JSONSchema{
					Type: "array",
					Items: &JSONSchema{
						Ref: "#/definitions/" + baseType,
					},
				}
			} else {
				fieldSchema = &JSONSchema{
					Ref: "#/definitions/" + baseType,
				}
			}
		} else {
			fieldSchema = getJSONType(field.Type)
		}
		if field.Doc != "" {
			fieldSchema.Description = field.Doc
		}
		cleanedJson := parseJSONTag(field.JSONName)
		schema.Properties[cleanedJson] = fieldSchema
		if requiredMap[field.Name] {
			required = append(required, cleanedJson)
		}
	}
	if len(required) > 0 {
		schema.Required = required
	}
	return schema
}

// GenerateMapDefinition creates a top-level JSON Schema definition for a map type (e.g. map[string]Something).
func GenerateMapDefinition(goType string) *JSONSchema {
	// Example: "map[string]interface{}" or "map[string]PortWhiteList"
	inside := goType[len("map["):]
	parts := strings.SplitN(inside, "]", 2)
	if len(parts) != 2 {
		// fallback to a generic object
		return &JSONSchema{
			Type:                 "object",
			AdditionalProperties: true,
		}
	}
	keyType := strings.TrimSpace(parts[0])   // e.g. "string"
	valueType := strings.TrimSpace(parts[1]) // e.g. "interface{}" or "PortWhiteList"

	// JSON only supports string keys as standard objects.
	if keyType != "string" {
		return &JSONSchema{
			Type:                 "object",
			AdditionalProperties: true,
		}
	}

	// If the value is interface{} or any => no constraints
	if valueType == "interface{}" || valueType == "any" {
		return &JSONSchema{
			Type:                 "object",
			AdditionalProperties: true,
		}
	}

	// If it's built-in, produce a simple type
	if !isCustomType(valueType) {
		return &JSONSchema{
			Type: "object",
			AdditionalProperties: &JSONSchema{
				Type: getBaseJSONType(valueType),
			},
		}
	}

	// Otherwise it's a custom type => reference
	return &JSONSchema{
		Type: "object",
		AdditionalProperties: &JSONSchema{
			Ref: "#/definitions/" + valueType,
		},
	}
}

// GenerateSliceDefinition creates a top-level JSON Schema definition for a slice type (e.g. []CertData).
func GenerateSliceDefinition(goType string) *JSONSchema {
	// e.g. "[]CertData"
	elemType := strings.TrimPrefix(goType, "[]")
	elemType = strings.TrimSpace(elemType)

	if !isCustomType(elemType) {
		return &JSONSchema{
			Type: "array",
			Items: &JSONSchema{
				Type: getBaseJSONType(elemType),
			},
		}
	}

	return &JSONSchema{
		Type: "array",
		Items: &JSONSchema{
			Ref: "#/definitions/" + elemType,
		},
	}
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

func handleMapField(fieldType string, pkgInfo *model.PackageInfo, dependencies map[string]bool) {
	// e.g. "map[string]RequestHeadersRewriteConfig"
	inside := fieldType[len("map["):]
	parts := strings.SplitN(inside, "]", 2)
	if len(parts) != 2 {
		return
	}
	valueType := strings.TrimSpace(parts[1])
	valueType = strings.TrimPrefix(valueType, "*")

	if isCustomType(valueType) {
		if !dependencies[valueType] {
			dependencies[valueType] = true
			if !strings.Contains(valueType, ".") {
				// find its definition, parse deeper
				for _, decl := range pkgInfo.Declarations {
					for _, depType := range decl.Types {
						if depType.Name == valueType {
							CollectTypeDefinitionDeps(depType, pkgInfo, dependencies)
						}
					}
				}
			}
		}
	}
}
