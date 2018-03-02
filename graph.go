package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// Graph models a graph
type Graph struct {
	Nodes    map[string]*Node
	NodeList []*Node
}

// AddNode checks if the given Node exists in the Graph
// If the Node doesn't exist the node is inserted
// Else the function returns
func (g *Graph) AddNode(n *Node) *Node {
	v, exists := g.Nodes[n.ID]
	if !exists {
		defer fmt.Printf("Node <%s> added to the graph\n", n.ID)
		// Adds node to the Graph map
		g.Nodes[n.ID] = n
		// Adds node to the Graph node list
		// Append will grow the array if needed
		g.NodeList = append(g.NodeList, n)
		// Returns the new node
		return n
	}
	fmt.Printf("Node <%s> already existed, no changes were made\nGraph Entry: <%s>\n", n.ID, n)
	// Returns the old node
	return v
}

// DeleteNode deletes a given node from the Graph
func (g *Graph) DeleteNode(n *Node) {
	delete(g.Nodes, n.ID)
}

// Size returns the number of nodes in the Graph
func (g Graph) Size() int {
	return len(g.NodeList)
}

// GenerateRandomPath creates a random path through all the nodes in the graph
// It switches nodes through a random index
// using the current system time as seed
func (g Graph) GenerateRandomPath() *Path {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	l := int32(len(g.NodeList))
	nodeStack := make([]*Node, l)
	copy(nodeStack, g.NodeList)

	var randomIndex int32
	var tempNode *Node
	for i := int32(0); i < l; i++ {
		randomIndex = randGen.Int31n(l)
		tempNode = nodeStack[i]
		nodeStack[i] = nodeStack[randomIndex]
		nodeStack[randomIndex] = tempNode
	}

	return &Path{nodeStack}
}

func (g Graph) String() string {
	return fmt.Sprintf("%s",
		func() string {
			var b bytes.Buffer
			for _, v := range g.Nodes {
				b.WriteString(v.String())
			}
			return b.String()
		}())
}
