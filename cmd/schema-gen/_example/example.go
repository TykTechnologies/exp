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
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	// Key embeds KeyRequest and KeyResponse.
	Key struct {
		KeyRequest  `json:""`
		KeyResponse `json:""`
	}
)

func (t *KeyRequest) Validate() error {
	if t.SessionID == "" {
		return errors.New("invalid KeyRequest, empty session")
	}
}
