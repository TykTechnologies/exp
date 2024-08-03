package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/pflag"
)

// Tables retrieves the list of tables in the current database schema along with their comments.
// If the Verbose flag in the command is set to true, it also retrieves the details of each table's columns.
func Tables(ctx context.Context, command *Command) error {
	tables, err := getTableList(command.DB, command.Verbose)
	if err != nil {
		return err
	}

	var showColumns bool
	pflag.BoolVar(&showColumns, "show-columns", false, "Show columns")
	pflag.Parse()

	if showColumns {
		for i, table := range tables {
			columns, err := getTableColumns(command.DB, table.Name, command.Verbose)
			if err != nil {
				return err
			}
			tables[i].Columns = columns
		}
	}

	output, err := json.Marshal(tables)
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}

func getTableList(db *sqlx.DB, verbose bool) ([]TableInfo, error) {
	var (
		query  = "select %s from information_schema.tables where table_schema=database()"
		fields = "table_name, table_comment"
	)

	rows, err := db.Queryx(fmt.Sprintf(query, fields))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []TableInfo
	for rows.Next() {
		var tableName, tableComment string
		if err := rows.Scan(&tableName, &tableComment); err != nil {
			return nil, err
		}

		table := TableInfo{
			Name:        tableName,
			Description: tableComment,
		}
		tables = append(tables, table)
	}

	return tables, nil
}

func getTableColumns(db *sqlx.DB, tableName string, verbose bool) ([]ColumnInfo, error) {
	var (
		query  = "select %s from information_schema.columns where table_schema=database() and table_name=?"
		fields = "column_name, column_type, column_comment"
	)
	if verbose {
		fields = "*"
	}

	rows, err := db.Queryx(fmt.Sprintf(query, fields), tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		column := make(map[string]interface{})
		if err := rows.MapScan(column); err != nil {
			return nil, err
		}

		field := make(map[string]string)
		for k, v := range column {
			field[strings.ToLower(k)] = dbValue(v)
		}

		columns = append(columns, field)
	}

	return columns, nil
}

func dbValue(in any) string {
	if v, ok := in.([]byte); ok {
		return string(v)
	}
	return fmt.Sprint(in)
}
