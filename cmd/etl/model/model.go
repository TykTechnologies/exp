package model

import (
	"github.com/jmoiron/sqlx"
)

type Record map[string]string
type Records []Record

type Command struct {
	Name string
	Args []string
	DB   *sqlx.DB

	Verbose bool
	Quiet   bool
}

type TableInfo struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Count       int          `json:"count"`
	Columns     []ColumnInfo `json:"columns,omitempty"`
}

type ColumnInfo map[string]string
