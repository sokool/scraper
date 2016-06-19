package graph

import (
	"sync"
	"fmt"
)

type Graph struct {
	noOfProcess int       //Number of workers/threads/gorutines for nodes processing
	visits      int       //Number of visited nodes
	inQueue     int       //Number of left nodes to visit
	queue       chan Node //Represents FIFO queue
	traveler    Traveler
}

func (this *Graph) push(n Node) *Graph {
	go func(node Node) {
		this.queue <- node
	}(n)
	return this
}

func (this *Graph) visitNeighbors(nodes []Node) []Node {
	delay := sync.WaitGroup{}
	var out []Node
	for _, n := range nodes {
		delay.Add(1)
		go func(node Node) {
			this.inQueue++
			childrens := this.traveler.Visit(node)
			for _, neighbor := range childrens {
				out = append(out, neighbor)
			}
			this.inQueue--
			delay.Done()
		}(n)
	}
	delay.Wait()

	return out
}

func (this *Graph) bfs(onDone chan <- bool) {
	for node := range this.queue {
		this.inQueue++
		neighbors := this.traveler.Visit(node)
		for _, toVisit := range this.visitNeighbors(neighbors) {
			this.push(toVisit)
		}
		this.inQueue--
		if (this.inQueue == 0) {
			//close(this.queue)
			//onDone <- true
		}
	}
}

func (this *Graph) GoBFS() int {
	waitOnDone := make(chan bool)
	for i := 1; i <= this.noOfProcess; i++ {
		go this.bfs(waitOnDone)
	}
	<-waitOnDone
	fmt.Println("DONE")
	//this.traveler.OnFinish()

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
