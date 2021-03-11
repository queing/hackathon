// Node Library
// Last Modified: Seonny, 2021.03.08

package node

// Node struct
// TODO Please implement this struct
type Node struct {
	id   int
	prev *Node
	next *Node
}

// Constructor of node
// TODO Please implement this constructor
func Constructor(id int) *Node {
	node := Node{id, nil, nil}

	return &node
}

// ID of the Node
// TODO Please implement this method
func (n *Node) ID() int {
	return 0
}

// Prev of the Node
// TODO Please implement this method
func (n *Node) Prev() *Node {
	return nil
}

// SetPrev is to set prev node of the node
// TODO Please implement this method
func (n *Node) SetPrev(node *Node) {
	return
}

// Next of the Node
// TODO Please implement this method
func (n *Node) Next() *Node {
	return nil
}

// SetNext is to set next node of the node
// TODO Please implement this method
func (n *Node) SetNext(node *Node) {
	return
}
