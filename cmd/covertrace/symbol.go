package main

import (
	"bufio"
	"os"
	"strings"
)

func getSymbolAndCoverage(filename string, startLine, endLine, numStmts, numCov int) (string, string, float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", "", 0, err
	}
	defer file.Close()

	var symbol, receiver string
	var coverage float64
	var found bool

	scanner := bufio.NewScanner(file)
	lineNum := 0
	funcLine := ""

	// Read the file and find the line number where the function starts
	for scanner.Scan() {
		lineNum++
		if lineNum == startLine {
			funcLine = scanner.Text()
			break
		}
	}

	// Rewind until a line with a "func" prefix is found
	for lineNum > 0 && !strings.HasPrefix(strings.TrimSpace(funcLine), "func ") {
		lineNum--
		file.Seek(0, 0) // rewind to the beginning
		scanner = bufio.NewScanner(file)
		for i := 0; i < lineNum; i++ {
			scanner.Scan()
		}
		funcLine = scanner.Text()
	}

	// Extract the receiver and function name from the line
	funcLine = strings.TrimSpace(funcLine)
	if strings.HasPrefix(funcLine, "func ") {
		funcLine = strings.TrimPrefix(funcLine, "func ")
		if strings.HasPrefix(funcLine, "(") {
			endIdx := strings.Index(funcLine, ")")
			receiver = funcLine[1:endIdx]
			parts := strings.SplitN(receiver, " ", 2)
			if len(parts) == 2 {
				receiver = parts[1]
			}
			receiver = strings.Trim(receiver, "*")
			funcLine = funcLine[endIdx+1:]
		}
		funcNameEndIdx := strings.IndexAny(funcLine, "(")
		if funcNameEndIdx != -1 {
			symbol = funcLine[:funcNameEndIdx]
		} else {
			symbol = funcLine
		}
		symbol = strings.TrimSpace(symbol)
		found = true
	}

	// Calculate coverage
	if numStmts > 0 {
		coverage = (float64(numCov) / float64(numStmts)) * 100
	}

	if !found {
		symbol = "Unknown"
	}

	return symbol, receiver, coverage, scanner.Err()
}
