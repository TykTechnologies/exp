package main

import (
	"bytes"
	"fmt"
	"os"
)

func WriteFile(filename string, data []byte, err error) error {
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

func RenderTemplate(templateName string, data map[string]interface{}) ([]byte, error) {
	t, err := loadTemplate(templateName)
	if err != nil {
		return nil, fmt.Errorf("error loading template: %w", err)
	}

	var output bytes.Buffer
	err = t.Execute(&output, data)
	return output.Bytes(), err
}
