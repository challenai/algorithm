package solution

import (
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {
	n1 := 4
	edges1 := [][]int{{1, 0}, {1, 2}, {1, 3}}
	r1 := []int{1}
	n2 := 6
	edges2 := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	r2 := []int{3, 4}
	n3 := 1
	edges3 := [][]int{}
	r3 := []int{0}
	n4 := 2
	edges4 := [][]int{{0, 1}}
	r4 := []int{0, 1}

	rsu1 := findMinHeightTrees2(n1, edges1)
	rsu2 := findMinHeightTrees2(n2, edges2)
	rsu3 := findMinHeightTrees2(n3, edges3)
	rsu4 := findMinHeightTrees2(n4, edges4)

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) || len(rsu3) != len(r3) || len(rsu4) != len(r4) {
		t.Fail()
	}
}
