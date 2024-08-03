package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

func Get(ctx context.Context, command *Command) error {
	var first bool
	pflag.BoolVar(&first, "first", true, "Return first record")
	pflag.Parse()

	table := command.Args[0]

	var values []any
	var params []string
	for _, arg := range command.Args[1:] {
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
	if first {
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
