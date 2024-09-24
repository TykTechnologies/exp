package model

import (
	"strings"
)

type Declaration struct {
	Kind DeclarationKind
	File string

	SelfContained bool

	Imports []string `json:",omitempty"`

	References map[string][]string `json:",omitempty"`

	Name     string   `json:",omitempty"`
	Names    []string `json:",omitempty"`
	Receiver string   `json:",omitempty"`

	Arguments []string `json:",omitempty"`
	Returns   []string `json:",omitempty"`

	Signature string `json:",omitempty"`
	Source    string `json:",omitempty"`
}

func (d *Declaration) Equal(in *Declaration) bool {
	if d.File == in.File && d.Kind == in.Kind && d.Name == in.Name {
		return true
	}
	return false
}

func (d *Declaration) Keys() []string {
	trimPath := "*."
	if d.Name != "" {
		return []string{
			strings.Trim(d.Receiver+"."+d.Name, trimPath),
		}
	}
	if len(d.Names) != 0 {
		result := make([]string, len(d.Names))
		for k, v := range d.Names {
			result[k] = strings.Trim(d.Receiver+"."+v, trimPath)
		}
	}
	return nil
}