package golangcilint

import "fmt"

func Convert(root *Root) *Summary {
	summary := &Summary{}
	fileMap := make(map[string]*File) // Maps file name to its File pointer

	for _, issue := range root.Issues {
		file, exists := fileMap[issue.Pos.Filename]
		if !exists {
			file = &File{Name: issue.Pos.Filename}
			summary.Files = append(summary.Files, file)
			fileMap[issue.Pos.Filename] = file
		}

		// Find or create the error in the file
		var found *Error
		for i, err := range file.Errors {
			if err.FromLinter == issue.FromLinter {
				found = file.Errors[i]
				break
			}
		}

		textWithLine := fmt.Sprintf("L%d: %s", issue.Pos.Line, issue.Text)

		if found != nil {
			found.Text = append(found.Text, textWithLine)
			found.Count++
		} else {
			file.Errors = append(file.Errors, &Error{
				FromLinter: issue.FromLinter,
				Text:       []string{textWithLine},
				Count:      1,
			})
		}
	}

	return summary
}
