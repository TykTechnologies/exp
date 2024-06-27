package semgrep

import (
	"flag"
)

type flags struct {
	input    string
	output   string
	format   string
	template string
}

func (f *flags) Bind() {
	flag.StringVar(&f.input, "i", "", "Input file: -i input.json; default uses standard input")
	flag.StringVar(&f.output, "o", "", "Output file: -o output.html; default uses standard output")
	flag.StringVar(&f.format, "f", "", "Shorthand to set the output template to `report-<format>.tpl`")
	flag.StringVar(&f.template, "template", "report.tpl", "Template to render (report.tpl is bundled)")
}

// Validate will evaluate *flags and modify them, return an error if any occurs.
// From then on, the *flags object is ready to use. This lets us have computed
// fields, and upgrade deprecated flags to newer fields...
func (f *flags) Validate() error {
	if f.format != "" {
		f.template = "report-" + f.format + ".tpl"
	}
	return nil
}
