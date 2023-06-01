package lint

import (
	"strings"
)

// FieldDocError holds a list of errors.
type FieldDocError struct {
	errs []string
}

// NewFieldDocError constructs a *FieldDocError.
// The internal errs field does not need initialization.
func NewFieldDocError() *FieldDocError {
	return &FieldDocError{}
}

// Error implements the error interface.
func (err *FieldDocError) Error() string {
	return strings.Join(err.errs, "\n")
}

// Append appends an error message to the error list.
func (err *FieldDocError) Append(errMsg ...string) {
	if len(errMsg) > 0 {
		for _, errText := range errMsg {
			if errText != "" {
				err.errs = append(err.errs, errText)
			}
		}
	}
}

// Empty returns true if there are no errors in the list.
func (err *FieldDocError) Empty() bool {
	return len(err.errs) == 0
}
