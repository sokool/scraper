package thief

import (
	"fmt"
)

type Thief struct {
	visitor   *bfs
	templates map[string]*Template
}

func (this *Thief) Run() {
	for _, template := range this.templates {
		var nodes []interface{}
		this.visitor.push(append(nodes, template.root))
		pages := this.visitor.find(template)
		fmt.Println(pages)
	}


}

func (this *Thief) Add(e *Template) *Thief {
	this.templates[e.root.url] = e
	return this
}

func New() *Thief {
	return &Thief{
		templates: make(map[string]*Template),
		visitor: newBFS(),
	}
}