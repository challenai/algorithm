package solution

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	nums1 := []int{3, 5, 7, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2}
	r1 := 1
	nums2 := []int{2, 2, 2, 0, 1}
	r2 := 0
	nums3 := []int{3, 4, 4, 5, 6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 1, 2}
	r3 := 1

	resu1 := findMin(nums1)
	resu2 := findMin(nums2)
	resu3 := findMin(nums3)
	// println(resu1)
	// println(resu2)
	// println(resu3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
