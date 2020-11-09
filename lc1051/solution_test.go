package solution

import (
	"testing"
)

func TestHeightChecker(t *testing.T) {
	nums1 := []int{1, 1, 4, 2, 1, 3}
	r1 := 3
	nums2 := []int{5, 1, 2, 3, 4}
	r2 := 5
	nums3 := []int{1, 2, 3, 4, 5}
	r3 := 0

	rsu1 := heightChecker(nums1)
	rsu2 := heightChecker(nums2)
	rsu3 := heightChecker(nums3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
