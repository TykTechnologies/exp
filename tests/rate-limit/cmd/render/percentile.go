package main

import (
	"sort"
)

// Percentile calculates the pth percentile of a slice of float64
func Percentile(data []float64, p float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sort.Float64s(data)
	k := float64(len(data)-1) * p / 100
	f := int(k)
	c := f + 1
	if c >= len(data) {
		return data[f]
	}
	return data[f] + (data[c]-data[f])*(k-float64(f))
}
