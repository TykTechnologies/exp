package example

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
	var err error
	return err
}
