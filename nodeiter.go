package sharded

import (
	"errors"
)

/*NodeIterator represents a structure that can be used to iterate over a list
of nodes for a specific identifier without needing to pre-calculate all nodes.

This is useful to iterate over failed nodes in order, as node nodes are only
computed when necessary.
*/
type NodeIterator struct {
	Identifier int
	Position   int

	current     int
	numNodes    int
	sortedNodes []int
	nodes       []int
}

//NewNodeIterator creates a new node iterator
func NewNodeIterator(identifier, numNodes int) *NodeIterator {
	return &NodeIterator{
		Identifier: identifier,

		current:  identifier,
		numNodes: numNodes,
	}
}

/*Get returns a node at a specific position in the list of nodes for the
identifier.
*/
func (ni *NodeIterator) Get(pos int) (int, error) {
	if pos >= ni.numNodes {
		return 0, errors.New("Number of nodes exceeded")
	}

	//The value has not been precalculated
	if pos >= len(ni.nodes) {
		for i := ni.Position; i <= pos; i++ {
			ni.Next()
		}
	}

	return ni.nodes[pos], nil
}

//Next returns the next node in a NodeIterator
func (ni *NodeIterator) Next() (int, error) {
	if ni.Position >= ni.numNodes {
		return 0, errors.New("Number of nodes exceeded")
	}

	//The value has been precalculated
	if ni.Position < len(ni.nodes) {
		ni.Position++
		return ni.nodes[ni.Position], nil
	}

	node := ni.current % (ni.numNodes - ni.Position)
	ni.sortedNodes, node = sortedInsert(ni.sortedNodes, node)
	ni.current /= (ni.numNodes - ni.Position)
	ni.Position++

	ni.nodes = append(ni.nodes, node)
	return node, nil
}

/*sortedInsert insert a node in a slice of sorted nodes and increments the node
node based on the number of nodes that are smaller that itself.
*/
func sortedInsert(sortedNodes []int, node int) ([]int, int) {
	for i, sortedNode := range sortedNodes {
		if node >= sortedNode {
			node++
		} else {
			return append(sortedNodes[:i], append([]int{node}, sortedNodes[i:]...)...), node
		}
	}
	return append(sortedNodes, node), node
}
