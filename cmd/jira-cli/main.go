package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ankitpokhrel/jira-cli/api"
	"github.com/ankitpokhrel/jira-cli/pkg/jira"
)

func main() {
	// Retrieve environment variables
	jiraURL := os.Getenv("JIRA_API_URL")
	jiraAPIToken := os.Getenv("JIRA_API_TOKEN")
	jiraAPIEmail := os.Getenv("JIRA_API_EMAIL")

	if jiraAPIToken == "" || jiraAPIEmail == "" {
		log.Fatalf("Environment variables JIRA_API_TOKEN and JIRA_API_EMAIL must be set")
	}

	// Initialize the API client
	authType := jira.AuthTypeBasic
	client := api.Client(jira.Config{
		Server:   jiraURL,
		AuthType: &authType,
		Login:    jiraAPIEmail,
		APIToken: jiraAPIToken,
	})

	// Example: Fetch a Jira issue by key
	issueKey := os.Args[1]
	issue, err := getJiraIssue(client, issueKey)
	if err != nil {
		log.Fatalf("Failed to get Jira issue: %v", err)
	}

	// Validate the issue state
	if err := validateTicketState(issue); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("Validation succeeded: Ticket is in a valid state\n")
	}
}

// getJiraIssue fetches issue details using the Jira client
func getJiraIssue(client *jira.Client, issueKey string) (*jira.Issue, error) {
	issue, err := client.GetIssue(issueKey)
	if err != nil {
		return nil, fmt.Errorf("error fetching issue %s: %v", issueKey, err)
	}
	return issue, nil
}
