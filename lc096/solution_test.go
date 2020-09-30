package solution

import (
	"testing"
)

func TestNumTrees(t *testing.T) {
	n1 := 3
	r1 := 5
	n2 := 4
	r2 := 9

	resu1 := numTrees(n1)
	resu2 := numTrees(n2)

	if r1 != resu1 || resu2 != r2 {
		t.Fail()
	}
}
