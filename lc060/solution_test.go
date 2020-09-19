package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	n1 := 3
	k1 := 3
	r1 := "213"
	n2 := 4
	k2 := 9
	r2 := "2314"
	resu1 := getPermutation(n1, k1)
	resu2 := getPermutation(n2, k2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
