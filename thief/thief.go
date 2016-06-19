package main

import (
	"runtime"
	"github.com/sokool/scraper/thief/model"
	"github.com/sokool/scraper/thief/graph"
	"fmt"
)

type Thief struct {
	list map[string]*model.Template
}

func (this *Thief) Run() {
	workers := 32
	fmt.Printf("Workers: %d\n", workers)
	for _, template := range this.list {
		graph.New(template, workers).GoBFS()
	}

}

func (this *Thief) Add(template *model.Template) *Thief {
	this.list[template.Name()] = template
	return this
}

func New() *Thief {
	return &Thief{
		list: make(map[string]*model.Template),
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	New().
	Add(model.FromJsonFile("/home/sokool/go/src/github.com/sokool/scraper/thief/otomoto.json")).
	Run()

}
