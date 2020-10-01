package solution

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums1 := []int{2, 2, 1, 3, 3, 4, 4, 7, 8, 9, 9, 8, 7}
	nums2 := []int{4, 1, 2, 1, 2}
	r1 := 1
	r2 := 4
	resu1 := singleNumber(nums1)
	resu2 := singleNumber(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
