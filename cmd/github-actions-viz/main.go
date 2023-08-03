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
	"sort"
	"strings"

	"github.com/nektos/act/pkg/model"
	"github.com/spf13/pflag"
)

func main() {
	if err := start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

type options struct {
	inputPath string
	writeOut  bool
	format    string
}

func start(context.Context) error {
	config := options{
		inputPath: ".",
		format:    "md",
	}
	pflag.StringVarP(&config.inputPath, "input-path", "i", config.inputPath, "input path")
	pflag.BoolVarP(&config.writeOut, "write-out", "w", config.writeOut, "write out as files")
	pflag.StringVar(&config.format, "format", config.format, "format (md, mermaid)")
	pflag.Parse()

	yamls, err := filepath.Glob(path.Join(config.inputPath, "*.yaml"))
	if err != nil {
		return err
	}
	ymls, err := filepath.Glob(path.Join(config.inputPath, "*.yml"))
	if err != nil {
		return err
	}

	files := []string{}
	files = append(files, yamls...)
	files = append(files, ymls...)

	for _, filename := range files {
		if path.Base(filename) == "Taskfile.yml" {
			continue
		}

		m, err := open(filename)
		if err != nil {
			return err
		}
		output := render(config, m, filename)
		if output != "" {
			fmt.Println(output)
		}
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

func render(config options, m *model.Workflow, filename string) string {
	var none string

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
	sort.SliceStable(rootJobs, func(i, j int) bool {
		if rootJobs[i].K < rootJobs[j].K {
			return true
		}
		return false
	})

	workflows := []string{}
	for _, v := range rootJobs {
		workflows = append(workflows, renderJob(v.K, v.V, outputs))
	}

	io.WriteString(buf, strings.Join(workflows, "\n"))
	io.WriteString(buf, "    }\n\n")

	if config.format == "md" {
		markdown := []string{}
		markdown = append(markdown, "# "+m.Name)
		markdown = append(markdown, fmt.Sprintf("```mermaid\n%s\n```", strings.TrimSpace(buf.String())))

		if config.writeOut {
			output := filename + ".md"
			fmt.Println(output)
			body := []byte(strings.Join(markdown, "\n\n") + "\n")
			if err := os.WriteFile(output, body, 0644); err != nil {
				panic(err)
			}
			return none
		}
	}

	if config.writeOut {
		output := filename + ".mermaid"
		fmt.Println(output)
		body := buf.Bytes()
		if err := os.WriteFile(output, body, 0644); err != nil {
			panic(err)
		}
		return none
	}

	return buf.String()
}

func isset(strs ...string) string {
	for _, str := range strs {
		if str != "" {
			return str
		}
	}
	return ""
}

func renderJob(key string, job *model.Job, outputs map[string][]string) string {
	indent := "        "
	name := isset(job.Name, key)
	if job.Name == "" && len(job.Steps) == 1 {
		name = isset(job.Name, job.Steps[0].Name, key)
	}

	if len(job.Steps) == 0 {
		result := []string{
			indent + fmt.Sprintf("%s: %s", key, name),
			indent + fmt.Sprintf("state %s {", key),
			indent + "    [*] --> Finish",
			indent + "}",
		}
		return strings.Join(result, "\n") + "\n"
	}

	result := []string{
		fmt.Sprintf("%s: %s", key, name),
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
		sort.Strings(val)
		for _, output := range val {
			result = append(result, fmt.Sprintf("    %s --> %s", to, output))
		}
	}

	result = append(result, "}")

	buf := new(bytes.Buffer)
	for _, line := range result {
		io.WriteString(buf, indent+line+"\n")
	}
	return buf.String()
}
