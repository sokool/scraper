package selector

import (
	. "github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"github.com/sokool/console"
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
	out.set(strings.TrimSpace(s.First().Text()))
}

func attr(s *Selection, in []string, out *result) {
	switch len(in) {
	case 2:
		data := make(map[string]string)
		s.Each(func(index int, item *Selection) {
			attribute, _ := item.First().Attr(in[0])
			data[strconv.Itoa(index)] = attribute
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
	data := make(map[string]string)
	s.Each(func(idx int, itm *Selection) {
		key := strings.TrimSpace(NewDocumentFromNode(itm.Children().Nodes[0]).Text())
		value := strings.TrimSpace(NewDocumentFromNode(itm.Children().Nodes[1]).Text())

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
	for name, value := range all.get().(map[string]string) {
		if name == in[0] {
			out.set(value)
			return
		}
	}
}