package model

import (
	"fmt"
	"net/url"
)

type configuration struct {
	Name   string`json:"name"`
	Url    string`json:"url"`
	Root   string
	Nodes  map[string]*node
	Schema *scheme
}

func (this *configuration) load() {
	//this.data = make(map[string]interface{})
	//this.Layout.onHit = this.found
}

func (this *configuration) node(name string) *node {
	if val, ok := this.Nodes[name]; ok {
		return val
	}

	panic(fmt.Sprintf("Configuration has no %s node", name))
}

func (this *configuration) root() *node {
	return this.nodeWithURL(this.Root, this.Url)
}

func (this *configuration) prepareURL(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		panic(fmt.Sprintf("Given link [%s] is not valid", link))
	}

	if u.Scheme != "" {
		return link
	}

	u, _ = url.Parse(this.Url)

	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, link)
}

func (this *configuration) nodeWithURL(name, link string) *node {
	return this.node(name).copy().setUrl(this.prepareURL(link))
}

func (this *configuration) neighborsFunc(of *node, filter func(*node)) {

	if !of.hasNeighbors() {
		n := &node{}
		filter(n)
	}

	for _, name := range of.Neighbors {
		n := this.node(name).copy()
		filter(n)
	}

}
