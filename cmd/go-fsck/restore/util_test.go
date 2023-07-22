package restore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsConflicting(t *testing.T) {
	testcases := []struct {
		in   []string
		want bool
	}{
		{
			in:  []string{`"net/tcp"`},
			out: false,
		},
		{
			in:  []string{`"text/template"`},
			out: true,
		},
		{
			in:  []string{`alias "text/template"`},
			out: true,
		},
	}

	for _, tc := range testcases {
		got := isConflicting(tc.in)
		assert.Equal(t, tc.want, got)
	}
}
