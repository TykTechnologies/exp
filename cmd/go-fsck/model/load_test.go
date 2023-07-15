package model_test

import (
	"testing"

	"github.com/kortschak/utter"
	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

func TestLoad(t *testing.T) {
	utter.Config.IgnoreUnexported = true
	utter.Config.OmitZero = true
	utter.Config.ElideType = true

	defs, err := model.Load(".")
	assert.NoError(t, err)
	assert.NotNil(t, defs)

	utter.Dump(defs)
}
