package solution

import (
	"testing"
)

func TestFindRightInterval(t *testing.T) {
	intervals1 := [][]int{{1, 2}}
	r1 := []int{-1}
	intervals2 := [][]int{{3, 4}, {2, 3}, {1, 2}}
	r2 := []int{-1, 0, 1}
	intervals3 := [][]int{{1, 4}, {2, 3}, {3, 4}}
	r3 := []int{-1, 2, -1}

	rsu1 := findRightInterval(intervals1)
	rsu2 := findRightInterval(intervals2)
	rsu3 := findRightInterval(intervals3)

	if rsu1[0] != r1[0] || rsu2[1] != r2[1] || rsu3[1] != r3[1] {
		t.Fail()
	}
}
