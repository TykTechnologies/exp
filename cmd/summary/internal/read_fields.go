package internal

import (
	"bufio"
	"io"
	"strings"
)

func ReadFields(reader io.Reader) ([][]string, error) {
	result := make([][]string, 0, 100)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) > 0 {
			result = append(result, fields)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
