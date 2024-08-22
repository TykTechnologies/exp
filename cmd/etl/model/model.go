package model

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

// Record represents the json encoding output.
type Record map[string]string

// RecordInput represents the named query parameter input.
type RecordInput map[string]any

// Record encodes RecordInput into Record.
func (f RecordInput) Record() Record {
	result := make(Record, len(f))
	for k, v := range f {
		result[strings.ToLower(k)] = dbValue(v)
	}
	return result
}

type Command struct {
	Name string
	Args []string
	DB   *sqlx.DB

	Verbose bool
	Quiet   bool
}

type TableInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Count       int      `json:"count"`
	Columns     []Record `json:"columns,omitempty"`
}
