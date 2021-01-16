package solution

import (
	"testing"
)

func TestMedianSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	r := []int{1, -1, -1, 3, 5, 6}

	resu := MedianSlidingWindow(nums, k)

	for i := 0; i < len(r); i++ {
		if resu[i] != r[i] {
			t.Fail()
		}
	}
}
