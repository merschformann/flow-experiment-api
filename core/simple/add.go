package simple

import "github.com/merschformann/flow-experiment-api/plugin"

// SimpleAdd implements a simple addition to test the performance of the plugin
// integration vs. a direct call.
func SimpleAdd(a, b int) int {
	plugin.Connect(con, &simpleAdd)
	return simpleAdd(a, b)
}
