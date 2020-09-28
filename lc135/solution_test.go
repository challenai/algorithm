package solution

import (
	"testing"
)

func TestCandy(t *testing.T) {
	nums1 := []int{1, 0, 2}
	r1 := 5
	nums2 := []int{1, 2, 2}
	r2 := 4
	resu1 := candy(nums1)
	resu2 := candy(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
