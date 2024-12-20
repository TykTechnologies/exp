package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.True(t, true)
	assert.NoError(t, GlobalFunc())
}
