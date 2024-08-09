package main

import (
	"context"
	"fmt"
	"io"

	"golang.org/x/exp/maps"

	"github.com/TykTechnologies/exp/cmd/etl/model"
)

type CommandHandlerFunc func(ctx context.Context, command *model.Command, r io.Reader) error

func HandleCommand(ctx context.Context, command *model.Command, r io.Reader) error {
	commandMap := map[string]CommandHandlerFunc{
		"insert":   Insert,
		"get":      Get,
		"list":     List,
		"tables":   Tables,
		"update":   Update,
		"query":    Query,
		"truncate": Truncate,
	}
	commands := maps.Keys(commandMap)

	if fn, ok := commandMap[command.Name]; ok {
		return fn(ctx, command, r)
	}
	return fmt.Errorf("unknown command: %s, supported %v", command.Name, commands)
}
