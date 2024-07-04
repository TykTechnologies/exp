package lint

import (
	flag "github.com/spf13/pflag"
	"golang.org/x/exp/slices"
)

type options struct {
	inputFile string
	rules     []string
	exclude   []string
	verbose   bool
	summary   bool
}

// GetRules traverses the rules and excludes any excluded rules.
func (o *options) GetRules() []string {
	if len(o.exclude) == 0 {
		return o.rules
	}

	result := make([]string, 0, len(o.rules))
	for _, ex := range o.rules {
		if slices.Contains(o.exclude, ex) {
			continue
		}
		result = append(result, ex)
	}
	return result
}

func NewOptions() *options {
	cfg := &options{
		inputFile: "schema.json",
		rules: []string{
			"require-comment",
			"require-field-comment",
			"require-dot-or-backtick",
			"require-prefix",
			"require-no-globals",
		},
		verbose: false,
	}
	flag.StringVarP(&cfg.inputFile, "input-file", "i", cfg.inputFile, "input file")
	flag.StringSliceVarP(&cfg.rules, "rules", "", cfg.rules, "linter rules to run")
	flag.StringSliceVarP(&cfg.exclude, "exclude", "", cfg.exclude, "linter rules to exlude")
	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	flag.BoolVarP(&cfg.summary, "summary", "", cfg.summary, "summarize linter issues")
	flag.Parse()

	return cfg
}
