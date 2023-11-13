package list

import (
	"strings"

	"golang.org/x/mod/semver"
)

// SanitizeList clears out trailing major versions from list
func SanitizeList(list []string) []string {
	var major = ".0.0"
	var minor = ".0"

	// Trim repeating major versions from end of list (4.0.0, 5.0.0)
	for len(list) > 1 {
		n1, n2 := list[len(list)-1], list[len(list)-2]

		if strings.HasSuffix(n1, major) && strings.HasSuffix(n2, major) {
			list = list[0 : len(list)-1]
			continue
		}

		break
	}

	// Trim repeating minor versions from end of list (4.2.0, 4.3.0)
	for len(list) > 1 {
		n1, n2 := list[len(list)-1], list[len(list)-2]

		if semver.Major(n1) != semver.Major(n2) {
			break
		}

		if strings.HasSuffix(n1, minor) && strings.HasSuffix(n2, minor) {
			list = list[0 : len(list)-1]
			continue
		}

		break
	}

	return list
}

// SanitizeSet clears out values in second parameter based on the first.
func SanitizeSet(set1 []string, set2 []string) []string {
	set1 = SanitizeList(set1)
	set2 = SanitizeList(set2)

	for k, removed := range set2 {
		remove := false

		for _, added := range set1 {
			// if added > removed
			remove = remove || semver.Compare(added, removed) > 0
			if remove {
				break
			}
		}

		if remove {
			set2[k] = ""
		}
	}

	result := []string{}
	for _, v := range set2 {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}
