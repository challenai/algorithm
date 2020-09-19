package solution

import (
	"testing"
)

func TestMerge(t *testing.T) {
	intervals1 := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	r1 := 3
	intervals2 := [][]int{
		[]int{1, 4},
		[]int{4, 5},
	}
	r2 := 1

	resu1 := merge(intervals1)
	resu2 := merge(intervals2)

	if len(resu1) != r1 || len(resu2) != r2 {
		t.Fail()
	}
}
