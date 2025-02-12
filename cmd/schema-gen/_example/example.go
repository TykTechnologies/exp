package example

import (
	"errors"
	"os"
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
