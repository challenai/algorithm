package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	matrix1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target1 := 3
	r1 := true
	matrix2 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target2 := 13
	r2 := false

	resu1 := searchMatrix(matrix1, target1)
	resu2 := searchMatrix(matrix2, target2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
