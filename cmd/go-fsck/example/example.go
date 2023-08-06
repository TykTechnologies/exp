package example

import (
	_ "net/http"
)

// Comment outer
var (
	// Comment inner
	exampleGroup1 = "Hello"
	exampleGroup2 = "There"
)

// Comment outer
var exampleGroup3 = "Sir"

// File represents a filename
type File string

// Body represends a decoded body
type Body struct {
	Name string
}

// Const comment
const E_WARNING = "warning" // const line comment

// Global func comment
func GlobalFunc() error {
	// holds the error
	var err error // the err var

	// inline comment
	err = nil

	return err
}
