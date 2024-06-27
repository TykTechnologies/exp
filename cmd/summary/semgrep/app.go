package semgrep

import (
	"encoding/json"
	"io"
	"os"
	"sort"
)

func readInput(input string) ([]byte, error) {
	if input == "" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(input)
}

func openOutput(output string) (io.WriteCloser, error) {
	if output == "" {
		return os.Stdout, nil
	}
	return os.Create(output)
}

type aggregateResult struct {
	CheckID string `json:"check_id"`
	Extra   struct {
		IsIgnored bool `json:"is_ignored"`
	} `json:"extra"`
}

type aggregateSource struct {
	Results []aggregateResult `json:"results,omitempty"`
}

type aggregateResultFlat struct {
	CheckID string

	CountIgnored int
	Count        int
}

func (r aggregateResult) Flatten() *aggregateResultFlat {
	var countIgnored int
	if r.Extra.IsIgnored {
		countIgnored++
	}

	return &aggregateResultFlat{
		CheckID:      r.CheckID,
		CountIgnored: countIgnored,
		Count:        1,
	}
}

func templateData(input []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	// main data from semgrep report json
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	// aggregated data for some grouping/ordering in reports
	var dataForAggregate aggregateSource
	if err := json.Unmarshal(input, &dataForAggregate); err != nil {
		return nil, err
	}

	// group results by group id
	checks := make(map[string]*aggregateResultFlat)
	for _, resultRaw := range dataForAggregate.Results {
		// holds a single result in a flat structure
		result := resultRaw.Flatten()

		check, ok := checks[result.CheckID]
		if !ok {
			// initial state
			checks[result.CheckID] = result
			continue
		}

		// increment totals
		check.Count += result.Count
		check.CountIgnored += result.CountIgnored
	}

	// flatten map to slice and order by count descending
	aggregateChecks := make([]*aggregateResultFlat, 0, len(checks))
	for _, check := range checks {
		aggregateChecks = append(aggregateChecks, check)
	}

	sort.SliceStable(aggregateChecks, func(i, j int) bool {
		return aggregateChecks[i].Count > aggregateChecks[j].Count
	})

	data["checks"] = aggregateChecks

	return data, nil
}
