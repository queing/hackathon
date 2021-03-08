// Node Library
// Last Modified: Seonny, 2021.03.06

package node

// Node struct
// TODO Please implement this struct
type Node struct {
	id int
	prev *Node
	next *Node
}

// Constructor of node
// TODO Please implement this constructor
func Constructor(id int) *Node {
	node := Node{ id, nil, nil }

	return &node
}