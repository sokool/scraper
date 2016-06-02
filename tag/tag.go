package tag

import (
	"reflect"
	"strings"
)

type Tag struct {
	namespaces map[string]*Namespace
	reflection reflect.Type
}

type Namespace struct {
	name    string
	nDisc   string
	aDisc   string
	actions map[string][]*action
}

type action struct {
	name   string
	params []string
}

func New(structure interface{}) *Tag {
	tag := &Tag{
		reflection: reflect.TypeOf(structure),
		namespaces: make(map[string]*Namespace),
	}

	//tag.buildStructure()
	return tag
}

func (t *Tag) AddNamespace(name, nD, aD string) *Namespace {
	namespace := &Namespace{
		name: name,
		nDisc: nD,
		aDisc: aD,
		actions: make(map[string][]*action),
	}

	for i := 0; i < t.reflection.NumField(); i++ {
		field := t.reflection.Field(i)
		tag := field.Tag.Get(name)
		if tag != "" {
			t.parse(field.Name, tag, namespace)
		}
	}

	t.namespaces[name] = namespace

	return namespace
}

func (t *Tag) parse(field, tag string, namespace *Namespace) {
	to := strings.Index(tag, namespace.nDisc)
	if to == -1 {
		out := strings.TrimSpace(tag)
		mPosition := strings.Index(out, " ")
		if (mPosition == -1) {
			namespace.addAction(field, &action{name:out})
			return
		}
		method := strings.TrimSpace(out[:mPosition])

		namespace.addAction(field, &action{name:method})
		t.parse(field, out[mPosition:], namespace)

		return
	}

	method := strings.TrimSpace(tag[:to])
	out := strings.TrimSpace(tag[to + len(namespace.nDisc):])
	f := strings.Index(out, namespace.nDisc)

	if f == -1 {
		params := strings.Split(out, namespace.aDisc)
		namespace.addAction(field, &action{name:method, params: params})

		//fmt.Printf("===========%s===========================\nMETHOD: %s\nPARAMS: %s\n", tag, method, params)
		return
	}

	o2 := strings.TrimSpace(out[:f])
	methodBegin := strings.LastIndex(o2, " ")
	params := strings.Split(o2[:methodBegin], namespace.aDisc)
	input := strings.TrimSpace(out[methodBegin:])
	//fmt.Printf("===========%s===========================\nMETHOD: %s\nPARAMS: %s\n", tag, method, params)
	namespace.addAction(field, &action{name:method, params: params})
	t.parse(field, input, namespace)

}

func (t *Tag) Namespace(name string) *Namespace {
	namespace, ok := t.namespaces[name];
	if !ok {
		return nil
	}

	return namespace
}

func (t *Tag) Get(namespace, field string, n int) (string, []string) {
	return t.Namespace(namespace).Get(field, n)
}

func (t *Tag) GetFunc(f func(field string, n int, value string)) {
	//for name, actions := range t.filed {
	//	for idx, value := range actions {
	//		f(name, idx, value)
	//	}
	//}
}

func (n *Namespace) addAction(field string, action *action) {
	n.actions[field] = append(n.actions[field], action)
}

func (n *Namespace) Get(action string, i int) (string, []string) {
	a := n.actions[action][i]
	return a.name, a.params
}

