package solution

import (
	"testing"
)

func TestJump(t *testing.T) {

	nums := []int{2, 3, 1, 1, 4}
	r := 2
	resu := jump(nums)

	if resu != r {
		t.Fail()
	}
}
