package solution

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-20, -50, -1, 0, 1, 2, -1, -4, -21, -16, -32, 200}
	r := [][]int{
		{-1, 0, 1},
		{-1, -1, 2},
	}

	resu := threeSum(nums)
	for i := 0; i < len(resu); i++ {
		for j := 0; j < len(resu[i]); j++ {
			print(resu[i][j], " ")
		}
		println()
	}

	if len(resu) != len(r) {
		t.Fail()
	}
}
