package main

// This app parses structs from a package in order to extract
// known data model struct fields into a json document.

import (
	"fmt"
	"os"

	"github.com/TykTechnologies/exp/cmd/schema-gen/extract"
)

func main() {
	if err := start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() (err error) {
	generators := []func() error{
		extract.Dump,
	}
	for _, generator := range generators {
		if err := generator(); err != nil {
			return err
		}
	}
	return nil
}
