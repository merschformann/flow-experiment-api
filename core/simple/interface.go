package simple

import "github.com/merschformann/flow-experiment-api/plugin"

var (
	con              = plugin.NewConnection()
	simpleAdd        func(a, b int) int
	newNodeContainer func(count int) NodeContainer
)
