package solution

import (
	"testing"
)

func TestNumArray(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(nums)

	if numArray.SumRange(0, 2) != 1 {
		t.Fail()
	} // return 1 ((-2) + 0 + 3)
	if numArray.SumRange(2, 5) != -1 {
		t.Fail()
	} // return -1 3 + (-5( + 2 + )-1)()
	if numArray.SumRange(0, 5) != -3 {
		t.Fail()
	} // return -3 ((-2) + 0 + 3 + -5)( + 2 + -1)()
}
