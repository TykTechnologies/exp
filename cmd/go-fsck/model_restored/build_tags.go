package model

import (
	"regexp"
	"strings"
)

func BuildTags(src []byte) []string {

	re := regexp.MustCompile(`(?m)^\s*//\s*\+build\s+(.*)$`)

	var buildTags []string

	matches := re.FindAllStringSubmatch(string(src), -1)
	for _, match := range matches {
		buildTag := strings.TrimSpace(match[1])
		buildTags = append(buildTags, buildTag)
	}

	if len(buildTags) == 1 && strings.HasPrefix(buildTags[0], "!") {
		return nil
	}

	return buildTags
}
