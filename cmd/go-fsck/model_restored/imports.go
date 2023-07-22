package model

import (
	"sort"

	"golang.org/x/exp/slices"
)

type Imports map[string][]string

func (i *Imports) Add(key, lit string) {
	data := *i
	if data == nil {
		data = make(Imports)
	}
	if set, ok := data[key]; ok {
		if slices.Contains(set, lit) {
			return
		}
		data[key] = append(set, lit)
		return
	}
	data[key] = []string{lit}
	*i = data
}

func (i Imports) Get(key string) []string {
	val, _ := i[key]
	if val != nil {
		sort.Strings(val)
	}
	return val
}
