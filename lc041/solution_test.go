package solution

import (
	"testing"
)

func TestFristMissingPositive(t *testing.T) {

	nums1 := []int{1, 2, 0}
	r1 := 3
	nums2 := []int{3, 4, -1, 1}
	r2 := 2
	nums3 := []int{7, 8, 9, 11, 12}
	r3 := 1

	resu1 := firstMissingPositive(nums1)
	resu2 := firstMissingPositive(nums2)
	resu3 := firstMissingPositive(nums3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
