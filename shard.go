package sharded

//GetNode returns the first data node index for a specific identifier.
func GetNode(identifier, numNodes int) int {
	return identifier % numNodes
}

/*GetNodes returns the complete list of all available data node indexes for
a specific identifier.
*/
func GetNodes(identifier, numNodes, numReplicas int) []int {
	var nodes, sNodes []int
	value := identifier
	for modulo := numNodes; modulo > numNodes-numReplicas; modulo-- {
		node := value % modulo
		sNodes, node = SortedInsert(sNodes, node)

		nodes = append(nodes, node)
		value /= modulo
	}

	return nodes
}

//GetNodeN returns the Nth data node index for a specific identifier.
func GetNodeN(identifier, numNodes, pos int) int {
	var node int
	var sNodes []int
	value := identifier
	for modulo := numNodes; modulo >= numNodes-pos; modulo-- {
		node = value % modulo
		sNodes, node = SortedInsert(sNodes, node)
		value /= modulo
	}
	return node
}

/*SortedInsert insert a node in a slice of sorted nodes and increments the node
value based on the number of nodes that are smaller that itself.
*/
func SortedInsert(sortedNodes []int, node int) ([]int, int) {
	for i, sortedNode := range sortedNodes {
		if node >= sortedNode {
			node++
		} else {
			return append(sortedNodes[:i], append([]int{node}, sortedNodes[i:]...)...), node
		}
	}
	return append(sortedNodes, node), node
}
