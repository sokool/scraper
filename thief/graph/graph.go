package graph

import (
	query "github.com/PuerkitoBio/goquery"
	"github.com/sokool/console"
)


type Graph struct {
	Root *Node `json:"root"`
	List []*Node `json:"nodes"`
}

func (this *Graph) OnNode(node interface{}) interface{} {
	item := node.(Node)
	item.document = getDocument(item.url)

	if item.document == nil {
		return nil
	}
	return item
}

func (this *Graph) OnNeighbor(neighbor interface{}) []interface{} {
	node := neighbor.(Node)
	var out []interface{}
	for _, neighbor := range node.neighbors {
		rule := this.rules[neighbor]
		node.document.Find(rule.selector).Each(func(n int, selection *query.Selection) {
			href, ok := selection.Attr("href")
			if (!ok) {
				return
			}
			//https://github.com/asaskevich/govalidator
			out = append(out, element{url:"http://www.homegate.ch" + href, neighbors: rule.nodes})
		})
	}

	return out
}

func (this *Graph) OnLast(node interface{}) {
	element := node.(element)
	data := make(map[string]interface{})
	for _, neighbor := range element.neighbors {
		rule := this.rules[neighbor]
		element.doc.Find(rule.selector).Each(func(i int, selection *query.Selection) {
			data[neighbor] = selection.Text()
		})
	}
	console.Log(data)
	//test := make(map[string]interface{})
	//test["dupa"] = "TADA"
	//data["test"] = test

	this.storage.Add(data)

}

func (this *Graph) OnFinish() {
	this.storage.Flush()
}




func (this *Graph) Neighbors(name string) {


}

func (this *Graph) Node(name string) {
}
