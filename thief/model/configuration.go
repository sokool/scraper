package model

type configuration struct {
	Name   string`json:"name"`
	Url    string`json:"url"`
	Root   string
	Nodes  map[string]*node
	Schema *scheme

	data   map[string]interface{}
}

func (this *configuration) load() {
	//this.data = make(map[string]interface{})
	//this.Layout.onHit = this.found
}


