package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type fileRecord struct {
	Name     string
	Basename string
	Dir      string
}

func newFileRecord(name string) fileRecord {
	return fileRecord{
		Name:     name,
		Basename: path.Base(name),
		Dir:      path.Dir(name),
	}
}

func isImageExtension(filename string) bool {
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".jpeg", ".jpg", ".gif", ".png":
		return true
	default:
		return false
	}
}

var ignoredFiles = map[string]bool{
	"index.html":   true,
	"index.json":   true,
	"Taskfile.yml": true,
}

func isIgnoredName(filename string) bool {
	if strings.HasPrefix(filename, ".") {
		return true
	}
	return ignoredFiles[filename]
}

func templateData(config *flags) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	data["title"] = config.title

	dirs := []fileRecord{}
	files := []fileRecord{}
	images := []fileRecord{}

	records, err := os.ReadDir(config.input)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		filename := record.Name()

		if isIgnoredName(filename) {
			continue
		}

		if record.IsDir() {
			dirs = append(dirs, newFileRecord(filename))
			continue
		}

		if isImageExtension(filename) {
			images = append(images, newFileRecord(filename))
			continue
		}

		files = append(files, newFileRecord(filename))
	}

	data["images"] = images
	data["files"] = files
	data["dirs"] = dirs

	return data, nil
}

func start(ctx context.Context, config *flags) error {
	if err := config.Validate(); err != nil {
		return err
	}

	templateData, err := templateData(config)
	if err != nil {
		return fmt.Errorf("error decoding template data: %w", err)
	}

	templateOutput, err := RenderTemplate(config.template, templateData)
	if err := WriteFile(path.Join(config.input, config.output), templateOutput, err); err != nil {
		return err
	}

	jsonOutput, err := json.MarshalIndent(templateData, "", "  ")
	if err := WriteFile(path.Join(config.input, config.outputJSON), jsonOutput, err); err != nil {
		return err
	}

	return nil
}
