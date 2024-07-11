package files

import (
	"bytes"
	"fmt"
	"os"
)

// File struct represents a Go source file with its components.
type File struct {
	Filename string
	Package  string
	Imports  []string
	Types    []string
}

// Flush writes the File's content to the specified Filename.
func (f *File) Flush() error {
	// Skip writing if there are no types to write
	if len(f.Types) == 0 {
		return nil
	}

	var buffer bytes.Buffer

	// Write package declaration
	buffer.WriteString(fmt.Sprintf("package %s\n\n", f.Package))

	// Write import statements
	if len(f.Imports) > 0 {
		buffer.WriteString("import (\n")
		for _, imp := range f.Imports {
			buffer.WriteString(fmt.Sprintf("\t%s\n", imp))
		}
		buffer.WriteString(")\n\n")
	}

	// Write types
	for _, typ := range f.Types {
		buffer.WriteString(fmt.Sprintf("%s\n\n", typ))
	}

	// Write buffer contents to file
	return os.WriteFile(f.Filename, buffer.Bytes(), 0644)
}
