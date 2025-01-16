package main

import (
	"fmt"

	"github.com/ankitpokhrel/jira-cli/pkg/jira"
)

// Global variable for allowed ticket states
var allowedStates = []string{
	"In Dev",
	"In Code Review",
	"Ready for Testing",
	"In Test",
	"In Progress",
	"In Review",
}

// allowedStatesMap is a map for efficient lookup of valid states
var allowedStatesMap = func() map[string]bool {
	m := make(map[string]bool)
	for _, state := range allowedStates {
		m[state] = true
	}
	return m
}()

// validateTicketState checks if the Jira ticket is in one of the allowed states.
func validateTicketState(i *jira.Issue) error {
	// Get the issue's current status
	status := i.Fields.Status.Name

	// Check if the status is allowed
	if !allowedStatesMap[status] {
		return fmt.Errorf(
			"ticket '%s' is in an invalid state: '%s'. Allowed states are: %v",
			i.Key, status, allowedStates,
		)
	}

	// Validation succeeded
	return nil
}
