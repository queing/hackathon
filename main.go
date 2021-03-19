package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	node "hackathon/node"
	queue "hackathon/queue"
)

func main() {
	testSynchronously()
}

func testSynchronously() {
	Q := queue.Constructor(1)

	// Track 1: Add 10000 Nodes.
	track1 := time.Now()

	for i := 0; i < 100000; i = i + 1 {
		N := node.Constructor(i)
		order := Q.Add(N)
		N.SetOrder(order)

		if N.ID() != order {
			msg := fmt.Sprintf("Add Error, expect order: %d, result order: %d", N.ID(), order)
			log.Fatal(msg)
		}
	}

	track1SpentTime := time.Since(track1)
	fmt.Printf("Track1 : %s\n", track1SpentTime)

	// Track 2: Find some nodes from queue.
	track2 := time.Now()

	for i := 0; i < 500; i = i + 1 {
		id := rand.Intn(10000) // 0 <= id < 10000
		N := Q.Find(id)

		if N.ID() != id {
			msg := fmt.Sprintf("Find Error, expect id: %d, result id: %d", id, N.ID())
			log.Fatal(msg)
		}
	}

	track2SpentTime := time.Since(track2)
	fmt.Printf("Track2 : %s\n", track2SpentTime)

	// Track 3: Peek a node
	track3 := time.Now()

	Q.Peek()

	track3SpentTime := time.Since(track3)
	fmt.Printf("Track3 : %s\n", track3SpentTime)

	// Track 4: Poll all nodes from queue.
	track4 := time.Now()

	for i := 0; i < 10000; i = i + 1 {
		Q.Poll()
	}

	track4SpentTime := time.Since(track4)
	fmt.Printf("Track4 : %s\n", track4SpentTime)
}

func testAsynchronously() {
	Q := queue.Constructor(1)

	// Track 1: Add 10000 Nodes.
	track1 := time.Now()
	ch := make(chan int, 100000)

	go func(ch chan<- int) {
		for i := 0; i < 100000; i = i + 1 {
				ch <- i
		}

		close(ch)
	}(ch)

	for id := range ch {
		N := node.Constructor(id)
		order := Q.Add(N)
		N.SetOrder(order)
	}

	track1SpentTime := time.Since(track1)
	fmt.Printf("Track1 : %s\n", track1SpentTime)
}
