package solution

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 7
	r := 2
	nums2 := []int{2, 4, 4, 3, 54, 3, 54, 54, 52}
	s2 := 543537
	r2 := 0

	resu := minSubArrayLen(s, nums)
	resu2 := minSubArrayLen(s2, nums2)

	if resu != r || resu2 != r2 {
		t.Fail()
	}
}
