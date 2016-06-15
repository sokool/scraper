package thief

import (
	"sync"
)

type visitor interface {
	OnNode(node) neighbor
	OnNeighbor(neighbor) []node
	OnLast(neighbor)
	OnFinish()
}

type bfs struct {
	workers int         //Number of workers for nodes processing
	visits  int         //Number of visited nodes
	inQueue int         //Number of left nodes to visit
	queue   chan []node //Represents FIFO queue
}

type node interface{}

type neighbor interface{}

func (this *bfs) push(n []node) {
	this.inQueue++
	go func(items []node) {
		this.queue <- items
	}(n)
}

func (this *bfs) find(r visitor) int {
	wait := make(chan bool)
	for i := 1; i <= this.workers; i++ {
		go func() {
			for nodes := range this.queue {
				delay := sync.WaitGroup{}
				var out []neighbor
				for _, item := range nodes {
					delay.Add(1)
					go func(n node) {
						result := r.OnNode(n)
						this.visits++
						if (result == nil) {
							return
						}
						out = append(out, result)
						delay.Done()
					}(item)
				}
				delay.Wait()
				for _, neighbor := range out {
					nodes := r.OnNeighbor(neighbor)
					if len(nodes) == 0 {
						r.OnLast(neighbor)
						continue
					}
					this.push(nodes)
				}
				this.inQueue--
				if (this.inQueue == 0) {
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
		visits: 0,
		inQueue: 0,
		queue: make(chan []node),
	}
}