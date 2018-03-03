package graph

import (
	"fmt"
)

// NodeArray represents an array of nodes
type NodeArray []*Node

// Reverse a NodeArray ordering
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
