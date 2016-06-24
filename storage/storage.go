package storage

type Storage interface {
	Add(map[string]interface{})
	Flush(name string)
	Count() int
}

var operations = map[string]func([]string) Storage{
	"xml" : XML,
	"json" :JSON,
}

func Get(name string, params []string) Storage {
	return operations[name](params)
}