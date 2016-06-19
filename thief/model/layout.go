package model

import (
	query "github.com/PuerkitoBio/goquery"
)

type layout struct {
	Root    string   `json:"root"`
	Actions []node `json:"actions"`
	it      int
	onHit   func(schema string, document *query.Document)
}

//func (this *layout) Neighbors(a *node) []*node {
//	var neighbors []*node
//	for _, name := range a.Childrens() {
//		child, ok := this.Action(name)
//		if (!ok) {
//			continue
//		}
//		neighbors = append(neighbors, &child)
//	}
//
//	return neighbors
//}
//
//func (this *layout) Action(name string) (node, bool) {
//	for _, action := range this.Actions {
//		if action.Name == name {
//			return action, true
//		}
//	}
//
//	return node{}, false
//}
//
//func (this *layout) Visit(in graph.Node) ([]graph.Node) {
//	var nodes []graph.Node
//	action := in.(*node)
//	document := http.Get(action.url)
//
//	if action.Schema != "" {
//		this.onHit(action.Schema, document)
//	}
//
//	if action.Selector != "" {
//		document.Find(action.Selector).Each(func(n int, selection *query.Selection) {
//			href, ok := selection.Attr("href")
//			if (!ok) {
//				return
//			}
//			for _, neighbor := range this.Neighbors(action) {
//				neighbor.url = href
//				nodes = append(nodes, neighbor)
//			}
//
//		})
//	}
//	return nodes
//}
