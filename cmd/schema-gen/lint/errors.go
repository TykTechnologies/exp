package lint

import (
	"strings"
)

// FieldDocError holds a list of errors.
type FieldDocError struct {
	errs []string
}

// Error implements the error interface.
func (err *FieldDocError) Error() string {
	return strings.Join(err.errs, "\n")
}

// Append appends an error message to the error list.
func (err *FieldDocError) Append(errMsg string) {
	err.errs = append(err.errs, errMsg)
}

// Empty returns true if there are no errors in the list.
func (err *FieldDocError) Empty() bool {
	return len(err.errs) == 0
}
