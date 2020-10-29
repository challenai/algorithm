package solution

import (
	"testing"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	A := []int{1, 2, 3, 4, 7}
	r := 4

	resu := numberOfArithmeticSlices(A)

	if r != resu {
		t.Fail()
	}
}
