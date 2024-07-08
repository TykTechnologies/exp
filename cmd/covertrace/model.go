package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type coverageInfo struct {
	File      string  `json:"file"`
	StartLine int     `json:"start_line"`
	EndLine   int     `json:"end_line"`
	NumStmts  int     `json:"num_stmts"`
	NumCov    int     `json:"num_cov"`
	Symbol    string  `json:"symbol"`
	Receiver  string  `json:"receiver,omitempty"`
	Coverage  float64 `json:"coverage"`
}

type yamlOutput struct {
	Types   []structType `yaml:"types"`
	Globals []globalFunc `yaml:"globals"`
	Totals  totalOutput  `yaml:"total"`
}

type structType struct {
	Struct string        `yaml:"struct"`
	Funcs  []funcDetails `yaml:"funcs"`
}

type funcDetails struct {
	Name     string `yaml:"name"`
	Coverage int    `yaml:"coverage"`
}

type globalFunc struct {
	Name     string `yaml:"name"`
	Coverage int    `yaml:"coverage"`
}

type totalOutput struct {
	Coverage struct {
		Funcs   int `yaml:"funcs"`
		Structs int `yaml:"structs"`
		Total   int `yaml:"total"`
	} `yaml:"coverage"`
}

func parseCoverageFile(filename string) ([]coverageInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var coverageData []coverageInfo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "mode:") {
			parts := strings.Fields(line)
			fileParts := strings.Split(parts[0], ":")
			file := fileParts[0]
			lines := strings.Split(fileParts[1], ",")
			startLine := parseLineNum(lines[0])
			endLine := parseLineNum(lines[1])
			numStmts := parseLineNum(parts[1])
			numCov := parseLineNum(parts[2])

			if numCov == 0 {
				continue
			}

			coverageData = append(coverageData, coverageInfo{
				File:      filepath.Base(file),
				StartLine: startLine,
				EndLine:   endLine,
				NumStmts:  numStmts,
				NumCov:    numCov,
			})
		}
	}
	return coverageData, scanner.Err()
}

func parseLineNum(line string) int {
	var num int
	fmt.Sscanf(line, "%d.%d", &num, new(int))
	return num
}
