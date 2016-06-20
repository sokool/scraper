package graph

type Traveler interface {
	Root() Node
	Visit(Node) []Node //sprawdzić czy mozę byc unexported
	Done()
}
