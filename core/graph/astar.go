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
func NewAStarSearch(heuristic func(node, goal *Node) float64) AStarSearch {
	plugin.Connect(con, &newAStarSearch)
	return newAStarSearch(heuristic)
}

var (
	con            = plugin.NewConnection()
	newAStarSearch func(heuristic func(node, goal *Node) float64) AStarSearch
)
