// Queue and Node Library
// Last Modified: Seonny, 2021.03.06
// * 아래 링크에 들어가시면 아래 메서드 들의 설명을 보실 수 있습니다.
// Reference: https://docs.oracle.com/javase/7/docs/api/java/util/Queue.html

package queue

import node "hackathon/node"

// Queue struct
// TODO Please implement this struct
type Queue struct {
	size int
}

// Constructor of queue
// TODO Please implement this constructor
func Constructor(size int) *Queue {
	queue := Queue{ size }

	return &queue
}

// Add the node into queue
// TODO Please implement this method
func (q *Queue) Add(node *node.Node) bool {
	return true
}

// Poll a node from queue head
// TODO Please implement this method
func (q *Queue) Poll() *node.Node {
	return nil
}

// Peek node from queue
// TODO Please implement this method
func (q *Queue) Peek() *node.Node {
	return nil
}

// Find the node
// TODO Please implement this method
func (q *Queue) Find(id int) int {
	return 0
}

// SetSize of the queue
// TODO Please implement this method
func (q *Queue) SetSize(size int) {
	return
}

// Size of the queue
// TODO Please implement this method
func (q *Queue) Size() int {
	return 0
}
