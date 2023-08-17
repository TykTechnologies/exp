package restore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsConflicting(t *testing.T) {
	testcases := []struct {
		in      []string
		wantErr bool
	}{
		{
			in:      []string{`"net/tcp"`},
			wantErr: false,
		},
		{
			in:      []string{`"text/template"`},
			wantErr: true,
		},
		{
			in:      []string{`alias "text/template"`},
			wantErr: true,
		},
	}

	for _, tc := range testcases {
		got := IsConflicting(tc.in)
		if tc.wantErr {
			assert.Error(t, got)
		} else {
			assert.NoError(t, got)
		}
		fmt.Println("err=", got)
	}
}
