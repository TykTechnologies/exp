package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var actionRegex = regexp.MustCompile(`uses:\s*([\w\/@.-]+)`)

// TraverseYAMLFiles traverses the .github/workflows directory and finds .yml files
func TraverseYAMLFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only process .yml files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yml") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// FindActionsInFile scans a .yml file and finds all `uses:` lines
func FindActionsInFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	actions := make(map[string]string)
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if match := actionRegex.FindStringSubmatch(line); match != nil {
			action := match[1]
			// Ignore local actions that start with './'
			if strings.HasPrefix(action, "./") {
				continue
			}
			actions[fmt.Sprintf("%s:%d", filePath, lineNumber)] = action
		}
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return actions, nil
}

// CompareActions detects mismatches between different versions of the same action
func CompareActions(actions map[string]string) {
	versions := make(map[string]map[string]string)

	for location, action := range actions {
		parts := strings.Split(action, "@")
		if len(parts) != 2 {
			fmt.Printf("Skipping malformed action in %s: %s\n", location, action)
			continue
		}
		name, version := parts[0], parts[1]
		if versions[name] == nil {
			versions[name] = make(map[string]string)
		}
		versions[name][version] = location
	}

	// Report mismatches
	for action, versionLocations := range versions {
		if len(versionLocations) > 1 {
			fmt.Printf("Action '%s' has multiple versions:\n", action)
			for version, location := range versionLocations {
				fmt.Printf("  - Version %s found at %s\n", version, location)
			}
		}
	}
}

func main() {
	// Specify the directory to search in
	root := ".github/workflows"

	// Step 1: Traverse the directory and find all .yml files
	files, err := TraverseYAMLFiles(root)
	if err != nil {
		fmt.Printf("Error traversing directory: %v\n", err)
		return
	}

	// Step 2: Scan each file for actions and track them
	allActions := make(map[string]string)
	for _, file := range files {
		actions, err := FindActionsInFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}
		for k, v := range actions {
			allActions[k] = v
		}
	}

	// Step 3: Compare actions and detect mismatches
	CompareActions(allActions)
}
