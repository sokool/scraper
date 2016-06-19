package model

type node struct {
	Selector  string
	Neighbors []string
	Schema    string

	url       string
}
