package model

import (
	"errors"
)

// ErrUnimplemented is a helper for commands in development
var ErrUnimplemented = errors.New("unimplemented")

// ErrUnexported exists to handle/skip unexported types.
var ErrUnexported = errors.New("expected: unexported type")
