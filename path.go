package main

import (
	"bytes"
	"fmt"
)

// Path represents a path through all Nodes in a Graph
type Path struct {
	Nodes []*Node
}

// TotalDist returns the total path distance as if the path is a cycle
func (p Path) TotalDist() float64 {
	l := len(p.Nodes)
	if l <= 1 {
		return 0
	}

	total := p.Nodes[0].Dist(p.Nodes[1]) + p.Nodes[0].Dist(p.Nodes[l-1])
	if l == 2 {
		return total
	}

	for i := 2; i < l; i++ {
		total += p.Nodes[i-1].Dist(p.Nodes[i])
	}
	return total
}

func (p Path) String() string {
	var path bytes.Buffer
	l := len(p.Nodes) - 1
	for i := 0; i < l; i++ {
		path.WriteString(p.Nodes[i].ID)
		path.WriteString(" -> ")
	}
	path.WriteString(p.Nodes[l].ID)
	path.WriteString(fmt.Sprintf("\nPath total distance: %f", p.TotalDist()))
	return path.String()
}
