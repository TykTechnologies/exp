package semgrep

import (
	"errors"
	"html/template"
	"io"
)

var (
	templates = make(map[string]*template.Template)
)

func loadTemplateFromCache(name string) (*template.Template, error) {
	t, ok := templates[name]
	if !ok {
		return nil, errors.New("no such template: " + name)
	}
	return t, nil
}

func setTemplateToCache(name string, tpl *template.Template) (*template.Template, error) {
	templates[name] = tpl
	return tpl, nil
}

func loadTemplateFromEmbedFS(name string) (*template.Template, error) {
	t, err := template.ParseFS(files, name)
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
		t, err := loader(name)
		if err == nil {
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
