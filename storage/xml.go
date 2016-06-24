package storage

import (
	"fmt"
	"os"
	"github.com/clbanning/mxj"
)

type xml struct {
	name string
	list []map[string]interface{}
}

func (this *xml) Add(o map[string]interface{}) {
	this.list = append(this.list, o)

}

func (this *xml) Count() int {
	return len(this.list)
}


func (this *xml) Flush(name string) {
	handler, err := os.Create(name + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	fmt.Fprintf(handler, "<%s>\n", "objects")
	for _, data := range this.list {
		o := mxj.Map(data)
		bytes, _ := o.Xml()
		fmt.Fprintf(handler, "%s\n", string(bytes))
	}
	fmt.Fprintf(handler, "</%s>\n", "objects")

}

func XML(filename []string) Storage {
	return &xml{
		name: filename[0],
	}
}

