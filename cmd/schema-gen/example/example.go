package example

import (
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
)
