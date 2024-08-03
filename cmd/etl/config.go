package main

import (
	"os"

	"github.com/spf13/pflag"
)

type Config struct {
	DSN     string
	Driver  string
	Folder  string
	Verbose bool
}

func NewConfig() (*Config, error) {
	var config Config
	pflag.StringVar(&config.DSN, "db-dsn", os.Getenv("DB_DSN"), "Database DSN")
	pflag.StringVar(&config.Driver, "db-driver", os.Getenv("DB_DRIVER"), "Database Driver")
	pflag.StringVarP(&config.Folder, "folder", "f", "output", "Folder with outputs")
	pflag.BoolVarP(&config.Verbose, "verbose", "v", false, "Folder with outputs")
	pflag.Parse()
	return &config, nil
}

func (c *Config) GetDSN() string {
	return c.DSN + "?parseTime=true"
}

func (c *Config) GetDriver() string {
	return c.Driver
}

var config = &Config{}
