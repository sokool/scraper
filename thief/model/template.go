package model

import (
	"github.com/sokool/scraper/thief/graph"
	"github.com/sokool/scraper/thief/http"
	query "github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"github.com/clbanning/mxj"
	"github.com/sokool/console"
	"github.com/sokool/scraper/storage"
)

type Object map[string]interface{}

type Template struct {
	config   *configuration
	onResult func(Object)
	storage  storage.Storage
}

func (this *Template) OnResult(f func(Object)) {
	this.onResult = f
}

func (this *Template) Name() string {
	return this.config.Name
}

func (this *Template) Root() graph.Node {
	return this.config.root()
}

func (this *Template) Done() {
	this.storage.Flush()
}

func (this *Template) onHit(doc *query.Document) {
	_, object := this.config.Schema.structure(doc)
	if this.onResult != nil {
		this.onResult(object)
	}

	this.storage.Add(object)
}

func (this *Template) Visit(in graph.Node) []graph.Node {
	var nodes []graph.Node
	action := in.(*node)
	document := http.Get(action.url)

	if !action.hasSelector() {
		this.onHit(document)
		return nodes
	}

	document.Find(action.Selector).Each(func(n int, selection *query.Selection) {
		href, _ := selection.Attr("href")
		this.config.neighborsFunc(action, func(n *node) {
			n.setUrl(this.config.prepareURL(href))
			nodes = append(nodes, n)
		})
	})

	return nodes
}

func FromJsonFile(file string) *Template {

	dat, e1 := ioutil.ReadFile(file)
	json, e2 := mxj.NewMapJson(dat)

	if e1 != nil || e2 != nil {
		console.Log(e1)
		console.Log(e2)
		return nil
	}

	var config *configuration
	json.Struct(&config)
	config.load()

	return &Template{
		config:config,
		storage: storage.Get(config.Schema.Storage, []string{config.Name}),
	}
}
