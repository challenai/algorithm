package solution

import (
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	r := 4
	resu := maximalSquare(matrix)

	if resu != r {
		t.Fail()
	}
}
