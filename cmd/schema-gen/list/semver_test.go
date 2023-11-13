package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeList(t *testing.T) {
	type testcase struct {
		input    []string
		expected []string
	}

	testcases := []testcase{
		testcase{
			input: []string{
				"v3.0.10",
				"v3.1.1",
				"v3.2.0",
				"v4.0.0",
				"v5.0.0",
			},
			expected: []string{
				"v3.0.10",
				"v3.1.1",
				"v3.2.0",
				"v4.0.0",
			},
		},
		testcase{
			input: []string{
				"v4.0.14",
				"v5.0.3",
				"v5.1.0",
				"v5.2.0",
			},
			expected: []string{
				"v4.0.14",
				"v5.0.3",
				"v5.1.0",
			},
		},
		testcase{
			input: []string{
				"v4.1.0",
				"v4.2.0",
				"v4.3.0",
			},
			expected: []string{
				"v4.1.0",
			},
		},
	}

	for _, tc := range testcases {
		got := SanitizeList(tc.input)
		want := tc.expected

		assert.Equal(t, want, got)
	}
}

func TestSanitizeSet(t *testing.T) {
	type testcase struct {
		added    []string
		removed  []string
		expected []string
	}

	testcases := []testcase{
		testcase{
			added: []string{
				"v3.0.10",
				"v3.1.1",
				"v3.2.0",
				"v4.0.0",
				"v5.0.0",
			},
			removed: []string{
				"v3.1.0",
			},
			expected: []string{},
		},
		testcase{
			added: []string{
				"v3.0.10",
				"v3.2.0",
				"v4.0.0",
				"v5.0.0",
			},
			removed: []string{
				"v3.1.0",
			},
			expected: []string{},
		},
		testcase{
			added: []string{
				"v4.0.14",
				"v5.0.3",
				"v5.1.0",
				"v5.2.0",
			},
			removed: []string{
				"v4.1.0",
				"v4.2.0",
				"v4.3.0",
			},
			expected: []string{},
		},
		testcase{
			added: []string{
				"v3.0.0",
			},
			removed: []string{
				"v5.0.0",
			},
			expected: []string{
				"v5.0.0",
			},
		},
	}

	for _, tc := range testcases {
		got := SanitizeSet(tc.added, tc.removed)
		want := tc.expected

		assert.Equal(t, want, got)
	}
}
