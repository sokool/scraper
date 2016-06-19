package model

import (
	"github.com/sokool/scraper/thief/graph"
	"github.com/sokool/scraper/thief/http"
	query "github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"github.com/clbanning/mxj"
	"github.com/sokool/console"
	"net/url"
	"fmt"
	"github.com/sokool/scraper/storage"
)

type Object map[string]interface{}

type Template struct {
	config   *configuration
	onResult func(Object)
	storage  storage.Storage

	hit      int
}

func (this *Template) OnResult(f func(Object)) {
	this.onResult = f
}

func (this *Template) Name() string {
	return this.config.Name
}

func (this *Template) Root() graph.Node {
	root := this.config.Nodes[this.config.Root]
	root.url = this.config.Url

	return root
}

func (this *Template) onHit(doc *query.Document) {
	_, object := this.config.Schema.structure(doc)
	if this.onResult != nil {
		this.onResult(object)
	}

	this.hit++
	this.storage.Add(object)
	if this.hit == 50 {
		this.storage.Flush()
	}
}

func (this *Template) Visit(in graph.Node) []graph.Node {
	action := in.(*node)
	document := http.Get(action.url)
	if action.Schema != "" {
		this.onHit(document)
	}

	var nodes []graph.Node
	if action.Selector != "" {

		document.Find(action.Selector).Each(func(n int, selection *query.Selection) {
			href, ok := selection.Attr("href")
			if (!ok) {
				return
			}
			for _, name := range action.Neighbors {
				n := this.config.Nodes[name]
				x := *n
				x.url = href
				this.fill(&x)
				nodes = append(nodes, &x)
			}

		})
	}
	return nodes
}

func (this *Template) fill(n *node) {
	url, err := url.Parse(n.url)
	if err != nil {
		return
	}

	if url.Scheme != "" {
		return
	}

	url, _ = url.Parse(this.config.Url)
	n.url = fmt.Sprintf("%s://%s/%s", url.Scheme, url.Host, n.url)

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
