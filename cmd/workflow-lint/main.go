package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

// Define a list of default latest versions for specific actions
var defaultUses = []string{
	"actions/setup-node@v4",
	"actions/cache@v4",
	"actions/setup-go@v5",
	"actions/download-artifact@v4",
	"actions/checkout@v4",
	"actions/setup-python@v5",
	"actions/upload-artifact@v4",
	// Add more default actions here if needed
}

var actionRegex = regexp.MustCompile(`uses:\s*([\w\/@.-]+)`)

// TraverseYAMLFiles traverses the .github/workflows directory and finds .yml files, skipping ignored files
func TraverseYAMLFiles(root string, ignoreFiles []string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories and check if file is ignored
		if info.IsDir() {
			return nil
		}

		// Check if the file should be ignored based on the `ignoreFiles` list
		for _, ignoreFile := range ignoreFiles {
			if filepath.Base(path) == ignoreFile {
				return nil // Skip this file
			}
		}

		// Only process .yml files
		if strings.HasSuffix(info.Name(), ".yml") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// ParseUses converts the given slice (either defaultUses or user-provided overrides) into a map
func ParseUses(uses []string) map[string]string {
	usesMap := make(map[string]string)
	for _, use := range uses {
		parts := strings.Split(use, "@")
		if len(parts) == 2 {
			usesMap[parts[0]] = parts[1] // Map action name to its latest version
		}
	}
	return usesMap
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
func CompareActions(actions map[string]string, latestDefaults map[string]string) map[string]string {
	versions := make(map[string]map[string]string)
	latestVersions := make(map[string]string)

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
		// Track the latest version for each action
		if latest, ok := latestVersions[name]; !ok || version > latest {
			latestVersions[name] = version
		}
	}

	// Override with the latestDefaults if a newer version exists
	for name, defaultVersion := range latestDefaults {
		latestVersions[name] = defaultVersion
	}

	return latestVersions
}

// UpdateActionVersion updates the action versions in the file to the latest version
func UpdateActionVersion(filePath string, actions map[string]string, latestVersions map[string]string) error {
	// Read the entire file
	input, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if match := actionRegex.FindStringSubmatch(line); match != nil {
			action := match[1]
			if strings.HasPrefix(action, "./") {
				continue // Ignore local actions
			}

			parts := strings.Split(action, "@")
			if len(parts) != 2 {
				continue
			}
			name, currentVersion := parts[0], parts[1]
			latestVersion, exists := latestVersions[name]

			// Update the action to the latest version if necessary
			if exists && currentVersion != latestVersion {
				lines[i] = strings.Replace(line, "@"+currentVersion, "@"+latestVersion, 1)
				fmt.Printf("Updating %s to %s in %s\n", action, name+"@"+latestVersion, filePath)
			}
		}
	}

	// Write the updated lines back to the file
	output := strings.Join(lines, "\n")
	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Define flags
	fixFlag := pflag.Bool("fix", false, "Automatically update actions to their latest versions")
	listFlag := pflag.Bool("list", false, "List the latest action versions used in the workflows")
	baseFlag := pflag.StringSlice("base", defaultUses, "List of action@version overrides")
	ignoreFlag := pflag.StringSlice("ignore", []string{}, "List of workflow files to ignore")

	pflag.Parse()

	// Specify the directory to search in
	root := ".github/workflows"

	// Step 1: Traverse the directory and find all .yml files, skipping ignored ones
	files, err := TraverseYAMLFiles(root, *ignoreFlag)
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

	// Step 3: Parse defaultUses and override if baseFlag is provided
	latestDefaults := ParseUses(*baseFlag)

	// Step 4: Compare actions and detect mismatches
	latestVersions := CompareActions(allActions, latestDefaults)

	// Step 5: Handle the -list flag to print the latest versions
	if *listFlag {
		fmt.Println("Latest action versions used:")
		for action, version := range latestVersions {
			fmt.Printf("%s@%s\n", action, version)
		}
		return
	}

	// Step 6: If --fix is passed, update the files to use the latest versions
	if *fixFlag {
		for _, file := range files {
			err := UpdateActionVersion(file, allActions, latestVersions)
			if err != nil {
				fmt.Printf("Error updating file %s: %v\n", file, err)
			}
		}
	}
}
