package solution

import (
	"testing"
)

func TestFindNextGreaterElements(t *testing.T) {
	nums := []int{1, 2, 1}
	r := []int{2, -1, 2}
	rsu := nextGreaterElements(nums)
	for i := 0; i < len(r); i++ {
		if rsu[i] != r[i] {
			t.Fail()
		}
	}
}
