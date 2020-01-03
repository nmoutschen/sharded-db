package sharded

import "testing"

func TestNodeInfoAddr(t *testing.T) {
	ni := NodeInfo{
		Hostname: "127.0.0.1",
		Port:     8080,
	}

	if ni.Addr() != "127.0.0.1:8080" {
		t.Errorf("ni.Addr() == %s; want %s", ni.Addr(), "127.0.0.1:8080")
	}
}
