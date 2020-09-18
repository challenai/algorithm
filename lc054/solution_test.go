package solution

import (
	"testing"
)

func TestSpiralOrder(t *testing.T) {

	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r1 := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	r2 := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	resu1 := spiralOrder(matrix1)
	resu2 := spiralOrder(matrix2)

	pass := true
	for i := 0; i < len(r1); i++ {
		if resu1[i] != r1[i] {
			pass = false
			break
		}
	}

	for i := 0; i < len(r2); i++ {
		if resu2[i] != r2[i] {
			pass = false
			break
		}
	}

	if !pass {
		t.Fail()
	}
}
