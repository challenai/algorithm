package solution

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	r := 4

	resu := longestConsecutive(nums)

	if resu != r {
		t.Fail()
	}
}
