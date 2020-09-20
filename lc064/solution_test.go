package solution

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	grid1 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	r1 := 7

	resu1 := minPathSum(grid1)
	if resu1 != r1 {
		t.Fail()
	}
}
