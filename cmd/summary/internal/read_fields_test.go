package internal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFields(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]string
	}{
		{
			name:     "Empty Input",
			input:    "",
			expected: [][]string{},
		},
		{
			name:     "Single Line",
			input:    "Hello World",
			expected: [][]string{{"Hello", "World"}},
		},
		{
			name:     "Multiple Lines",
			input:    "Hello World\nTesting\n123 456",
			expected: [][]string{{"Hello", "World"}, {"Testing"}, {"123", "456"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			result, err := ReadFields(reader)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
