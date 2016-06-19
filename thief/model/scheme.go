package model

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sokool/scraper/thief/selector"
)

type scheme struct {
	Storage string
	Name    string
	Fields  map[string]*element
}

func (this *scheme) structure(doc *goquery.Document) (string, map[string]interface{}) {
	out := make(map[string]interface{})

	for name, field := range this.Fields {
		value := selector.Run(doc, field.Selector)
		if len(value) == 1 {
			out[name] = value[""]
		} else {
			out[name] = value
		}

	}

	return "x", out
}