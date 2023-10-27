package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	_ "modernc.org/sqlite"

	"golang.org/x/exp/maps"

	"github.com/TykTechnologies/exp/cmd/go-fsck/extract"
	"github.com/TykTechnologies/exp/cmd/go-fsck/lint"
	"github.com/TykTechnologies/exp/cmd/go-fsck/restore"
	"github.com/TykTechnologies/exp/cmd/go-fsck/stats"
)

func main() {
	if err := start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() (err error) {
	commands := map[string]func() error{
		"extract": extract.Run,
		"restore": restore.Run,
		"stats":   stats.Run,
		"lint":    lint.Run,
	}
	commandList := maps.Keys(commands)
	sort.Strings(commandList)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go-fsck <command> help")
		fmt.Printf("Available commands: %s\n", strings.Join(commandList, ", "))
		return nil
	}

	commandFn, ok := commands[os.Args[1]]
	if ok {
		return commandFn()
	}

	return fmt.Errorf("Unknown command: %q", os.Args[1])
}
