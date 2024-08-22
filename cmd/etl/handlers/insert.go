package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TykTechnologies/exp/cmd/etl/model"
)

// isInputFromPipe checks if there's input from a pipe
func isInputFromPipe() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error stating stdin:", err)
		return false
	}
	return fi.Mode()&os.ModeNamedPipe != 0 || fi.Size() > 0
}

func InsertRequest(r io.Reader, args []string) ([]model.RecordInput, error) {
	if !isInputFromPipe() {
		return []model.RecordInput{model.RecordInput{}}, nil
	}

	input, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	multi := false
	if string(input[0:1]) == "[" {
		multi = true
	}

	var records []model.RecordInput
	if multi {
		if err := json.Unmarshal(input, &records); err != nil {
			return nil, err
		}
	} else {
		var record model.RecordInput
		if err := json.Unmarshal(input, &record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func buildInsertQuery(table string, data model.RecordInput) (string, []any) {
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

func Insert(ctx context.Context, command *model.Command, r io.Reader) error {
	records, err := InsertRequest(r, command.Args[1:])
	if err != nil {
		return err
	}

	args := command.Args
	table := args[0]

	params, err := decodeQueryParameters(args[1:])
	if err != nil {
		return err
	}

	// append with commandline args
	for i, r := range records {
		for k, v := range params {
			r[k] = v
		}
		records[i] = r
	}

	tx, err := command.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var updates int64
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
		updates += rowsAffected
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	if !command.Quiet {
		fmt.Printf("%s: stored %d records, %d rows affected\n", table, len(records), updates)
	}
	return nil
}
