package limiters

import (
	"errors"
	"fmt"
)

var (
	// ErrLimitExhausted is returned by the Limiter in case the number of requests overflows the capacity of a Limiter.
	ErrLimitExhausted = errors.New("requests limit exhausted")

	// ErrRaceCondition is returned when there is a race condition while saving a state of a rate limiter.
	ErrRaceCondition = errors.New("race condition detected")
)

func errorsWrap(err error, message string) error {
	if err != nil {
		return fmt.Errorf(message+": %w", err)
	}
	return nil
}
