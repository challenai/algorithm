package solution

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	grid1 := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	r1 := 2
	grid2 := [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	r2 := 3

	resu1 := uniquePathsWithObstacles(grid1)
	resu2 := uniquePathsWithObstacles(grid2)
	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
