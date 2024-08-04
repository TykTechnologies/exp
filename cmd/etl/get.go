package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func Get(ctx context.Context, command *Command, _ io.Reader) error {
	var all bool

	flagSet := NewFlagSet("List")
	flagSet.BoolVar(&all, "all", false, "Return all records")
	if err := flagSet.Parse(command.Args); err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}
	args := flagSet.Args()

	table := args[0]

	var values []any
	var params []string
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "-") {
			continue
		}

		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			parts[1] = strings.Trim(parts[1], "'\"")

			params = append(params, parts[0]+" = ?")
			values = append(values, parts[1])
			continue
		}
		params = append(params, arg)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", table, strings.Join(params, " "))
	if command.Verbose {
		fmt.Println("--", query)
	}
	rows, err := command.DB.Queryx(query, values...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var results []Record
	for rows.Next() {
		row := make(map[string]any)
		if err := rows.MapScan(row); err != nil {
			return err
		}

		result := make(map[string]string, len(row))
		for k, v := range row {
			result[strings.ToLower(k)] = dbValue(v)
		}

		results = append(results, result)
	}

	var output []byte
	if !all {
		var res *Record
		if len(results) > 0 {
			res = &results[0]
		}
		if res == nil {
			return errors.New("no results")
		}

		output, err = json.Marshal(res)
	} else {
		output, err = json.Marshal(results)
	}

	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
