package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"golang.org/x/exp/maps"
)

func UpdateRequest(r io.Reader, args []string) (Records, error) {
	input, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	multi := false
	if string(input[0:1]) == "[" {
		multi = true
	}

	var records Records
	if multi {
		if err := json.Unmarshal(input, &records); err != nil {
			return nil, err
		}
	} else {
		var record Record
		if err := json.Unmarshal(input, &record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func buildUpdateQuery(table string, data Record, whereKeys []string) (string, []any) {
	var query = fmt.Sprintf("UPDATE %s SET ", table)

	keys := maps.Keys(data)

	values := []any{}
	for i, key := range keys {
		prefix := ", "
		if i == 0 {
			prefix = ""
		}
		query += fmt.Sprintf("%s%s=? ", prefix, key)
		values = append(values, data[key])
	}
	query += "WHERE "
	for i, key := range whereKeys {
		prefix := ", "
		if i == 0 {
			prefix = ""
			query += fmt.Sprintf("%s%s=? ", prefix, key)
		}
		values = append(values, data[key])
	}

	fmt.Println(query)

	return query, values
}

func Update(ctx context.Context, command *Command, r io.Reader) error {
	records, err := UpdateRequest(r, command.Args[1:])
	if err != nil {
		return err
	}

	table := command.Args[0]
	where := command.Args[1:]

	tx, err := command.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	count := 0
	for _, record := range records {
		query, params := buildUpdateQuery(table, record, where)
		_, err := tx.Exec(query, params...)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				continue // Ignore unique constraint errors
			}
			if strings.Contains(err.Error(), "Duplicate entry") {
				continue // Ignore unique constraint errors
			}
			return err
		}
		count++
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	fmt.Printf("%d rows updated\n", count)
	return nil
}
