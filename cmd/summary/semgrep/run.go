package semgrep

import (
	"flag"
	"fmt"
)

func Run() error {
	config := &flags{}
	config.Bind()
	flag.Parse()

	if err := config.Validate(); err != nil {
		return err
	}

	in, err := readInput(config.input)
	if err != nil {
		return err
	}

	out, err := openOutput(config.output)
	if err != nil {
		return err
	}
	defer out.Close()

	templateData, err := templateData(in)
	if err != nil {
		return fmt.Errorf("error decoding template data: %w", err)
	}

	return RenderTemplate(out, config.template, templateData)
}
