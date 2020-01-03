package sharded

import "fmt"

//ClusterInfo represents the overall configuration of the cluster
type ClusterInfo struct {
	NumNodes    int
	NumReplicas int
	Nodes       []NodeInfo
}

//NodeInfo represents information relative to a node in the cluster
type NodeInfo struct {
	ID       int
	Hostname string
	Port     int
}

//Addr returns a hostname:port representation of the node
func (ni *NodeInfo) Addr() string {
	return fmt.Sprintf("%s:%d", ni.Hostname, ni.Port)
}
