package lint

import (
	"strings"
)

// LintError holds a list of errors.
type LintError struct {
	errs []string
}

// NewLintError constructs a *LintError.
// The internal errs field does not need initialization.
func NewLintError() *LintError {
	return &LintError{}
}

// Error implements the error interface.
func (err *LintError) Error() string {
	return strings.Join(err.errs, "\n")
}

// Combine adds *LintErrors into receiver.
func (err *LintError) Combine(add *LintError) {
	if add.Empty() {
		return
	}

	err.Append(add.errs...)
}

// Append appends an error message to the error list.
func (err *LintError) Append(errMsg ...string) {
	for _, errText := range errMsg {
		if errText != "" {
			err.errs = append(err.errs, errText)
		}
	}
}

// Empty returns true if there are no errors in the list.
func (err *LintError) Empty() bool {
	return err == nil || len(err.errs) == 0
}
