package thief

import (
	"github.com/clbanning/mxj"
	"github.com/sokool/console"
	"io/ioutil"
	"github.com/sokool/scraper/thief/scheme"
	"github.com/sokool/scraper/thief/graph"
)

type Configuration struct {
	Url     string `json:"url"`
	Storage string `json:"storage"`
	Graph   *graph.Graph `json:"graph"`
	Schemas []*scheme.Scheme `json:"schemes"`
}

//func (this *configuration) Url() *url.URL {
//	return this.url
//}
//
//func (this *configuration) buildNode(path string) *graph.Node {
//	var node *graph.Node
//
//	o, _ := this.json.ValueForPath(path)
//	m := o.(map[string]interface{})
//
//	s := mxj.Map(m)
//
//	e := s.Struct(&node)
//
//	if e != nil {
//		panic(e)
//		return nil
//	}
//
//	return node
//}

//func (this *configuration) buildSchemes() *configuration {
//	//s, _ := this.json.ValuesForPath("schemes")
//	//for _, o := range s {
//	//	var scheme *scheme.Scheme
//	//	j := mxj.Map(o.(map[string]interface{}))
//	//
//	//	for name, a := range j {
//	//		console.Log(name, a)
//	//	}
//	//	j.Struct(&scheme)
//	//
//	//	this.schemas[scheme.Name] = scheme
//	//}
//
//	return this
//}
//
//func (this *configuration) buildRoot() *configuration {
//	this.root = this.buildNode("root")
//	return this
//}
//
//func (this *configuration) buildStorage() *configuration {
//	name, _ := this.json.ValueForPathString("storage")
//	this.storage = storage.Get(name)
//
//	return this
//}
//
//func (this *configuration) buildURL() *configuration {
//	var err error
//	link, _ := this.json.ValueForPathString("url")
//	this.url, err = url.Parse(link)
//	if err != nil {
//		panic(err)
//	}
//
//	return this
//}

//func (this *configuration) buildNodes() *configuration {
//	//s, _ := this.json.ValuesForPath("nodes")
//	//for _, o := range s {
//	//	var node *graph.Node
//	//	j := mxj.Map(o.(map[string]interface{}))
//	//
//	//	j.Struct(&node)
//	//
//	//	this.nodes[node.Name] = node
//	//}
//
//	return this
//}

//func (this *configuration) Node(name string) *graph.Node {
//	return this.nodes[name]
//}

//func (c *configuration) NodeNeighbors(name string) map[string]*graph.Node {
//	neighbors := make(map[string]*graph.Node)
//	nString := strings.TrimSpace(c.Node(name).Neighbors)
//	if nString == "" {
//		return neighbors
//	}
//
//	for _, index := range strings.Split(nString, ",") {
//		console.Log(index)
//		node := c.Node(strings.TrimSpace(index))
//		neighbors[node.Name] = node
//	}
//	return neighbors
//}

//func (c *Configuration) Root() *graph.Node {
//	return c.root
//}
//
//func (c *Configuration) Storage() func(...string) storage.Storage {
//	return c.storage
//}

//func (c *configuration) Scheme(name string) *scheme.Scheme {
//	return c.schemas[name]
//}

func Config(filepath string) *Configuration {

	dat, e1 := ioutil.ReadFile(filepath)
	json, e2 := mxj.NewMapJson(dat)

	if e1 != nil || e2 != nil {
		console.Log(e1)
		console.Log(e2)
		return nil
	}

	var config *Configuration
	json.Struct(&config)

	//config.
	//buildURL().
	//buildStorage().
	//buildRoot().
	//buildNodes().
	//buildSchemes()


	return config
}