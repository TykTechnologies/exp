package model

import (
	"strings"
)

type (
	DeclarationKind string

	Declaration struct {
		Kind      DeclarationKind
		File      string
		Name      string   `json:",omitempty"`
		Names     []string `json:",omitempty"`
		Receiver  string   `json:",omitempty"`
		Signature string   `json:",omitempty"`
		Source    string
	}
)

func (d DeclarationKind) String() string {
	return string(d)
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
