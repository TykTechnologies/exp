package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func decodeQueryParameters(args []string) (map[string]any, error) {
	result := map[string]any{}
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			parts[1] = strings.Trim(parts[1], "'\"")

			if strings.HasPrefix(parts[1], "@") {
				contents, err := os.ReadFile(parts[1][1:])
				if err != nil {
					return nil, err
				}
				parts[1] = string(contents)
			}

			k, v := parts[0], parts[1]
			result[k] = v
			continue
		}
	}

	return result, nil
}

func Query(ctx context.Context, command *Command, _ io.Reader) error {
	flagSet := NewFlagSet("Query")
	if err := flagSet.Parse(command.Args); err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}
	args := flagSet.Args()

	query, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}

	queryParams, err := decodeQueryParameters(args[1:])
	if err != nil {
		return err
	}

	rows, err := command.DB.NamedQuery(string(query), queryParams)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Process the rows and write the response
	results := []map[string]string{}
	for rows.Next() {
		row := map[string]interface{}{}
		if err := rows.MapScan(row); err != nil {
			return err
		}

		result := make(map[string]string, len(row))
		for k, v := range row {
			result[strings.ToLower(k)] = dbValue(v)
		}

		results = append(results, result)
	}

	return json.NewEncoder(os.Stdout).Encode(results)
}
