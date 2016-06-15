package thief

import (
	query "github.com/PuerkitoBio/goquery"
	"github.com/sokool/scraper/storage"
	"net/url"
	"strings"
	"github.com/sokool/console"
)

type Template struct {
	root    element          //root element of graph
	rules   map[string]*rule //set of rules based on each neighbor is processed
	storage storage.Storage  //
}

type element struct {
	url       string
	doc       *query.Document
	neighbors []string
}

type rule struct {
	selector string
	nodes    []string
}

func (this *Template) OnNode(node interface{}) interface{} {
	item := node.(element)
	item.doc = getDocument(item.url)

	if item.doc == nil {
		return nil
	}
	return item
}

func (this *Template) OnNeighbor(neighbor interface{}) []interface{} {
	a := neighbor.(element)
	var out []interface{}
	for _, neighbor := range a.neighbors {
		rule := this.rules[neighbor]
		a.doc.Find(rule.selector).Each(func(n int, selection *query.Selection) {
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

func (this *Template) OnLast(node interface{}) {
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

func (this *Template) OnFinish() {
	this.storage.Flush()
}

func (this *Template) Add(name string, selector interface{}, neighbors ...string) (*Template) {
	var nodes []string
	for _, neighbor := range neighbors {
		nodes = append(nodes, neighbor)
	}

	switch c := selector.(type) {
	case string :
		this.rules[name] = &rule{c, nodes}
		break
	}
	return this

}

func NewScheme(format string, link string, nodes ...string) *Template {
	u, _ := url.Parse(link)
	return &Template{
		root: element{url:link, neighbors: nodes},
		storage: storage.Get(format)(strings.Replace(u.Host, ".", "-", -1)),
		rules: make(map[string]*rule),
	}
}