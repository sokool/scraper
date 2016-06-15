package graph

import (
	query "github.com/PuerkitoBio/goquery"
)

type Node struct {
	Name      string //"nextList",
	Selector  string //"a[rel=next]",
	Neighbors string //"documents, nextList",
	Schema    string //"partial"

	url       string
	document  *query.Document
}

