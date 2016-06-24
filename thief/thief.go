package main

import (
	"runtime"
	"github.com/sokool/scraper/thief/model"
	"github.com/sokool/scraper/thief/graph"
	"fmt"
	"github.com/sokool/scraper/thief/ui"

)

type Thief struct {
	list map[string]*model.Template
}

func (this *Thief) Run() {
	workers := 32
	fmt.Printf("Workers: %d\n", workers)
	for _, template := range this.list {
		fmt.Println()
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

	configPath := "/home/sokool/go/src/github.com/sokool/scraper/thief/config/"

	New().
	Add(model.FromJsonFile(configPath + "otomoto.json")).
	//Add(model.FromJsonFile(configPath + "onet.json")).
	//Add(model.FromJsonFile(configPath + "homegate.json")).
	Run()

	ui.Render()
}
