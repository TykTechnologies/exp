package files

import (
	"os"
	"testing"
)

func TestFile_Flush(t *testing.T) {
	// Create a temporary file for testing
	tempFile := "test_file.go"
	defer os.Remove(tempFile)

	// Create a File struct with test data
	file := &File{
		Filename: tempFile,
		Package:  "main",
		Imports:  []string{`"fmt"`, `"os"`},
		Types: []string{
			"type MyType struct {\n\tName string\n}",
			"type AnotherType struct {\n\tAge int\n}",
		},
	}

	// Flush the file content to disk
	err := file.Flush()
	if err != nil {
		t.Fatalf("Flush() error = %v", err)
	}

	// Read the file content to verify
	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("failed to read the test file: %v", err)
	}

	expectedContent := `package main

import (
	"fmt"
	"os"
)

type MyType struct {
	Name string
}

type AnotherType struct {
	Age int
}

`

	if string(content) != expectedContent {
		t.Errorf("file content mismatch.\nExpected:\n%s\nGot:\n%s", expectedContent, string(content))
	}
}
