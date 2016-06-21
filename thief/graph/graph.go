package graph

import (
	"sync"
)

type Graph struct {
	noOfProcess int       //Number of workers/threads/gorutines for nodes processing
	visits      int       //Number of visited nodes
	inQueue     int       //Number of left nodes to visit
	queue       chan Node //Represents FIFO queue
	traveler    Traveler
}

func (this *Graph) push(n Node) *Graph {
	this.inQueue++
	go func(node Node) {
		this.queue <- node
	}(n)
	return this
}

func (this *Graph) open(n Node) []Node {
	childrens := this.traveler.Visit(n)
	return childrens
}

func (this *Graph) openNeighbors(nodes []Node) []Node {
	delay := sync.WaitGroup{}
	var out []Node
	for _, n := range nodes {
		delay.Add(1)
		go func(node Node) {
			for _, neighbor := range this.open(node) {
				out = append(out, neighbor)
			}
			delay.Done()
		}(n)
	}
	delay.Wait()

	return out
}

func (this *Graph) bfs(onDone chan <- bool) {
	for node := range this.queue {
		neighbors := this.open(node)
		for _, toVisit := range this.openNeighbors(neighbors) {
			this.push(toVisit)
		}
		this.inQueue--
		if (this.inQueue == 0) {
			close(this.queue)
			onDone <- true
		}
	}
}

func (this *Graph) GoBFS() int {
	waitOnDone := make(chan bool)
	for i := 1; i <= this.noOfProcess; i++ {
		go this.bfs(waitOnDone)
	}
	<-waitOnDone
	this.traveler.Done()

	return 0
}

func New(traveler Traveler, workers int) *Graph {
	graph := &Graph{
		noOfProcess: workers,
		visits:      0,
		inQueue:     0,
		queue:       make(chan Node),
		traveler:    traveler,
	}
	graph.push(traveler.Root())
	return graph
}
