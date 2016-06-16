package graph

type Traveler interface {
	// Dla in pobiera DOCUMENT dla out ustawia URLE
	Visit(Node) []Node //findNeighbor?
	//OnLast(neighbor interface{}) //onLastNode?
	//OnFinish()
}
