package solution

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	r := [][]int{
		{-1, 0, 0, 1},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
	}

	resu := fourSum(nums, target)
	// for i := 0; i < len(resu); i++ {
	// 	for j := 0; j < len(resu[i]); j++ {
	// 		print(resu[i][j], " ")
	// 	}
	// 	println()
	// }

	if len(resu) != len(r) {
		t.Fail()
	}
}
