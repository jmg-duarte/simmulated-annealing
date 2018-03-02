package main

import (
	"fmt"
	"log"
)

type NodeArray []*Node

// sort.Interface Methods
func (n NodeArray) Len() int {
	return len(n)
}

// Less compares the node "j" with "i" neighbourhood
// Returns "j->(i - 1)" < "j->(i + 1)"
// Being that "->" the represents the distance between the nodes
// "->" is also associative so "i->j" == "j->i"
func (n NodeArray) Less(i, j int) bool {
	l := n.Len()

	if l <= 0 {
		log.Fatal("Path.Less: path is empty")
	}

	if i > l || j > l {
		log.Fatal("Path.Less: cannot check out of bounds elements")
	}

	if i == j {
		return true
	}

	// Decrement the index to avoid 3 subtractions
	decIndex := i - 1
	// Declare the needed nodes for comparison
	var iNodeInc, iNodeDec, jNode *Node
	// Get the "i + 1" node in a circular path
	iNodeInc = n[(i+1)%l]
	// Get the "i - 1" node in a circular path
	if decIndex < 0 {
		iNodeDec = n[l%decIndex]
	} else {
		iNodeDec = n[decIndex]
	}
	// Get the "j" node
	jNode = n[j]

	jIncDist := jNode.Dist(iNodeInc)
	jDecDist := jNode.Dist(iNodeDec)

	return jDecDist < jIncDist
}

func (n NodeArray) Swap(i, j int) {
	l := n.Len()

	if l <= 0 {
		log.Fatal("Path.Swap: path is empty")
	}

	if i > l || j > l {
		log.Fatal("Path.Swap: cannot swap out of bounds elements")
	}

	if i == j {
		return
	}

	t := n[i]
	n[i] = n[j]
	n[j] = t
}

func (n NodeArray) Reverse() NodeArray {
	l := len(n)
	rev := make(NodeArray, l)
	copy(rev, n)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		rev[i] = n[j]
		rev[j] = n[i]
	}
	return rev
}

func (n NodeArray) String() string {
	return fmt.Sprint([]*Node(n))
}
