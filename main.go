package main

import (
	node "hackathon/node"
	queue "hackathon/queue"
)

func main() {
	a := queue.Constructor(1, 2)
	b := node.Constructor(1)

	a.Add(b)
	a.Find(b.ID())
}
