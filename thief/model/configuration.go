package model

import (
	"fmt"
	"net/url"
)

type configuration struct {
	Name   string
	Url    string
	Root   string
	Nodes  map[string]*node
	Schema *scheme
}

func (this *configuration) node(name string) *node {
	if n, ok := this.Nodes[name]; ok {
		return n
	}

	panic(fmt.Errorf("Node [%s] has not been found in configuration", name))
}

func (this *configuration) root() *node {
	return this.nodeWithURL(this.Root, this.Url)
}

func (this *configuration) prepareURL(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		panic(fmt.Errorf("Given URL [%s] is not valid", link))
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
