package main

import (
	"strings"
)

type TypeInfo struct {
	Type, Format string
}

func TypeAlias(kind string) (TypeInfo, bool) {
	typeInfo, ok := typeAliases[kind]
	if ok {
		return typeInfo, ok
	}
	typeInfo.Type = kind

	if strings.HasPrefix(kind, "[]") {
		format := kind[2:]
		return TypeInfo{
			Type:   "array",
			Format: format,
		}, true
	}

	if strings.HasPrefix(kind, "map[") {
		format := strings.Split(kind, "]")[1]
		return TypeInfo{
			Type:   "map",
			Format: format,
		}, true
	}

	return typeInfo, false
}

var typeAliases = map[string]TypeInfo{
	// rich types
	"[]string":               {Type: "array", Format: "string"},
	"map[string]interface{}": {Type: "object"},
	"time.Time":              {Type: "string", Format: "date-time"},
	"*time.Time":             {Type: "string", Format: "date-time"},
	"string":                 {Type: "string"},

	// proto numeric types
	"int":      {Type: "integer"},
	"int32":    {Type: "integer", Format: "int32"},
	"uint32":   {Type: "integer", Format: "uint32"},
	"sint32":   {Type: "integer", Format: "int32"},
	"fixed32":  {Type: "integer", Format: "int32"},
	"sfixed32": {Type: "integer", Format: "int32"},

	// proto numeric types, 64bit
	"int64":    {Type: "integer", Format: "int64"},
	"uint64":   {Type: "integer", Format: "uint64"},
	"sint64":   {Type: "integer", Format: "int64"},
	"fixed64":  {Type: "integer", Format: "int64"},
	"sfixed64": {Type: "integer", Format: "int64"},

	"double":  {Type: "number", Format: "double"},
	"float":   {Type: "number", Format: "float"},
	"float64": {Type: "number", Format: "double"},

	// effectively copies google.protobuf.BytesValue
	"bytes": {
		Type:   "string",
		Format: "byte",
	},

	// It is what it is
	"bool": {
		Type:   "boolean",
		Format: "boolean",
	},

	"google.protobuf.Timestamp": {
		Type:   "string",
		Format: "date-time",
	},
	"google.protobuf.Duration": {
		Type: "string",
	},
	"google.protobuf.StringValue": {
		Type: "string",
	},
	"google.protobuf.BytesValue": {
		Type:   "string",
		Format: "byte",
	},
	"google.protobuf.Int32Value": {
		Type:   "integer",
		Format: "int32",
	},
	"google.protobuf.UInt32Value": {
		Type:   "integer",
		Format: "uint32",
	},
	"google.protobuf.Int64Value": {
		Type:   "string",
		Format: "int64",
	},
	"google.protobuf.UInt64Value": {
		Type:   "string",
		Format: "uint64",
	},
	"google.protobuf.FloatValue": {
		Type:   "number",
		Format: "float",
	},
	"google.protobuf.DoubleValue": {
		Type:   "number",
		Format: "double",
	},
	"google.protobuf.BoolValue": {
		Type:   "boolean",
		Format: "boolean",
	},
	"google.protobuf.Empty": {},
}
