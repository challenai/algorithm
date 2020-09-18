package solution

import (
	"testing"
)

func TestRotate(t *testing.T) {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r1 := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	r2 := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	rotate(matrix1)
	rotate(matrix2)

	pass := true
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			if matrix1[i][j] != r1[i][j] {
				pass = false
				break
			}
		}
	}
	for i := 0; i < len(matrix2); i++ {
		for j := 0; j < len(matrix2[i]); j++ {
			if matrix2[i][j] != r2[i][j] {
				pass = false
				break
			}
		}
	}
	if !pass {
		t.Fail()
	}
}
