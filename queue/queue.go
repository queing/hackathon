// Queue and Node Library
// Last Modified: Seonny, 2021.03.08
// * 아래 링크에 들어가시면 아래 메서드 들의 설명을 보실 수 있습니다.
// Reference: https://docs.oracle.com/javase/7/docs/api/java/util/Queue.html

package queue

import (
	node "hackathon/node"
)

// Queue struct
// TODO Please implement this struct
type Queue struct {
	id int
	head *node.Node
	tail *node.Node
}

// Constructor of queue
// TODO Please implement this constructor
func Constructor(id int) *Queue {
	queue := Queue{id, nil, nil}

	return &queue
}

// ID of the queue
// TODO Please implement this method
func (q *Queue) ID() int {
	return q.id
}

// Add the node into queue
// TODO Please implement this method
// * Return order of node
func (q *Queue) Add(node *node.Node) int {
	if q.head == nil && q.tail == nil {
		q.head = node
		q.tail = node
		return 0
	}

	order := q.tail.Order() + 1
	
	q.tail.SetNext(node)
	q.tail = node

	return order
}

// Poll a node from queue head
// TODO Please implement this method
func (q *Queue) Poll() *node.Node {
	if q.head == nil {
		return nil
	}

	result := q.head
	q.head = q.head.Next()
	
	return result
}

// Peek a node from queue
// TODO Please implement this method
func (q *Queue) Peek() *node.Node {
	return q.head
}

// Find the node by id
// TODO Please implement this method
// * Return pointer of node and order
func (q *Queue) Find(id int) (*node.Node, int) {
	count := 0

	for current := q.head; current != nil; current = current.Next() {
		if current.ID() == id {
			return current, count
		}

		count = count + 1
	}

	return nil, 0
}

// Count of all nodes in queue
func (q *Queue) Count() int {
	count := 0

	for current := q.head; current != nil; current = current.Next() {
		count = count + 1
	}

	return count
}