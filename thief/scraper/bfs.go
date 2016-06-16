package scraper

import (
	"sync"
)

type visitor interface {
	OnNode(node interface{}) (neighbor interface{})        //findNeighbor?
	OnNeighbor(neighbor interface{}) (nodes []interface{}) //findNode?
	OnLast(neighbor interface{})                           //onLastNode?
	OnFinish()
}

type bfs struct {
	workers int                //Number of workers/threads/gorutines for nodes processing
	visits  int                //Number of visited nodes
	inQueue int                //Number of left nodes to visit
	queue   chan []interface{} //Represents FIFO queue
}

type neighbor interface{}

func (this *bfs) push(node []interface{}) {
	this.inQueue++
	go func(n []interface{}) {
		this.queue <- n
	}(node)
}

func (this *bfs) find(r visitor) int {
	wait := make(chan bool)
	for i := 1; i <= this.workers; i++ {
		go func() {
			for nodes := range this.queue {
				delay := sync.WaitGroup{}
				var out []interface{}
				for _, item := range nodes {
					delay.Add(1)
					go func(node interface{}) {
						result := r.OnNode(node) //pobiera dokument
						this.visits++
						if result == nil {
							return
						}
						out = append(out, result)
						delay.Done()
					}(item)
				}
				delay.Wait()
				for _, neighbor := range out {
					nodes := r.OnNeighbor(neighbor) //szuka urli url'e
					if len(nodes) == 0 {
						r.OnLast(neighbor)
						continue
					}
					this.push(nodes)
				}
				this.inQueue--
				if this.inQueue == 0 {
					close(this.queue)
					wait <- false
					return
				}
			}
		}()
	}
	<-wait
	r.OnFinish()
	return this.visits
}

func newBFS() *bfs {
	return &bfs{
		workers: 32,
		visits:  0,
		inQueue: 0,
		queue:   make(chan []interface{}),
	}
}
