package solution

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums1 := []int{1, 1, 2}
	r1 := 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r2 := 5
	resu1 := removeDuplicates(nums1)
	resu2 := removeDuplicates(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
