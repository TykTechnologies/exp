package example

import (
	"errors"
	"os"
)

// Status is an enum type.
type Status int

const (
	Active Status = iota
	Inactive
	Suspended
)

// Role is an enum type for user roles.
type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
	Guest Role = "guest"
)

// Type declaration doc, File is an alias of os.File.
type File = os.File

// Declaration doc.
type (
	// KeyRequest doc.
	KeyRequest struct {
		// SessionID doc
		SessionID string `json:"session_id"` // SessionID comment
	}

	// KeyResponse doc.
	KeyResponse struct {
		id string `json:"-"`

		Status  string `json:"status"`
		Message string `json:"message"`
	}

	// Key embeds KeyRequest and KeyResponse.
	Key struct {
		KeyRequest  `json:""`
		KeyResponse `json:""`
	}

	NamedRequests map[string]*KeyRequest
)

const foo = "bar"

var bar = "baz"

type keyRequest struct{}

func (t *KeyRequest) Validate() error {
	if t.SessionID == "" {
		return errors.New("invalid KeyRequest, empty session")
	}
}

func Validate() error {
	return nil
}

func validate() error {
	return nil
}
