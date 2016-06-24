package selector

import (
	"github.com/PuerkitoBio/goquery"
)

type Selector struct {
	commands map[string]*command
}

func (self *Selector) Append(name, selector string) *Selector {
	self.commands[name] = parse(selector)
	return self
}

func (self *Selector) Execute(d *goquery.Document) map[string]interface{} {
	output := make(map[string]interface{})
	for name, cmd := range self.commands {
		result := cmd.Run(d)
		if len(result) == 1 {
			output[name] = result[0]
		} else {
			output[name] = result
		}

	}
	return output
}

func New() *Selector {
	return &Selector{
		commands: make(map[string]*command),
	}
}
