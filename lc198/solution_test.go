package solution

import (
	"testing"
)

func TestRob(t *testing.T) {
	nums1 := []int{1, 2, 3, 1}
	r1 := 4
	nums2 := []int{2, 7, 9, 3, 1}
	r2 := 12

	resu1 := rob(nums1)
	resu2 := rob(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
