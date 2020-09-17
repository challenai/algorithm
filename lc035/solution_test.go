package solution

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	r1 := 2
	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	r2 := 1
	nums3 := []int{1, 3, 5, 6}
	target3 := 7
	r3 := 4
	nums4 := []int{1, 3, 5, 6}
	target4 := 0
	r4 := 0

	resu1 := searchInsert(nums1, target1)
	resu2 := searchInsert(nums2, target2)
	resu3 := searchInsert(nums3, target3)
	resu4 := searchInsert(nums4, target4)

	println(resu1)
	println(resu2)
	println(resu3)
	println(resu4)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 {
		t.Fail()
	}
}
