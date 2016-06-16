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

func (this *Graph) Push(n Node) {
	this.inQueue++
	go func(node Node) {
		this.queue <- node
	}(n)
}

func (this *Graph) process(onDone chan <- bool) {
	delay := sync.WaitGroup{}
	for node := range this.queue {
		this.traveler.Visit(node)
		for _, neighbor := range node.Neighbors() {
			delay.Add(1)
			go this.traveler.Visit(neighbor)
		}
	}
	delay.Wait()
}

func (this *Graph) GoBFS() int {
	waitOnDone := make(chan bool)
	for i := 1; i <= this.noOfProcess; i++ {
		go this.process(waitOnDone)
	}
	<-waitOnDone
	//this.traveler.OnFinish()

	return 0
}

//func (this *Graph) bfs(t Traveler) int {
//      wait := make(chan bool)
//      for i := 1; i <= this.noOfProcess; i++ {
//            go func() {
//                  for nodes := range this.queue {
//                        delay := sync.WaitGroup{}
//                        var out []interface{}
//                        for _, item := range nodes {
//                              delay.Add(1)
//                              go func(node interface{}) {
//                                    result := t.Visit(node)
//                                    this.visits++
//                                    if (result == nil) {
//                                          return
//                                    }
//                                    out = append(out, result)
//                                    delay.Done()
//                              }(item)
//                        }
//                        delay.Wait()
//                        for _, neighbor := range out {
//                              nodes := t.visitNeighbors(neighbor)
//                              if len(nodes) == 0 {
//                                    t.OnLast(neighbor)
//                                    continue
//                              }
//                              this.push(nodes)
//                        }
//                        this.inQueue--
//                        if (this.inQueue == 0) {
//                              close(this.queue)
//                              wait <- false
//                              return
//                        }
//                  }
//            }()
//      }
//      <-wait
//      t.OnFinish()
//
//      return this.visits
//}

func New(traveler Traveler, workers int) *Graph {
	return &Graph{
		noOfProcess: workers,
		visits:      0,
		inQueue:     0,
		queue:       make(chan Node),
		traveler:    traveler,
	}
}
