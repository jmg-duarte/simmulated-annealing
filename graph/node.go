package graph

import (
	"fmt"
	"math"
)

// Node models a node in a Graph
type Node struct {
	ID          string
	XCoordinate float64
	YCoordinate float64
	//Neighbors   []*Node
}

// Equals checks if the given neighbor Node is the same as self
func (n Node) Equals(neighbour *Node) bool {
	return n.ID == neighbour.ID &&
		n.XCoordinate == neighbour.XCoordinate &&
		n.YCoordinate == neighbour.YCoordinate
}

// Dist calculates the Pythagorean distance between two Nodes
func (n Node) Dist(neighbour *Node) float64 {

	x := math.Abs(n.XCoordinate - neighbour.XCoordinate)
	y := math.Abs(n.YCoordinate - neighbour.YCoordinate)
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func (n Node) String() string {
	return fmt.Sprintf("Node: %s\nLocation: (%f, %f)\n", n.ID, n.XCoordinate, n.YCoordinate)
}
