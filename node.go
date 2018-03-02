package main

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
func (n Node) Equals(neighbor *Node) bool {
	return n.ID == neighbor.ID &&
		n.XCoordinate == neighbor.XCoordinate &&
		n.YCoordinate == neighbor.YCoordinate
}

// Dist calculates the Pythagorean distance between two Nodes
func (n Node) Dist(neighbor *Node) float64 {
	x := math.Abs(n.XCoordinate - neighbor.XCoordinate)
	y := math.Abs(n.YCoordinate - neighbor.YCoordinate)
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func (n Node) String() string {
	return fmt.Sprintf("Node: %s\nLocation: (%f, %f)\n", n.ID, n.XCoordinate, n.YCoordinate)
}
