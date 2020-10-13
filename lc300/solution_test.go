package solution

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	r := 4

	nums2 := []int{1, 12, 3, 2, 5, 8, 3, 8, 4, 7}
	r2 := 5

	resu1 := lengthOfLIS(nums)
	resu2 := lengthOfLIS(nums2)

	if resu1 != r || resu2 != r2 {
		t.Fail()
	}
}
