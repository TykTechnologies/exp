package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

type Targets struct {
	Version string `json:"-"`

	Defaults map[string]string `json:"defaults,omitempty"`
	Targets  []*Target         `json:"targets"`
}

func NewTargets() *Targets {
	return &Targets{
		Defaults: make(map[string]string),
	}
}

type Target struct {
	Defaults      map[string]string `json:"defaults,omitempty"`
	Branch        []string          `json:"branch"`
	RequiredTests []string          `json:"required_tests" yaml:"required_tests"`
}

const tpl = `[
{{ range $target := .Targets }}
{{ range $branch := .Branch }}
    {
       branch   = {{ $branch }}
       required_tests = {{ $target.RequiredTests | json }}
{{ range $k, $v := $target.Defaults }}
       {{ $k }} = {{ $v }}
{{ end }}
    },
{{ end }}{{ end }}
]`

func start() error {
	data, err := os.ReadFile("main.yml")
	if err != nil {
		return nil
	}

	t := NewTargets()
	yaml.Unmarshal(data, t)

	for k, v := range t.Defaults {
		for _, target := range t.Targets {
			// Overwrite defaults if none
			if target.Defaults == nil {
				target.Defaults = t.Defaults
				continue
			}

			// Apply root defaults to empty targets
			if val := target.Defaults[k]; val == "" {
				target.Defaults[k] = v
			}
		}
	}

	funcs := template.FuncMap{
		"json": func(in interface{}) (string, error) {
			out, err := json.Marshal(in)
			return string(out), err
		},
	}

	render, err := template.New("").Funcs(funcs).Parse(tpl)
	if err != nil {
		return err
	}

	out := new(bytes.Buffer)
	if err := render.Execute(out, t); err != nil {
		return err
	}

	var line []byte
	for {
		line, err = out.ReadBytes('\n')
		// skip empty lines
		if trimmed := bytes.TrimSpace(line); len(trimmed) == 0 {
			continue
		}

		os.Stdout.Write(bytes.TrimRight(line, "\n"))
		os.Stdout.WriteString("\n")

		if err != nil {
			break
		}
	}
	if errors.Is(err, io.EOF) {
		err = nil
	}
	return err
}
