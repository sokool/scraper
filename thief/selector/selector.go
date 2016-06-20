package selector

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func Parse(input string) (string, map[string][]string) {

	var selectorOut string
	var actionOut map[string][]string

	if strings.Index(input, "|") == -1 {
		input = input + "|text"
	}

	o := strings.Split(input, "|")
	selectorOut = strings.TrimSpace(o[0])
	if selectorOut == "" {
		return selectorOut, actionOut
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

	return selectorOut, actionOut
}

func Run(document *goquery.Document, in string) map[string]string {
	selector, actions := Parse(in)
	out := make(map[string]string)
	document.Find(selector).Each(func(number int, element *goquery.Selection) {
		for name, params := range actions {

			for name, value := range operations[name](element, params) {
				out[name] = value
			}
		}
	})
	return out
}
