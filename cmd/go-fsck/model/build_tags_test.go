package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TykTechnologies/exp/cmd/go-fsck/model"
)

func TestBuildTags(t *testing.T) {
	src := `
// +build debug
// +build linux

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`

	want := []string{"debug", "linux"}
	got := model.BuildTags([]byte(src))

	assert.Equal(t, want, got)
}
