package model

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sokool/scraper/thief/selector"
	"github.com/sokool/scraper/thief/filter"
)

type scheme struct {
	Storage  string
	Fields   map[string]*field

	selector *selector.Selector
}

func (this *scheme) scrape(d *goquery.Document) Object {
	if this.selector == nil {
		this.selector = selector.New()
		for name, field := range this.Fields {
			this.selector.Append(name, field.Selector)
		}
	}

	object := this.selector.Execute(d)
	this.filter(object)
	return object
}

func (this *scheme) filter(o Object) {
	var err error
	for name, field := range this.Fields {
		if field.Filters == "" {
			continue
		}
		item, ok := o[name]
		if !ok {
			continue
		}

		switch item.(type) {
		case string:
			o[name], err = filter.Run(field.Filters, item.(string))
			if err != nil {
				panic(err)
			}
			break
		}

	}
}