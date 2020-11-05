package solution

import (
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	n1 := 5
	r1 := 2
	n2 := 8
	r2 := 3
	n3 := 137
	r3 := 16

	rsu1 := arrangeCoins(n1)
	rsu2 := arrangeCoins(n2)
	rsu3 := arrangeCoins(n3)

	if r1 != rsu1 || r2 != rsu2 || rsu3 != r3 {
		t.Fail()
	}
}
