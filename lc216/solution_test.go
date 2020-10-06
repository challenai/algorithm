package solution

import (
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	k1 := 3
	n1 := 7
	r1 := [][]int{{1, 2, 4}}
	k2 := 3
	n2 := 9
	r2 := [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}
	k3 := 3
	n3 := 2
	r3 := [][]int{}
	k4 := 9
	n4 := 45
	r4 := [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}

	rsu1 := combinationSum3(k1, n1)
	rsu2 := combinationSum3(k2, n2)
	rsu3 := combinationSum3(k3, n3)
	rsu4 := combinationSum3(k4, n4)

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) || len(rsu3) != len(r3) || len(rsu4) != len(r4) {
		t.Fail()
	}
}
