package selector

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type Command struct {
	selector string
	actions  map[string][]string
}

func Parse(input string) *Command {
	var selectorOut string
	var actionOut map[string][]string

	if strings.Index(input, "|") == -1 {
		input = input + "|text"
	}

	o := strings.Split(input, "|")
	selectorOut = strings.TrimSpace(o[0])
	if selectorOut == "" {
		return &Command{
			selector: selectorOut,
			actions: actionOut,
		}
	}

	actionOut = make(map[string][]string)
	for _, action := range o[1:] {
		params := strings.Split(action, ":")
		options := []string{}
		for _, element := range params[1:] {
			element = strings.TrimSpace(element)
			if element == "" {
				continue
			}
			options = append(options, element)
		}

		actionOut[strings.TrimSpace(params[0])] = options
	}

	return &Command{
		selector: selectorOut,
		actions: actionOut,
	}
}

func (this *Command) Run(document *goquery.Document) map[string]string {
	out := make(map[string]string)
	document.Find(this.selector).Each(func(number int, element *goquery.Selection) {
		for name, params := range this.actions {
			for name, value := range operations[name](element, params) {
				out[name] = value
			}
		}
	})
	return out
}
