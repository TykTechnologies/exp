package golangcilint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Updated test function to match the new structure of Error with Text as []string
func TestConvert(t *testing.T) {
	root := &Root{
		Issues: []Issue{
			{
				FromLinter: "errcheck",
				Text:       "Error return value of `f.Close` is not checked",
				Pos: Position{
					Filename: "cli/importer/importer.go",
					Line:     10,
				},
			},
			{
				FromLinter: "errcheck",
				Text:       "Error return value of `os.Open` is not checked",
				Pos: Position{
					Filename: "cli/importer/importer.go",
					Line:     20,
				},
			},
			{
				FromLinter: "staticcheck",
				Text:       "S1000: should use for range instead of for { select {} }",
				Pos: Position{
					Filename: "cli/exporter/exporter.go",
					Line:     15,
				},
			},
		},
	}

	expected := &Summary{
		Files: []*File{
			{
				Name: "cli/importer/importer.go",
				Errors: []*Error{
					{
						FromLinter: "errcheck",
						Text:       []string{"L10: Error return value of `f.Close` is not checked", "L20: Error return value of `os.Open` is not checked"},
						Count:      2,
					},
				},
			},
			{
				Name: "cli/exporter/exporter.go",
				Errors: []*Error{
					{
						FromLinter: "staticcheck",
						Text:       []string{"L15: S1000: should use for range instead of for { select {} }"},
						Count:      1,
					},
				},
			},
		},
	}

	result := Convert(root)

	assert.Equal(t, expected, result, "The converted summary does not match the expected output.")
}
