package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := start(ctx); err != nil {
		log.Fatal(err)
	}
}

func start(ctx context.Context) error {
	if len(os.Args) < 2 {
		return errors.New("usage: etl <command> <tableName> [options]")
	}

	config := NewConfig()
	args, err := config.ParseFlags()
	if err != nil {
		return err
	}

	db, err := sqlx.Open(config.GetDriver(), config.GetDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	command := Command{
		DB:      db,
		Name:    args[0],
		Args:    args[1:],
		Verbose: config.Verbose,
	}

	return HandleCommand(ctx, &command, os.Stdin)
}
