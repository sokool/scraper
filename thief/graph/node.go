package graph

type Node interface {
	Neighbors() []Node
}
