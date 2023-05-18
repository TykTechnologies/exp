package structs

import (
	"encoding/json"
	"flag"
	"os"
)

type Request struct {
	rootName    string
	pkgPath     string
	ignoreFiles []string
}

func Dump() (err error) {
	var (
		outputFile = "schema.json"
		sourcePath = "."
	)
	flag.StringVar(&outputFile, "o", outputFile, "output file")
	flag.StringVar(&sourcePath, "i", sourcePath, "source path")
	flag.Parse()

	return write(outputFile, sourcePath)
}

func write(filename string, inputPackage string) error {
	sts, err := Extract(inputPackage)
	if err != nil {
		return err
	}

	return dump(filename, sts)
}

func dump(filename string, data interface{}) error {
	println(filename)

	dataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, dataBytes, 0644) //nolint:gosec
}
