package solution

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	nums1 := []int{3, 4, 5, 1, 2}
	r1 := 1
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	r2 := 0

	resu1 := findMin(nums1)
	resu2 := findMin(nums2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
