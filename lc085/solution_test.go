package solution

import (
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	r1 := 6
	matrix2 := [][]byte{{'0'}, {'1'}}
	r2 := 1
	matrix3 := [][]byte{{'1', '1'}}
	r3 := 2

	resu1 := maximalRectangle(matrix1)
	resu2 := maximalRectangle(matrix2)
	resu3 := maximalRectangle(matrix3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
