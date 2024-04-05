package coverfunc

import (
	"fmt"
)

// CoverageInfo represents information about coverage for a specific function.
type CoverageInfo struct {
	Filename string
	Line     int
	Function string
	Percent  float64
}

// PackageInfo represents information about coverage for a package.
type PackageInfo struct {
	Package   string
	Functions int
	Coverage  float64
}

// String returns a string representation of a PackageInfo.
func (p PackageInfo) String() string {
	return fmt.Sprintf("%s, symbols %d, coverage %.2f%%", p.Package, p.Functions, p.Coverage)
}

// FileInfo represents information about coverage for a file.
type FileInfo struct {
	Filename  string
	Functions int
	Coverage  float64
}

// String returns a string representation of a FileInfo.
func (f FileInfo) String() string {
	return fmt.Sprintf("%s, symbols %d, coverage %.2f%%", f.Filename, f.Functions, f.Coverage)
}
