package limiters

import (
	"log"
)

// StdLogger implements the Logger interface.
type StdLogger struct{}

// NewStdLogger creates a new instance of StdLogger.
func NewStdLogger() *StdLogger {
	return &StdLogger{}
}

// Log delegates the logging to the std logger.
func (l *StdLogger) Log(v ...interface{}) {
	log.Println(v...)
}
