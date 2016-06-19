package storage

import ()
import "github.com/sokool/console"

type data struct {
	data []map[string]interface{}
}

func (this *data) Add(in map[string]interface{}) {
	this.data = append(this.data, in)
	console.Log(in)
}

func (this *data) Flush() {

}

func STRUCT(params  []string) Storage {
	return &data{
	}
}
