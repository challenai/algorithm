package solution

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	resu := twoSum2(nums, target)
	if len(resu) != 2 && resu[0]+resu[1] != 1 {
		t.Fail()
	}
}
