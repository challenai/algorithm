package solution

import (
	"testing"
)

func TestRemoveDulplicates(t *testing.T) {

	nums1 := []int{1, 1, 1, 2, 2, 3}
	r1 := 5
	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	r2 := 7
	resu1 := removeDulplicate(nums1)
	resu2 := removeDulplicate(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
