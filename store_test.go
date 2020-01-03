package sharded

import (
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	clusterInfo := ClusterInfo{7, 3, []NodeInfo{}}
	nodeInfo := NodeInfo{1, "127.0.0.1", 8080}
	s := NewStore(clusterInfo, nodeInfo)
	data := NewData(137, time.Now(), "TestStore")

	// Store data
	if err := s.Store(data); err != nil {
		t.Errorf("s.Store() returned an error: %s", err.Error())
	}

	//Access with s.Load()
	if d, ok := s.Load(data.ID); !ok || d != data {
		t.Errorf("s.Load() did not return the right data")
	}

	//Check for s.LoadNode()
	nodeIter := NewNodeIterator(data.ID, clusterInfo.NumNodes)
	for i := 0; i < clusterInfo.NumReplicas; i++ {
		n, _ := nodeIter.Next()
		if n != nodeInfo.ID {
			ds, ok := s.LoadNode(n)
			if !ok {
				t.Errorf("s.LoadNode() returned ok == false")
			}
			if len(ds) == 0 {
				t.Errorf("s.LoadNode() does not contain data after s.Load()")
				continue
			}
			if d, ok := ds[data.ID]; !ok || d != data {
				t.Errorf("s.LoadNode() did not return the right data")
			}
		}
	}

	//Delete data
	if err := s.Delete(data.ID); err != nil {
		t.Errorf("s.Delete() returned an error: %s", err.Error())
	}

	//Access with s.Load()
	if _, ok := s.Load(data.ID); ok {
		t.Errorf("s.Load() returned data after s.Delete()")
	}

	//Check for s.LoadNode()
	nodeIter = NewNodeIterator(data.ID, clusterInfo.NumNodes)
	for i := 0; i < clusterInfo.NumReplicas; i++ {
		n, _ := nodeIter.Next()
		if n != nodeInfo.ID {
			ds, ok := s.LoadNode(n)
			if !ok {
				t.Errorf("s.LoadNode() returned ok == false")
			}
			if len(ds) != 0 {
				t.Errorf("s.LoadNode() contains data after s.Delete()")
			}
		}
	}

	if err := s.Delete(data.ID); err != nil {
		t.Errorf("s.Delete() returned an error after second call: %s", err.Error())
	}
}
