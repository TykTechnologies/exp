package render

import (
	"embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/*.md
var fs embed.FS

func listEmbeddedFiles(embedFS embed.FS, match string) []string {
	var matchedFiles []string
	files, _ := embedFS.ReadDir("testdata")
	for _, file := range files {
		if strings.Contains(file.Name(), match) {
			matchedFiles = append(matchedFiles, file.Name())
		}
	}
	return matchedFiles
}

func TestRender(t *testing.T) {
	inputFiles := listEmbeddedFiles(fs, "-input.md")

	for _, inputFile := range inputFiles {
		filename := "testdata/" + inputFile
		t.Run(filename, func(t *testing.T) {
			in, err := fs.ReadFile(filename)
			assert.NoError(t, err)

			want, err := fs.ReadFile(strings.ReplaceAll(filename, "-input.", "-want."))
			assert.NoError(t, err)

			got := Render(in)

			if false {
				fmt.Printf("source ======\n%s\n", string(in))
				fmt.Printf("want ========\n%s\n", string(want))
				fmt.Printf("got =========\n%s\n", string(got))
			}

			assert.NotEqual(t, string(in), string(want), "Input/wanted output should differ")
			assert.NotEqual(t, string(in), string(got), "Input/actual output should differ")
			assert.Equal(t, string(want), string(got), "Expected case to match wanted markdown")
		})
	}
}
