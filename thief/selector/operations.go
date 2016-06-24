package selector

import (
	. "github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/sokool/console"
	"bitbucket.org/gotamer/cases"
)

type invoker func(*Selection, []string, *result)

type result struct {
	values interface{}
}

func (this *result) set(v interface{}) *result {
	this.values = v
	return this
}

func (this *result) get() interface{} {
	return this.values
}

var operations = map[string]invoker{
	"attr" : attr,
	"map": arrayMap,
	"node": node,
	"after": after,
	"text": text,
	"html": html,
}

func text(s *Selection, params []string, out *result) {
	if len(params) == 1 {
		data := []interface{}{}
		s.Each(func(idx int, itm *Selection) {
			data = append(data, strings.TrimSpace(itm.Text()))
		})
		out.set(data)
	} else {

		out.set(strings.TrimSpace(s.First().Text()))
	}
}

func attr(s *Selection, in []string, out *result) {
	switch len(in) {
	case 2:
		//data := make(map[string]interface{})
		data := []interface{}{}
		s.Each(func(index int, item *Selection) {
			attribute, _ := item.First().Attr(in[0])
			data = append(data, attribute)
		})

		out.set(data)
		break
	case 1:
		attribute, _ := s.First().Attr(in[0])

		out.set(attribute)
		break
	default:
		console.Log("error")
	}
}

func arrayMap(s *Selection, in []string, out *result) {
	data := make(map[string]interface{})
	s.Each(func(idx int, itm *Selection) {
		key := cases.Camel(NewDocumentFromNode(itm.Children().Nodes[0]).Text())
		value := strings.TrimSpace(NewDocumentFromNode(itm.Children().Nodes[1]).Text())
		if key == "" {
			return
		}
		data[key] = value
	})

	out.set(data)
}

func html(s *Selection, in []string, out *result) {
	html, _ := s.Html()
	out.set(html)
}

func node(s *Selection, in []string, out *result) {
	//number, _ := strconv.Atoi(in[0])
	//if len(s.Nodes) > 0 {
	//	o := NewDocumentFromNode(s.Nodes[0])
	//	a := NewDocumentFromNode(o.Nodes[0])
	//	console.Log(a.Text(), number)
	//}

	//node := s.Children().Nodes[number]
	//
	//out.set(strings.TrimSpace(s.FindNodes(node).Text()))
}

func after(s *Selection, in []string, out *result) {
	all := &result{}
	arrayMap(s, in, all)
	for name, value := range all.get().(map[string]interface{}) {
		if name == in[0] {
			out.set(value)
			return
		}
	}
}