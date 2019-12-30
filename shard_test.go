package sharded

import (
	"testing"
)

func TestGetNode(t *testing.T) {
	testCases := []struct {
		Identifier int
		NumNodes   int
		Expected   int
	}{
		{137, 7, 4},
		{1243, 4, 3},
		{1364, 26, 12},
		{8346, 10, 6},
		{7295, 27, 5},
		{7169, 25, 19},
		{596, 25, 21},
		{2620, 7, 2},
		{6710, 28, 18},
		{7577, 29, 8},
		{8931, 29, 28},
		{8968, 17, 9},
		{1659, 21, 0},
		{6104, 26, 20},
		{1557, 23, 16},
		{6805, 22, 7},
		{7162, 21, 1},
		{8247, 2, 1},
		{6447, 14, 7},
		{4768, 9, 7},
		{6943, 28, 27},
		{3463, 9, 7},
		{5661, 11, 7},
		{426, 14, 6},
		{1618, 17, 3},
		{8283, 14, 9},
		{8202, 20, 2},
		{932, 18, 14},
		{8847, 6, 3},
		{9810, 2, 0},
		{6505, 12, 1},
		{2942, 10, 2},
		{931, 18, 13},
		{1603, 28, 7},
		{7288, 27, 25},
		{2541, 11, 0},
		{3822, 3, 0},
		{7479, 8, 7},
		{3522, 16, 2},
		{705, 2, 1},
		{6966, 26, 24},
		{9578, 7, 2},
		{6752, 27, 2},
		{9255, 19, 2},
		{3619, 10, 9},
		{3631, 12, 7},
		{7200, 10, 0},
		{2323, 3, 1},
		{1964, 28, 4},
		{1445, 2, 1},
		{3068, 23, 9},
	}

	for _, tc := range testCases {
		val := GetNode(tc.Identifier, tc.NumNodes)
		if val != tc.Expected {
			t.Errorf("GetNode(%d, %d) == %d; want %d", tc.Identifier, tc.NumNodes, val, tc.Expected)
		}
	}
}

func TestGetNodes(t *testing.T) {
	testCases := []struct {
		Identifier  int
		NumNodes    int
		NumReplicas int
		Expected    []int
	}{
		{137, 7, 3, []int{4, 1, 5}},
		{36677, 12, 5, []int{5, 10, 8, 0, 4}},
		{0, 5, 3, []int{0, 1, 2}},
	}

	for _, tc := range testCases {
		val := GetNodes(tc.Identifier, tc.NumNodes, tc.NumReplicas)
		if len(val) != len(tc.Expected) {
			t.Errorf("len(GetNodes(%d, %d, %d)) == %d; want %d", tc.Identifier, tc.NumNodes, tc.NumReplicas, len(val), len(tc.Expected))
		}
		for i := range tc.Expected {
			if val[i] != tc.Expected[i] {
				t.Errorf("GetNodes(%d, %d, %d)[%d] == %v; want %v", tc.Identifier, tc.NumNodes, tc.NumReplicas, i, val[i], tc.Expected[i])
			}
		}
	}
}

func TestGetNodeN(t *testing.T) {
	testCases := []struct {
		Identifier int
		NumNodes   int
		Expected   []int
	}{
		{137, 7, []int{4, 1, 5}},
		{36677, 12, []int{5, 10, 8, 0, 4}},
		{0, 5, []int{0, 1, 2}},
	}

	for _, tc := range testCases {
		for i, expected := range tc.Expected {
			val := GetNodeN(tc.Identifier, tc.NumNodes, i)
			if val != expected {
				t.Errorf("GetNodeN(%d, %d, %d) == %d; want %d", tc.Identifier, tc.NumNodes, i, val, expected)
			}
		}
	}
}

func TestSortedInsert(t *testing.T) {
	testCases := []struct {
		Input    []int
		Expected []int
	}{
		{[]int{4, 1, 3}, []int{1, 4, 5}},
		{[]int{0, 0, 0}, []int{0, 1, 2}},
	}

	for _, tc := range testCases {
		var val []int
		for _, n := range tc.Input {
			val, _ = SortedInsert(val, n)
		}
		if len(val) != len(tc.Expected) {
			t.Errorf("len(sortedNodes) == %d; want %d", len(val), len(tc.Expected))
		}
		for i := range tc.Expected {
			if val[i] != tc.Expected[i] {
				t.Errorf("sortedNodes[%d] == %d; want %d", i, val[i], tc.Expected[i])
			}
		}
	}
}
