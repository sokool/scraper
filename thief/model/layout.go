package model

import (
	"github.com/sokool/scraper/thief/graph"
	"github.com/sokool/scraper/thief/http"
)

type Layout struct {
	Root    Action   `json:"root"`
	Actions []Action `json:"actions"`
}

func (this *Layout) Visit(in graph.Node) []graph.Node {
	action := in.(*Action)
	action.document = http.Get(action.url)

	var out []graph.Node
	out = append(out, &Action{})

	return out
	//for _, neighbor := range node.neighbors {
	//      rule := this.rules[neighbor]
	//      node.document.Find(rule.selector).Each(func(n int, selection *query.Selection) {
	//            href, ok := selection.Attr("href")
	//            if (!ok) {
	//                  return
	//            }
	//            //https://github.com/asaskevich/govalidator
	//            out = append(out, element{url:"http://www.homegate.ch" + href, neighbors: rule.nodes})
	//      })
	//}
	//
	//console.Log(action)
}

//func (this *Layout) neighbors(a *Action) {
//      neighbors := make(map[string]*graph.Node)
//      nString := strings.TrimSpace(c.Node(name).Neighbors)
//      if nString == "" {
//            return neighbors
//      }
//
//      for _, index := range strings.Split(nString, ",") {
//            console.Log(index)
//            node := c.Node(strings.TrimSpace(index))
//            neighbors[node.Name] = node
//      }
//      return neighbors
//}

//func (this *Layout) OnNode(node *graph.Node) interface{} {
//      item := node.(Action)
//      item.document = http.Get(item.url)
//
//      if item.document == nil {
//            return nil
//      }
//      return item
//}

//func (this *Layout) OnNeighbor(neighbor interface{}) []interface{} {
//      node := neighbor.(Action)
//      var out []interface{}
//      for _, neighbor := range node.neighbors {
//            rule := this.rules[neighbor]
//            node.document.Find(rule.selector).Each(func(n int, selection *query.Selection) {
//                  href, ok := selection.Attr("href")
//                  if (!ok) {
//                        return
//                  }
//                  //https://github.com/asaskevich/govalidator
//                  out = append(out, element{url:"http://www.homegate.ch" + href, neighbors: rule.nodes})
//            })
//      }
//
//      return out
//}
//
//func (this *Layout) OnLast(node interface{}) {
//      element := node.(element)
//      data := make(map[string]interface{})
//      for _, neighbor := range element.neighbors {
//            rule := this.rules[neighbor]
//            element.doc.Find(rule.selector).Each(func(i int, selection *query.Selection) {
//                  data[neighbor] = selection.Text()
//            })
//      }
//      console.Log(data)
//      //test := make(map[string]interface{})
//      //test["dupa"] = "TADA"
//      //data["test"] = test
//
//      this.storage.Add(data)
//
//}
//
//func (this *Layout) OnFinish() {
//      this.storage.Flush()
//}
//
//func (this *Layout) Neighbors(name string) {
//
//}
//
//func (this *Layout) Node(name string) {
//}
