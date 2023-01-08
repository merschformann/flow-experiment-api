package graph

import (
	"github.com/merschformann/flow-experiment-api/plugin"
)

type Node struct {
	Edges []*Edge
	ID    int
}

type Edge struct {
	From, To *Node
	Cost     float64
}

type AStarSearch interface {
	Search(start, goal *Node) []*Node
}

// NewAStarSearch returns a new a-star search instance.
func NewAStarSearch() AStarSearch {
	plugin.Connect(con, &newAStarSearch)
	return newAStarSearch()
}

var (
	con            = plugin.NewConnection()
	newAStarSearch func() AStarSearch
)
