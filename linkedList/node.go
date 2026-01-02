package linkedlist

type Node struct {
	Value int
	Next  *Node
}

type NodeDoubleLink struct {
	value    int
	previous *NodeDoubleLink
	next     *NodeDoubleLink
}

func NewNode(value int, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

func NewNodeDoubleLink(value int, previous *NodeDoubleLink, next *NodeDoubleLink) *NodeDoubleLink {
	return &NodeDoubleLink{
		value:    value,
		previous: previous,
		next:     next,
	}
}
