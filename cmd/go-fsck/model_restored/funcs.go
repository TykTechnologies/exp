package model

import (
	"go/ast"
	"regexp"
)

func appendIfNotExists(slice []string, element string) []string {
	for _, s := range slice {
		if s == element {
			return slice
		}
	}
	return append(slice, element)
}

func getBuildTags(file *ast.File) []string {

	re := regexp.MustCompile(`^\s*//\s*\+build\s+(.*)$`)

	var buildTags []string

	if file.Doc != nil {
		for _, comment := range file.Doc.List {
			match := re.FindStringSubmatch(comment.Text)
			if len(match) > 1 {
				buildTags = append(buildTags, match[1])
			}
		}
	}

	return buildTags
}
