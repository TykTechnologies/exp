package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// Render tries to scan the input markdown file and re-render it according to
// some common style rules that enforce spacing and bullet list consistency.
func Render(input []byte) []byte {
	options := blackfriday.WithExtensions(blackfriday.CommonExtensions)
	parser := blackfriday.New(options)

	// Add newline after we see an ending colon
	// Blackfriday already doesn't know how to parse this otherwise.
	pattern := regexp.MustCompile(`:\s*\n*`)
	input = pattern.ReplaceAll(input, []byte(":\n\n"))

	ast := parser.Parse(input)

	return render(ast)
}

func dump(in any) {
	b, _ := json.MarshalIndent(in, "", "  ")
	fmt.Println(string(b))
}

func render(parsedMarkdown *blackfriday.Node) []byte {
	var buf bytes.Buffer

	inList := false

	parsedMarkdown.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		switch node.Type {
		case blackfriday.Text:
			if entering {
				line := strings.TrimSpace(string(node.Literal))
				if inList && strings.Contains(line, "\n") {
					// shim to add spacing after list item
					line = strings.ReplaceAll(line, "\n", "\n\n")
				}
				buf.WriteString(line)
			}
		case blackfriday.Code:
			if entering {
				buf.WriteString("```\n")
			} else {
				buf.WriteString("\n```\n\n")
			}
		case blackfriday.CodeBlock:
			if entering {
				buf.WriteString("```\n" + strings.TrimSpace(string(node.Literal)) + "\n```\n\n")
			}
		case blackfriday.Paragraph:
			if !entering {
				if inList {
					buf.WriteString("\n")
				} else {
					buf.WriteString("\n\n")
				}
			}
		case blackfriday.Heading:
			if entering {
				level := strings.Repeat("#", node.HeadingData.Level)
				buf.WriteString(fmt.Sprintf("%s ", level))
			} else {
				buf.WriteString("\n\n")
			}
		case blackfriday.List:
			if entering {
				inList = true
			} else {
				inList = false
				buf.WriteString("\n") // Add newline after last list item
			}
		case blackfriday.Item:
			if entering {
				buf.WriteString("- ")
			}
		case blackfriday.Document:
			return blackfriday.GoToNext
		default:
			fmt.Printf("Unhandled blackfriday type: %v\n", node.Type)
		}
		return blackfriday.GoToNext
	})

	result := buf.Bytes()
	if len(result) > 1 {
		return result[:len(result)-1]
	}
	return result
}
