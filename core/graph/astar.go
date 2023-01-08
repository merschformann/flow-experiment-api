package graph

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
