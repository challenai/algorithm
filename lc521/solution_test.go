package solution

import (
	"testing"
)

func TestFindLUSlength(t *testing.T) {

	a1 := "abc"
	b1 := "cdc"
	r1 := 3
	a2 := "aaa"
	b2 := "bbb"
	r2 := 3
	a3 := "aaa"
	b3 := "aaa"
	r3 := -1

	rsu1 := findLUSlength(a1, b1)
	rsu2 := findLUSlength(a2, b2)
	rsu3 := findLUSlength(a3, b3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
