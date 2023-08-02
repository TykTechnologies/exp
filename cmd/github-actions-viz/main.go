package main

// This collects individual github actions yamls, tries to parse them
// and render documentation with a mermaidjs state diagram output.

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/nektos/act/pkg/model"
	"github.com/spf13/pflag"
)

func main() {
	if err := start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func start(context.Context) error {
	var inputPath = "."
	pflag.StringVarP(&inputPath, "input-path", "i", inputPath, "input path")
	pflag.Parse()

	files, err := filepath.Glob(path.Join(inputPath, "*.yml"))
	if err != nil {
		return err
	}

	for _, filename := range files {
		if path.Base(filename) == "Taskfile.yml" {
			continue
		}

		m, err := open(filename)
		if err != nil {
			return err
		}
		fmt.Println(render(m, filename))
	}

	return nil
}

func open(filename string) (*model.Workflow, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Can't open %s, %w", filename, err)
	}
	defer f.Close()

	m, err := model.ReadWorkflow(f)
	if err != nil {
		return nil, fmt.Errorf("Can't parse workflow %s, %w", filename, err)
	}

	return m, nil
}

const header = `stateDiagram-v2
    workflow : %s - %s
    state workflow {
`

func render(m *model.Workflow, filename string) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, header, filename, m.Name)

	type wrap struct {
		K string
		V *model.Job
	}

	// map job step onto next jobs if any
	outputs := map[string][]string{}

	rootJobs := make([]wrap, 0, len(m.Jobs))
	for key, job := range m.Jobs {
		needs := job.Needs()
		rootJobs = append(rootJobs, wrap{key, job})
		for _, need := range needs {
			outputs[need] = append(outputs[need], key)
		}
	}

	workflows := []string{}
	for _, v := range rootJobs {
		workflows = append(workflows, renderJob(v.K, v.V, outputs))
	}

	io.WriteString(buf, strings.Join(workflows, "\n"))
	io.WriteString(buf, "    }\n\n")

	return buf.String()
}

func renderJob(key string, job *model.Job, outputs map[string][]string) string {
	result := []string{
		fmt.Sprintf("%s: %s", key, job.Name),
		fmt.Sprintf("state %s {", key),
	}
	type wrap struct {
		from, to, name string
	}

	from := "[*]"
	steps := []wrap{}
	for stepIndex, step := range job.Steps {
		if step.Name != "" {
			to := fmt.Sprintf("step%d%s", stepIndex, key)
			steps = append(steps, wrap{
				from: from,
				to:   to,
				name: step.Name,
			})
			from = to
		}
	}

	var to string
	for _, step := range steps {
		result = append(result, fmt.Sprintf("    %s --> %s", step.from, step.to))
		result = append(result, fmt.Sprintf("    %s : %s", step.to, step.name))
		to = step.to
	}

	if val, ok := outputs[key]; ok {
		for _, output := range val {
			result = append(result, fmt.Sprintf("    %s --> %s", to, output))
		}
	}

	result = append(result, "}")

	buf := new(bytes.Buffer)
	indent := "        "
	for _, line := range result {
		io.WriteString(buf, indent+line+"\n")
	}
	return buf.String()
}
