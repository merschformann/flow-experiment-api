package simple

import "github.com/merschformann/flow-experiment-api/plugin"

type NodeStruct struct {
	PreviousStruct *NodeInterface
	NextStruct     *NodeInterface
	ValueStruct    int
}

type NodeInterface interface {
	PreviousInterface() *NodeInterface
	NextInterface() *NodeInterface
	ValueInterface() int
}

func (ns *NodeStruct) PreviousInterface() *NodeInterface {
	return ns.PreviousStruct
}

func (ns *NodeStruct) NextInterface() *NodeInterface {
	return ns.NextStruct
}

func (ns *NodeStruct) ValueInterface() int {
	return ns.ValueStruct
}

type NodeContainer interface {
	NodeAtInterface(index int) NodeInterface
	NodeAtStruct(index int) NodeStruct
	NodeAtAny(index int) any
}

func NewNodeContainer(count int) NodeContainer {
	plugin.Connect(con, &newNodeContainer)
	return newNodeContainer(count)
}
