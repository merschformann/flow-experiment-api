package simple

import "github.com/merschformann/flow-experiment-api/plugin"

type SimpleAdder interface {
	Add(a, b int) int
}

type simpleAdder struct {
}

func (s simpleAdder) Add(a, b int) int {
	return a + b
}

func NewSimpleAdder() SimpleAdder {
	plugin.Connect(con, &newSimpleAdder)
	return simpleAdder{}
}

// SimpleAdd implements a simple addition to test the performance of the plugin
// integration vs. a direct call.
func SimpleAdd(a, b int) int {
	plugin.Connect(con, &simpleAdd)
	return simpleAdd(a, b)
}
