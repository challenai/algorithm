package solution

import (
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	r := 6

	resu := subsetsWithDup(nums)
	if len(resu) != r {
		t.Fail()
	}
}
