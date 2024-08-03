package main

import (
	"github.com/jmoiron/sqlx"
)

type Record map[string]string
type Records []Record

type Command struct {
	Name    string
	Args    []string
	DB      *sqlx.DB
	Verbose bool
}

type TableInfo struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Columns     []ColumnInfo `json:"columns"`
}

type ColumnInfo map[string]string
