package main

import "strings"

func title(s string) string {
	return strings.Split(s, "\n")[0]
}

func description(s string) string {
	lines := strings.Split(s, "\n")
	lines = lines[1:]

	lastline := ""
	clean := []string{}

	for _, line := range lines {
		if strings.HasPrefix(line, "swagger:") {
			continue
		}
		if lastline == "" && line == "" {
			continue
		}
		lastline = line

		clean = append(clean, line)
	}

	out := strings.Join(clean, "\n")
	out = strings.TrimSpace(out)
	return out
}
