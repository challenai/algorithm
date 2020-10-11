package solution

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	t1 := 5
	t2 := 20
	r1 := true
	r2 := false

	resu1 := searchMatrix(matrix, t1)
	resu2 := searchMatrix(matrix, t2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
