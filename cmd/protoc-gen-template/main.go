package main

import (
	"flag"
	"path/filepath"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

type templateData struct {
	Plugin *protogen.Plugin
	File   *protogen.File
}

func generateFiles(templatePath string, plugin *protogen.Plugin, file *protogen.File) error {
	// ensure trailing slash for templatePath
	templatePath = strings.TrimRight(templatePath, "/") + "/"

	// list all available templates
	files, err := filepath.Glob(templatePath + "*.tmpl")
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	for _, tplFile := range files {
		filename := file.GeneratedFilenamePrefix + "_" + filepath.Base(strings.TrimSuffix(tplFile, ".tmpl"))

		g := plugin.NewGeneratedFile(filename, file.GoImportPath)

		err = tpl.ExecuteTemplate(g, filepath.Base(tplFile), &templateData{
			Plugin: plugin,
			File:   file,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var flags flag.FlagSet
	var (
		templatePath = flags.String("template_dir", "./templates", "")
	)

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}

			if err := generateFiles(*templatePath, plugin, f); err != nil {
				return err
			}
		}
		return nil
	})
}
