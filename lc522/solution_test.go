package solution

import (
	"testing"
)

func TestFindLUSlength(t *testing.T) {

	strs1 := []string{"jj", "aba", "cdc", "eae", "fd", "fd"}
	r1 := 3
	strs2 := []string{"a", "aaa", "aaa", "aa"}
	r2 := -1

	rsu1 := findLUSlength(strs1)
	rsu2 := findLUSlength(strs2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
