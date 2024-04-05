package coverfunc

import (
	"strconv"
	"strings"
)

// Parse parses the coverage data into CoverageInfo.
func Parse(data [][]string) []CoverageInfo {
	var coverageInfos []CoverageInfo

	for _, line := range data {
		filenameAndLine := strings.Split(line[0], ":")
		filename := filenameAndLine[0]
		if filename == "total" {
			continue
		}

		// Assuming that line number is always present and can be parsed
		lineNumber, _ := strconv.Atoi(filenameAndLine[1])

		info := CoverageInfo{
			Filename: filename,
			Line:     lineNumber,
			Function: line[1],
		}
		percent, _ := strconv.ParseFloat(strings.TrimSuffix(line[2], "%"), 64)
		info.Percent = percent
		coverageInfos = append(coverageInfos, info)
	}

	return coverageInfos
}
