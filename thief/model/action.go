package model

import (
	query "github.com/PuerkitoBio/goquery"
	"github.com/sokool/scraper/thief/graph"
)

type Action struct {
	element
	Name      string                  //"nextList",
	Selector  string                  //"a[rel=next]",
	Actions   string `json:"actions"` //"documents, nextList",
	Schema    string                  //"partial"
	neighbors []graph.Node
}

func (this *Action) add(a Action) {
	this.neighbors = append(this.neighbors, graph.Node(a))
}

func (this Action) Neighbors() []graph.Node {

	return this.neighbors
}

type element struct {
	url      string
	document *query.Document
}
