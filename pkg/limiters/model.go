package limiters

import (
	"time"
)

// Clock encapsulates a system Clock.
// Used
type Clock interface {
	// Now returns the current system time.
	Now() time.Time
}

// Logger wraps the Log method for logging.
type Logger interface {
	// Log logs the given arguments.
	Log(v ...interface{})
}
