package selector

import (
	"strings"
	. "github.com/PuerkitoBio/goquery"
)

type command struct {
	selector string
	actions  []*action
}

type action struct {
	name   string
	params []string
}

func (self *action) call(s *Selection) *result {
	result := &result{}
	operations[self.name](s, self.params, result)

	return result
}

func (self *command) Run(d *Document) []interface{} {
	var out []interface{}
	for _, action := range self.actions {
		out = append(out, action.call(d.Find(self.selector)).get())
	}

	return out
}

func parse(input string) *command {
	var selectorOut string
	var actionOut []*action

	if strings.Index(input, "|") == -1 {
		input = input + "|text"
	}

	o := strings.Split(input, "|")
	selectorOut = strings.TrimSpace(o[0])
	if selectorOut == "" {
		return &command{
			selector: selectorOut,
			actions: actionOut,
		}
	}

	for _, element := range o[1:] {
		params := strings.Split(element, ":")
		options := []string{}
		for _, element := range params[1:] {
			element = strings.TrimSpace(element)
			if element == "" {
				continue
			}
			options = append(options, element)
		}

		actionOut = append(actionOut, &action{strings.TrimSpace(params[0]), options})
	}

	return &command{
		selector: selectorOut,
		actions: actionOut,
	}
}




