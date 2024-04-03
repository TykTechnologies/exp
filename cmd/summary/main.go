package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	_ "modernc.org/sqlite"

	"golang.org/x/exp/maps"

	"github.com/TykTechnologies/exp/cmd/summary/golangcilint"
	"github.com/TykTechnologies/exp/cmd/summary/lsof"
	"github.com/TykTechnologies/exp/cmd/summary/vet"
)

func main() {
	if err := start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() (err error) {
	commands := map[string]func() error{
		"vet":          vet.Run,
		"lsof":         lsof.Run,
		"golangcilint": golangcilint.Run,
	}
	commandList := maps.Keys(commands)
	sort.Strings(commandList)

	if len(os.Args) < 2 {
		fmt.Println("Usage: summary <command> help")
		fmt.Printf("Available commands: %s\n", strings.Join(commandList, ", "))
		return nil
	}

	commandFn, ok := commands[os.Args[1]]
	if ok {
		return commandFn()
	}

	return fmt.Errorf("Unknown command: %q", os.Args[1])
}
