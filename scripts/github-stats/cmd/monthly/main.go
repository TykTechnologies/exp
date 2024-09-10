package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func main() {
	var filename = "data/github-actions-runs.json"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if err := start(filename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type WorkflowRun struct {
	Path      string    `json:"path"`
	Status    string    `json:"conclusion"`
	CreatedAt time.Time `json:"created_at"`
	Branch    string    `json:"head_branch"`
}

type Message struct {
	WorkflowRuns []WorkflowRun `json:"workflow_runs"`
}

// Grouping structure to hold runs by date and status (success or failure)
type GroupedRuns struct {
	Date    string
	Status  string
	Branch  string
	Success int
	Failure int
}

func start(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error opening gh actions runs: %w", err)
	}
	defer f.Close()

	var total int64
	groupedRuns := make(map[string]map[string]map[string]int) // Date -> Branch -> Status (success/failure) -> Count

	// Define the cutoff date (30 days ago)
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	dec := json.NewDecoder(f)
	for {
		var m Message
		err := dec.Decode(&m)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return fmt.Errorf("Error decoding json message: %w", err)
		}

		for _, run := range m.WorkflowRuns {
			// Filter out runs that are not within the last 30 days
			if run.CreatedAt.Before(thirtyDaysAgo) {
				continue
			}

			// Only include specific workflow path
			if run.Path != ".github/workflows/ci-tests.yml" {
				continue
			}

			// Group by date, branch, and status (success or failure)
			date := run.CreatedAt.Format("2006-01-02")
			branch := run.Branch
			status := run.Status // "success" or "failure"

			if groupedRuns[date] == nil {
				groupedRuns[date] = make(map[string]map[string]int)
			}

			if groupedRuns[date][branch] == nil {
				groupedRuns[date][branch] = make(map[string]int)
			}

			groupedRuns[date][branch][status]++
			total++
		}
	}

	// Output the results
	fmt.Printf("Total runs: %d\n\n", total)

	fmt.Println("Grouped runs by date, branch, and status (success/failure):")
	var sortedDates []string
	for date := range groupedRuns {
		sortedDates = append(sortedDates, date)
	}
	// Sort dates in descending order (newest to oldest)
	sort.Sort(sort.Reverse(sort.StringSlice(sortedDates)))

	for _, date := range sortedDates {
		branches := groupedRuns[date]
		totalSuccess := 0
		totalFailure := 0
		totalRuns := 0

		// Count total success and failure for the day
		for _, statuses := range branches {
			totalSuccess += statuses["success"]
			totalFailure += statuses["failure"]
		}
		totalRuns = totalSuccess + totalFailure

		// Calculate the failure rate as a percentage
		failureRate := float64(totalFailure) / float64(totalRuns) * 100

		// Print the summary for the day
		fmt.Printf("Date: %s | Success: %d | Failure: %d | Failure Rate: %.2f%%\n", date, totalSuccess, totalFailure, failureRate)

		// Print branch-level details, skipping branches with failure: 0
		for branch, statuses := range branches {
			success := statuses["success"]
			failure := statuses["failure"]
			// Only print if there are failures
			if failure > 0 {
				fmt.Printf("  Branch: %s, Success: %d, Failure: %d\n", branch, success, failure)
			}
		}
	}

	return nil
}
