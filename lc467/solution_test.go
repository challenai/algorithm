package solution

import (
	"testing"
)

func TestFindSubstringInWraproundString(t *testing.T) {
	p1 := "a"
	r1 := 1
	p2 := "cac"
	r2 := 2
	p3 := "zab"
	r3 := 6

	rsu1 := findSubstringInWraproundString(p1)
	rsu2 := findSubstringInWraproundString(p2)
	rsu3 := findSubstringInWraproundString(p3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
