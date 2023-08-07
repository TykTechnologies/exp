package restore

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsConflicting(t *testing.T) {
	testcases := []struct {
		in   []string
		want bool
	}{
		{
			in:   []string{`"net/tcp"`},
			want: false,
		},
		{
			in:   []string{`"text/template"`},
			want: true,
		},
		{
			in:   []string{`alias "text/template"`},
			want: true,
		},
	}

	for _, tc := range testcases {
		got := isConflicting(tc.in)
		assert.Equal(t, tc.want, got, "in: "+strings.Join(tc.in, ", "))
	}
}
