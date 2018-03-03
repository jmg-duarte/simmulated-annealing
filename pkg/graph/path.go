package graph

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// Path represents a path through all Nodes in a Graph
type Path struct {
	Nodes NodeArray
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

	for i := 1; i < l; i++ {
		total += p.Nodes[i-1].Dist(p.Nodes[i])
	}
	return total
}

// GenerateNeighbour takes a Path p and applies the following changes
// Generates two random indexes and switches said elements
// Reverses all elements in between
func (p Path) GenerateNeighbour() *Path {
	l := len(p.Nodes)
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	nodes := make(NodeArray, l)
	copy(nodes, p.Nodes)

	randI := randGen.Intn(l)
	randJ := randGen.Intn(l)

	tempNode := nodes[randI]
	nodes[randI] = nodes[randJ]
	nodes[randJ] = tempNode
	if randI > randJ {
		copy(nodes[randJ:randI], nodes[randJ:randI].Reverse())
	} else {
		copy(nodes[randI:randJ], nodes[randI:randJ].Reverse())
	}
	return &Path{nodes}
}

// Stringer
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
