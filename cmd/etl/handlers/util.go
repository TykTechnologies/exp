package handlers

import "github.com/spf13/pflag"

// NewFlagSet is used for command flags.
func NewFlagSet(name string) *pflag.FlagSet {
	fs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	fs.SetInterspersed(true)
	return fs
}
