package jsonschema

import "strings"

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
