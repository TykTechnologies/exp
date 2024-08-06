package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func Get(ctx context.Context, command *Command, _ io.Reader) error {
	var all bool

	flagSet := NewFlagSet("Get")
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

			if parts[1] == "NULL" {
				params = append(params, parts[0]+" IS NULL")
				continue
			}

			params = append(params, parts[0]+" = ?")
			values = append(values, parts[1])
			continue
		}
		params = append(params, arg)
	}

	if len(params) == 0 {
		params = append(params, "1=1")
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", table, strings.Join(params, " "))
	if !all {
		query = query + " LIMIT 0,1"
	}
	if command.Verbose {
		fmt.Printf("-- %s %#v %#v\n", query, params, values)
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

	if len(results) == 0 {
		os.Exit(1)
	}

	var data any = results
	if !all {
		data = &results[0]
	}

	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
