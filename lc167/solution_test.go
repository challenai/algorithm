package solution

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	r1 := []int{1, 2}
	nums2 := []int{2, 3, 4}
	target2 := 6
	r2 := []int{1, 3}
	nums3 := []int{-1, 0}
	target3 := -1
	r3 := []int{1, 2}

	resu1 := twoSum(nums1, target1)
	resu2 := twoSum(nums2, target2)
	resu3 := twoSum(nums3, target3)

	if resu1[0] != r1[0] || resu2[1] != r2[1] || resu3[0] != r3[0] {
		t.Fail()
	}
}
