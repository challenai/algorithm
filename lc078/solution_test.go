package solution

import (
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	r := 8

	resu := subsets2(nums)
	// for i := 0; i < len(resu); i++ {
	// 	for j := 0; j < len(resu[i]); j++ {
	// 		print(resu[i][j])
	// 	}
	// 	println()
	// }

	if len(resu) != r {
		t.Fail()
	}
}
