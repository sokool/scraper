package model

import (
	"strings"
)

type node struct {
	Selector  string
	Neighbors []string

	url       string
}

func (this *node) hasNeighbors() bool {
	return len(this.Neighbors) > 0
}

func (this *node) hasSelector() bool {
	return strings.TrimSpace(this.Selector) != ""
}

func (this *node) copy() *node {
	node := *this
	return &node
}

func (this *node) setUrl(link string) *node {
	this.url = link
	return this
}
