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
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	HeadCommit HeadCommit `json:"head_commit"`
}

type HeadCommit struct {
	Author Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Count int    `json:"-"`
}

func (a *Author) String() string {
	return a.Name // + " <" + a.Email + ">"
}

func (w *WorkflowRun) Duration() time.Duration {
	if w.CreatedAt.IsZero() || w.UpdatedAt.IsZero() {
		return time.Duration(0)
	}

	start, end := time.Since(w.CreatedAt), time.Since(w.UpdatedAt)
	return start - end
}

type Message struct {
	WorkflowRuns []WorkflowRun `json:"workflow_runs"`
}

func start(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error opening gh actions runs: %w", err)
	}
	defer f.Close()

	var total, skipped int64
	var totalDuration time.Duration
	var longest time.Duration

	authors := []*Author{}
	authorMap := map[string]*Author{}

	findAuthor := func(in Author) *Author {
		var author *Author
		var ok bool

		author, ok = authorMap[in.Email]
		if ok {
			return author
		}
		author, ok = authorMap[in.Name]
		if ok {
			return author
		}

		author = &Author{Name: in.Name, Email: in.Email}

		authorMap[in.Name] = author
		authorMap[in.Email] = author

		authors = append(authors, author)
		return author
	}

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
			if run.Path != ".github/workflows/ci-tests.yml" {
				continue
			}

			if run.Status != "completed" {
				continue
			}

			duration := run.Duration()

			if duration > time.Hour {
				skipped++
				continue
			}

			if duration > longest {
				longest = duration
				// fmt.Println("New longest duration:", longest)
			}

			author := findAuthor(run.HeadCommit.Author)
			author.Count++

			totalDuration += duration
			total++
		}
	}

	avgPerRun := totalDuration / time.Duration(total)
	fmt.Printf("Total runs: %d, skipped %d, duration total %s, %s per run\n", total, skipped, totalDuration, avgPerRun)

	fmt.Println("ROI calculations:")

	optimizations := []int{5, 15, 30, 60}
	for _, t := range optimizations {
		fmt.Printf("Optimizing %d seconds saves %s CI time per annum\n", t, time.Duration(total)*time.Duration(t)*time.Second)
	}

	sort.Slice(authors, func(i, j int) bool {
		return authors[i].Count > authors[j].Count
	})

	fmt.Println("Authors:")
	for idx, author := range authors {
		fmt.Printf("#%d %s, count %d\n", idx+1, author, author.Count)
	}

	return nil
}
