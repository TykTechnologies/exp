package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func InsertRequest(r io.Reader, args []string) (Records, error) {
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

func buildInsertQuery(table string, data Record) (string, []any) {
	// Step 1: List the keys
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	// Step 2: Construct placeholders string
	placeholders := strings.Repeat("?, ", len(keys))
	placeholders = strings.TrimSuffix(placeholders, ", ")

	// Step 3: Create a slice of values
	values := make([]any, 0, len(data))
	for _, key := range keys {
		values = append(values, data[key])
	}

	// Step 4: Construct ON DUPLICATE KEY UPDATE clause
	updates := make([]string, 0, len(data))
	for _, key := range keys {
		updates = append(updates, fmt.Sprintf("%s = VALUES(%s)", key, key))
	}

	// Step 5: Create the query string
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s",
		table, strings.Join(keys, ", "), placeholders, strings.Join(updates, ", "))

	return query, values
}

func Insert(ctx context.Context, command *Command, r io.Reader) error {
	records, err := InsertRequest(r, command.Args[1:])
	if err != nil {
		return err
	}

	table := command.Args[0]

	var names []string
	var values []string
	for _, arg := range command.Args[1:] {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			parts[1] = strings.Trim(parts[1], "'\"")

			if strings.HasPrefix(parts[1], "@") {
				contents, err := os.ReadFile(parts[1][1:])
				if err != nil {
					return err
				}
				parts[1] = string(contents)
			}

			names = append(names, parts[0])
			values = append(values, parts[1])
			continue
		}
	}

	// append with commandline args
	for k, name := range names {
		for i, r := range records {
			r[name] = values[k]
			records[i] = r
		}
	}

	tx, err := command.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var inserts, updates int64
	for _, record := range records {
		query, params := buildInsertQuery(table, record)

		if command.Verbose {
			fmt.Printf("-- %s %#v\n", query, params)
		}

		result, err := tx.Exec(query, params...)
		if err != nil {
			return err
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected > 1 {
			updates++
		} else {
			inserts++
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	fmt.Printf("%d inserts, %d updates\n", inserts, updates)
	return nil
}
