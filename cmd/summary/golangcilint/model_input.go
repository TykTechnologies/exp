package golangcilint

// Root struct to capture the entire JSON structure.
type Root struct {
	Issues []Issue `json:"Issues"`
}

// Issue struct represents each issue in the JSON array.
type Issue struct {
	FromLinter           string   `json:"FromLinter"`
	Text                 string   `json:"Text"`
	Severity             string   `json:"Severity"`
	SourceLines          []string `json:"SourceLines"`
	Pos                  Position `json:"Pos"`
	ExpectNoLint         bool     `json:"ExpectNoLint"`
	ExpectedNoLintLinter string   `json:"ExpectedNoLintLinter"`
}

// Position struct represents the position of each issue.
type Position struct {
	Filename string `json:"Filename"`
	Offset   int    `json:"Offset"`
	Line     int    `json:"Line"`
	Column   int    `json:"Column"`
}
