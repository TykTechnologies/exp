package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/pflag"
)

func List(ctx context.Context, command *Command) error {
	var offset, limit int
	pflag.IntVar(&offset, "offset", 0, "Offset for the results")
	pflag.IntVar(&limit, "limit", 10, "Limit the number of results")
	pflag.Parse()

	var err error
	var rows *sqlx.Rows
	table := command.Args[0]
	if len(command.Args) > 1 {
		query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", table)
		rows, err = command.DB.Queryx(query, command.Args[1])
	} else {
		query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT %d OFFSET %d", table, limit, offset)
		rows, err = command.DB.Queryx(query)
	}
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
	if len(command.Args) > 1 {
		var res *Record
		if len(results) > 0 {
			res = &results[0]
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
