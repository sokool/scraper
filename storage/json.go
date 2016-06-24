package storage

import (
	"fmt"
	"os"
	"github.com/clbanning/mxj"
)

type json struct {
	name    string
	objects []map[string]interface{}
}

func (this *json) Add(o map[string]interface{}) {
	this.objects = append(this.objects, o)
}

func (this *json) Count() int {
	return len(this.objects)
}

func (this *json) Flush(name string) {
	handler, err := os.Create(name + ".json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	fmt.Fprintf(handler, "[")
	for _, object := range this.objects {
		bytes, _ := mxj.Map(object).Json()
		fmt.Fprintf(handler, "%s,\n", string(bytes))
	}
	fmt.Fprintf(handler, "]")

	//reset objects list?
	this.objects = []map[string]interface{}{}
}

func JSON(filename []string) Storage {
	return &json{
		name: filename[0],
	}
}


