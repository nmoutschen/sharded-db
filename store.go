package sharded

import "errors"

//Store represents the internal storage of the node
type Store struct {
	//Info contains the information about the current node
	info        NodeInfo
	clusterInfo ClusterInfo

	//data is the main data store for pieces of information
	data map[int]*Data
	/*nodeData maintains a copy of the pointers per other nodes that should
	have a copy of the pieces of information stored in this node. This is used
	when another node dies and need to stream data back.*/
	nodeData map[int]map[int]*Data
}

//NewStore creates a new data Store
func NewStore(clusterInfo ClusterInfo, nodeInfo NodeInfo) *Store {
	s := &Store{
		info:        nodeInfo,
		clusterInfo: clusterInfo,

		data:     make(map[int]*Data),
		nodeData: map[int]map[int]*Data{},
	}

	//Create all nodeData for each node
	for i := 0; i < s.clusterInfo.NumNodes; i++ {
		if i != s.info.ID {
			s.nodeData[i] = make(map[int]*Data)
		}
	}

	return s
}

//Delete deletes a piece of information
func (s *Store) Delete(dataID int) error {
	if _, ok := s.data[dataID]; !ok {
		return nil
	}

	//Delete in external node representations
	ni := NewNodeIterator(dataID, s.clusterInfo.NumNodes)
	for i := 0; i < s.clusterInfo.NumReplicas; i++ {
		n, err := ni.Next()
		if err != nil {
			return err
		}
		if n != s.info.ID {
			delete(s.nodeData[n], dataID)
		}
	}

	//Delete data in internal store
	delete(s.data, dataID)
	return nil
}

/*Load retrieves a piece of information

This is a purely internal operation and doesn't actually retrieve data from
distant nodes if this node does not have that piece of information.
*/
func (s *Store) Load(dataID int) (data *Data, ok bool) {
	data, ok = s.data[dataID]
	return data, ok
}

/*LoadNode returns a map containing all the known data for a given node

This is used when another node dies and need to stream the data to return to a
consistent state.
*/
func (s *Store) LoadNode(nodeID int) (data map[int]*Data, ok bool) {
	data, ok = s.nodeData[nodeID]
	return data, ok
}

/*Store stores or updates a piece of information

This is a purely internal operation and doesn't actually send the data to
distant nodes.
*/
func (s *Store) Store(data *Data) error {
	if _, ok := s.data[data.ID]; ok && s.data[data.ID].Time.After(data.Time) {
		return errors.New("Skip storing stale data")
	}

	//Store data in internal store
	s.data[data.ID] = data

	//Store in external node representations
	ni := NewNodeIterator(data.ID, s.clusterInfo.NumNodes)
	for i := 0; i < s.clusterInfo.NumReplicas; i++ {
		n, err := ni.Next()
		if err != nil {
			return err
		}
		if n != s.info.ID {
			s.nodeData[n][data.ID] = data
		}
	}

	return nil
}
