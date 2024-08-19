package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/TykTechnologies/exp/cmd/etl/model"
)

func Query(ctx context.Context, command *model.Command, _ io.Reader) error {
	flagSet := NewFlagSet("Query")
	if err := flagSet.Parse(command.Args); err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}
	args := flagSet.Args()

	query, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}

	params, err := decodeQueryParameters(args[1:])
	if err != nil {
		return err
	}

	if command.Verbose {
		log.Printf("-- %s %#v\n", query, params)
	}

	rows, err := command.DB.NamedQuery(string(query), params)
	if err != nil {
		return err
	}
	defer rows.Close()

	results, err := scanAllRecords(rows)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(results)
}
