package limiters

import (
	"time"
)

// SystemClock implements the Clock interface by using the real system clock.
type SystemClock struct {
}

// NewSystemClock creates a new instance of SystemClock.
func NewSystemClock() *SystemClock {
	return &SystemClock{}
}

// Now returns the current system time.
func (c *SystemClock) Now() time.Time {
	return time.Now()
}
