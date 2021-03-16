// Node Library
// Last Modified: Seonny, 2021.03.08

package node

// Node struct
// TODO Please implement this struct
type Node struct {
	id   int
	order int
	next *Node
}

// Constructor of node
// TODO Please implement this constructor
func Constructor(id int) *Node {
	node := Node{id, -1, nil}

	return &node
}

// ID of the Node
// TODO Please implement this method
func (n *Node) ID() int {
	return n.id
}

// Order of the Node
// TODO Please implement this method
func (n *Node) Order() int {
	return n.order
}

// SetOrder is to set order of the node
func (n *Node) SetOrder(order int) {
	n.order = order
	return
}

// Next of the Node
// TODO Please implement this method
func (n *Node) Next() *Node {
	return n.next
}

// SetNext is to set next node of the node
// TODO Please implement this method
func (n *Node) SetNext(node *Node) {
	n.next = node
	return
}
