// Queue and Node Library
// Last Modified: Seonny, 2021.03.08
// * 아래 링크에 들어가시면 아래 메서드 들의 설명을 보실 수 있습니다.
// Reference: https://docs.oracle.com/javase/7/docs/api/java/util/Queue.html

package queue

import node "hackathon/node"

// Queue struct
// TODO Please implement this struct
type Queue struct {
	id int
	size int
}

// Constructor of queue
// TODO Please implement this constructor
func Constructor(id int, size int) *Queue {
	queue := Queue{ id, size }

	return &queue
}

// Size of the queue
// TODO Please implement this method
func (q *Queue) Size() int {
	return 0
}

// SetSize is to set size of the queue
// TODO Please implement this method
func (q *Queue) SetSize(size int) {
	return
}

// Add the node into queue
// TODO Please implement this method
// * Return order of node
func (q *Queue) Add(node *node.Node) int {
	return 0
}

// Poll a node from queue head
// TODO Please implement this method
func (q *Queue) Poll() *node.Node {
	return nil
}

// Peek a node from queue
// TODO Please implement this method
func (q *Queue) Peek() *node.Node {
	return nil
}

// Find the node by id
// TODO Please implement this method
// * Return pointer of node and order
func (q *Queue) Find(id int) (*node.Node, int) {
	return nil, 0
}
