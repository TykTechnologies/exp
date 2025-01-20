package example

import (
        "sync/atomic"
)

type ExampleBadStruct struct {
        // ruleid: find-atomic-value-usage
        Some atomic.Value
}

func ExampleBad() {
        // ruleid: find-atomic-value-usage
        var v atomic.Value
        v.Store("example")
}

func ExampleGood() {
	var p atomic.Pointer[string]
	p.Store(new(string))
}
