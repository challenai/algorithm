package solution

import (
	"testing"
)

func TestSetZeroes(t *testing.T) {
	input1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	input2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	r1 := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	r2 := [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0},
	}

	setZeroes(input1)
	setZeroes(input2)
	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r1[0]); j++ {
			if input1[i][j] != r1[i][j] {
				t.Fail()
				return
			}
		}
	}

	for i := 0; i < len(r2); i++ {
		for j := 0; j < len(r2[0]); j++ {
			if input2[i][j] != r2[i][j] {
				t.Fail()
				return
			}
		}
	}
}
