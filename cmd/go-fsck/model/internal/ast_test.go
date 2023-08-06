package internal_test

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model/internal"
)

const src = `package example

// Global func comment
func GlobalFunc() error {
	// holds the error
	var err error	// the err var

	// inline comment
	err = nil

	return err
}`

func TestPrint(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	assert.NoError(t, err)

	var out strings.Builder
	assert.NoError(t, internal.PrintSource(internal.CommentedNode(f, f), fset, &out))

	assert.Equal(t, src, strings.TrimSpace(out.String()))
}
