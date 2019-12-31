package sharded

import (
	"testing"
)

func TestNodeIterator(t *testing.T) {
	testCases := []struct {
		Identifier int
		NumNodes int
		Expected []int
	}{
		{137, 7, []int{4, 1, 5}},
		{36677, 12, []int{5, 10, 8, 0, 4}},
		{0, 5, []int{0, 1, 2}},
	}

	for _, tc := range testCases {
		ni := NewNodeIterator(tc.Identifier, tc.NumNodes)
		for _, expected := range tc.Expected {
			val, _ := ni.Next()
			if val != expected {
				t.Errorf("ni.Next() == %d; want %d", val, expected)
			}
		}

		for i, expected := range tc.Expected {
			val, _ := ni.Get(i)
			if val != expected {
				t.Errorf("ni.Get(%d) == %d with precomputed values; want %d", i, val, expected)
			}
		}

		ni = NewNodeIterator(tc.Identifier, tc.NumNodes)
		for i, expected := range tc.Expected {
			val, _ := ni.Get(i)
			if val != expected {
				t.Errorf("ni.Get(%d) == %d without precomputed values; want %d", i, val, expected)
			}
		}

		ni = NewNodeIterator(tc.Identifier, tc.NumNodes)
		for i := len(tc.Expected)-1; i >= 0; i-- {
			val, _ := ni.Get(i)
			if val != tc.Expected[i] {
				t.Errorf("ni.Get(%d) == %d in reverse; want %d", i, val, tc.Expected[i])
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
			val, _ = sortedInsert(val, n)
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
