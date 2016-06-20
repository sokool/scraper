package selector

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

var operations = map[string]func(*goquery.Selection, []string) map[string]string{
	"attr" : attr,
	"map": arrayMap,
	"node": node,
	"after": after,
	"text": text,
	"html": html,
}

func text(selection *goquery.Selection, in []string) map[string]string {
	out := make(map[string]string)
	out[""] = strings.TrimSpace(selection.Text())

	return out
}

func attr(selection *goquery.Selection, in []string) map[string]string {
	out := make(map[string]string)
	a, _ := selection.Attr(in[0])
	out[""] = strings.TrimSpace(a)
	return out
}

func arrayMap(selection *goquery.Selection, in []string) map[string]string {
	out := make(map[string]string)

	nodes := selection.Children().Nodes
	if len(nodes) >= 2 {
		key := strings.TrimSpace(goquery.NewDocumentFromNode(selection.Children().Nodes[0]).Text())
		value := strings.TrimSpace(goquery.NewDocumentFromNode(selection.Children().Nodes[1]).Text())

		out[key] = value
	}

	return out
}

func html(selection *goquery.Selection, in []string) map[string]string {
	out := make(map[string]string)
	out[""], _ = selection.Html()

	return out
}

func node(selection *goquery.Selection, in []string) map[string]string {
	out := make(map[string]string)

	number, _ := strconv.Atoi(in[0])
	node := selection.Contents().Nodes[number]
	out[""] = strings.TrimSpace(selection.FindNodes(node).Text())

	return out
}

func after(selection *goquery.Selection, in []string) map[string]string {
	out := make(map[string]string)

	for name, value := range arrayMap(selection, in) {
		if name == in[0] {
			out[""] = value
			return out
		}
	}

	return out
}