package main

import (
	"runtime"
	"github.com/sokool/scraper/thief/model"
	"github.com/sokool/scraper/thief/graph"
)

type Thief struct {
	config map[string]*model.Configuration
}

func (this *Thief) Run() {
	for _, config := range this.config {
		graph.New(config.Layout, 8)
		//var nodes []interface{}
		//this.graph.push(append(nodes, config.root))
		//pages := this.graph.find(config)
		//fmt.Println(pages)
	}

}

func (this *Thief) Add(config *model.Configuration) *Thief {
	this.config[config.Name] = config
	return this
}

func New() *Thief {
	return &Thief{
		config: make(map[string]*model.Configuration),
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	New().
	Add(model.Config("/home/sokool/go/src/github.com/sokool/scraper/thief/homegate.json")).
	Run()

}
