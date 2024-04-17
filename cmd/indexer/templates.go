package main

import (
	"errors"
	"html/template"
	"io"
)

var (
	templates = make(map[string]*template.Template)

	errNotExists = errors.New("no such file in cache")
)

func loadTemplateFromCache(name string) (*template.Template, error) {
	if t, ok := templates[name]; ok {
		return t, nil
	}
	return nil, errNotExists
}

func setTemplateToCache(name string, tpl *template.Template) (*template.Template, error) {
	templates[name] = tpl
	return tpl, nil
}

func loadTemplateFromEmbedFS(name string) (*template.Template, error) {
	t, err := template.ParseFS(embeddedFiles, name)
	if err != nil {
		return nil, err
	}

	return setTemplateToCache(name, t)
}

func loadTemplateFromFilesystem(name string) (*template.Template, error) {
	t, err := template.ParseFiles(name)
	if err != nil {
		return nil, err
	}

	return setTemplateToCache(name, t)
}

func loadTemplate(name string) (*template.Template, error) {
	loaders := []func(string) (*template.Template, error){
		loadTemplateFromCache,
		loadTemplateFromEmbedFS,
		loadTemplateFromFilesystem,
	}

	for _, loader := range loaders {
		if t, err := loader(name); err == nil {
			return t, nil
		}
	}

	return nil, errors.New("no such template: " + name)
}

func RenderTemplate(w io.Writer, templateName string, data map[string]interface{}) error {
	t, err := loadTemplate(templateName)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}
