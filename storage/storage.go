package storage

type Storage interface {
	Add(map[string]interface{})
	Flush()
}

var operations = map[string]func(...string) Storage{
	"xml" : XML,
	"json" :JSON,
	"struct": STRUCT,
}

func Get(name string) func(...string) Storage {
	return operations[name]
}