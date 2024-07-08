package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/spf13/pflag"
)

func main() {
	var inputFile string
	var skipSummary bool
	var outputJSON bool

	pflag.StringVarP(&inputFile, "input", "i", "", "Input coverage file")
	pflag.BoolVar(&skipSummary, "skip-summary", false, "Skip summary output")
	pflag.BoolVar(&outputJSON, "json", false, "Output detailed data in JSON format")
	pflag.Parse()

	if inputFile == "" {
		fmt.Println("Usage: go run main.go -i <coverage_file> [--skip-summary] [--json]")
		return
	}

	coverageData, err := parseCoverageFile(inputFile)
	if err != nil {
		fmt.Println("Error parsing coverage file:", err)
		return
	}

	structMap := make(map[string]map[string]int)
	funcMap := make(map[string]int)

	for i := 0; i < len(coverageData); {
		cov := coverageData[i]
		symbol, receiver, coverage, err := getSymbolAndCoverage(cov.File, cov.StartLine, cov.EndLine, cov.NumStmts, cov.NumCov)
		if err != nil {
			fmt.Println("Error getting symbol or coverage:", err)
			return
		}

		if coverage > 0 {
			coverageData[i].Symbol = symbol
			coverageData[i].Receiver = receiver
			coverageData[i].Coverage = coverage

			if receiver != "" {
				if structMap[receiver] == nil {
					structMap[receiver] = make(map[string]int)
				}
				structMap[receiver][symbol] += cov.NumCov
				funcMap[fmt.Sprintf("%s.%s", receiver, symbol)] += cov.NumCov
			} else {
				funcMap[symbol] += cov.NumCov
			}
			i++
		} else {
			coverageData = append(coverageData[:i], coverageData[i+1:]...)
		}
	}

	if skipSummary {
		if outputJSON {
			printJSON(coverageData)
		} else {
			printRawYaml(coverageData)
		}
	} else {
		if outputJSON {
			printSummaryJSON(structMap, funcMap, coverageData)
		} else {
			summarizeYaml(structMap, funcMap)
		}
	}
}

func summarizeYaml(structMap map[string]map[string]int, funcMap map[string]int) {
	var yamlOut yamlOutput
	var structsCoverage int
	var globalsCoverage int

	for structName, funcs := range structMap {
		var functions []funcDetails
		for funcName, covStmts := range funcs {
			functions = append(functions, funcDetails{
				Name:     funcName,
				Coverage: covStmts,
			})
			structsCoverage += covStmts
		}
		yamlOut.Types = append(yamlOut.Types, structType{
			Struct: structName,
			Funcs:  functions,
		})
	}

	for funcName, covStmts := range funcMap {
		if !strings.Contains(funcName, ".") {
			yamlOut.Globals = append(yamlOut.Globals, globalFunc{
				Name:     funcName,
				Coverage: covStmts,
			})
			globalsCoverage += covStmts
		}
	}

	totalCoverage := structsCoverage + globalsCoverage
	yamlOut.Totals = totalOutput{
		Coverage: struct {
			Funcs   int `yaml:"funcs"`
			Structs int `yaml:"structs"`
			Total   int `yaml:"total"`
		}{
			Funcs:   globalsCoverage,
			Structs: structsCoverage,
			Total:   totalCoverage,
		},
	}

	yamlData, err := yaml.Marshal(yamlOut)
	if err != nil {
		fmt.Println("Error marshalling to YAML:", err)
		return
	}
	fmt.Println(string(yamlData))
}

func printJSON(coverageData []coverageInfo) {
	jsonOutput, err := json.MarshalIndent(coverageData, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonOutput))
}

func printRawYaml(coverageData []coverageInfo) {
	yamlOutput, err := yaml.Marshal(coverageData)
	if err != nil {
		fmt.Println("Error marshalling to YAML:", err)
		return
	}
	fmt.Println(string(yamlOutput))
}

func printSummaryJSON(structMap map[string]map[string]int, funcMap map[string]int, coverageData []coverageInfo) {
	var yamlOut yamlOutput
	var structsCoverage int
	var globalsCoverage int

	for structName, funcs := range structMap {
		var functions []funcDetails
		for funcName, covStmts := range funcs {
			functions = append(functions, funcDetails{
				Name:     funcName,
				Coverage: covStmts,
			})
			structsCoverage += covStmts
		}
		yamlOut.Types = append(yamlOut.Types, structType{
			Struct: structName,
			Funcs:  functions,
		})
	}

	for funcName, covStmts := range funcMap {
		if !strings.Contains(funcName, ".") {
			yamlOut.Globals = append(yamlOut.Globals, globalFunc{
				Name:     funcName,
				Coverage: covStmts,
			})
			globalsCoverage += covStmts
		}
	}

	totalCoverage := structsCoverage + globalsCoverage
	yamlOut.Totals = totalOutput{
		Coverage: struct {
			Funcs   int `yaml:"funcs"`
			Structs int `yaml:"structs"`
			Total   int `yaml:"total"`
		}{
			Funcs:   globalsCoverage,
			Structs: structsCoverage,
			Total:   totalCoverage,
		},
	}

	summaryJSON, err := json.MarshalIndent(yamlOut, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling summary to JSON:", err)
		return
	}
	fmt.Println(string(summaryJSON))
}
