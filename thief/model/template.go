package model

import (
	"github.com/sokool/scraper/thief/graph"
	"github.com/sokool/scraper/thief/http"
	query "github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"github.com/clbanning/mxj"
	"github.com/sokool/console"
	"github.com/sokool/scraper/storage"
	"fmt"
	"strconv"
	"time"
)

type Object map[string]interface{}

type Template struct {
	config   *configuration
	onResult func(Object)
	storage  storage.Storage
	stats    *Stats
}

func (this *Template) OnResult(f func(Object)) *Template {
	this.onResult = f

	return this
}

func (this *Template) Name() string {
	return this.config.Name
}

func (this *Template) Root() graph.Node {
	return this.config.root()
}

func (this *Template) Done() {
	ustamp := strconv.Itoa(int(time.Now().Unix()))
	this.storage.Flush(ustamp + "_" + this.Name())
}

func (this *Template) onHit(d *query.Document) {
	object := this.config.Schema.scrape(d)
	if this.onResult != nil {
		this.onResult(object)
	}

	this.storage.Add(object)

	if this.storage.Count() == 200 {
		this.Done()
	}

}

func (this *Template) Visit(in graph.Node) []graph.Node {
	var nodes []graph.Node
	action := in.(*node)
	this.stats.OpenNodes++
	document := http.Get(action.url)
	this.stats.Waiting--
	if document == nil {
		this.stats.OpenNodes--
		this.stats.Errors++
		return nodes
	}

	if !action.hasSelector() {
		this.stats.DocumentsFound++
		this.stats.OpenNodes--

		this.onHit(document)

		fmt.Printf("\r%s", this.stats)
		return nodes
	}

	document.Find(action.Selector).Each(func(n int, selection *query.Selection) {
		href, _ := selection.Attr("href")
		this.config.neighborsFunc(action, func(n *node) {
			n.setUrl(this.config.prepareURL(href))
			nodes = append(nodes, n)
		})
	})

	this.stats.Waiting = len(nodes) + this.stats.Waiting
	this.stats.OpenNodes--
	fmt.Printf("\r%s", this.stats)
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

	return &Template{
		config:config,
		storage: storage.Get(config.Schema.Storage, []string{config.Name}),
		stats: &Stats{},
	}
}
