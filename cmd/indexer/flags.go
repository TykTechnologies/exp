package main

import (
	"flag"
	"path/filepath"
)

type flags struct {
	input    string
	title    string
	template string

	output     string
	outputJSON string
}

func (f *flags) Bind() {
	flag.StringVar(&f.title, "t", "", "Title of index page to render")
	flag.StringVar(&f.input, "i", ".", "Input folder (default: current folder)")
	flag.StringVar(&f.output, "o", "index.html", "Output filename or absolute filepath")
	flag.StringVar(&f.outputJSON, "json", "index.json", "Output json directory listing")
	flag.StringVar(&f.template, "template", "index.tpl", "Template to render (index.tpl is bundled)")
}

// Validate will evaluate *flags and modify them, return an error if any occurs.
// From then on, the *flags object is ready to use. This lets us have computed
// fields, and upgrade deprecated flags to newer fields...
func (f *flags) Validate() error {
	if f.title == "" {
		inputAbs, err := filepath.Abs(f.input)
		if err != nil {
			return err
		}

		inputBase := filepath.Base(inputAbs)

		f.title = "Index of " + inputBase + "/"
	}
	return nil
}
