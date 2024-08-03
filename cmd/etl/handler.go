package main

import (
	"context"
	"fmt"
	"io"
)

func HandleCommand(ctx context.Context, command *Command, r io.Reader) error {
	switch command.Name {
	case "insert":
		return Insert(ctx, command, r)
	case "get":
		return Get(ctx, command)
	case "list":
		return List(ctx, command)
	case "tables":
		return Tables(ctx, command)
	default:
		return fmt.Errorf("unknown command: %s", command.Name)
	}
}
