package solution

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	r1 := "fl"
	strs2 := []string{"dog", "racecar", "car"}
	r2 := ""

	resu1 := longestCommonPrefix(strs1)
	resu2 := longestCommonPrefix(strs2)

	if resu1 != r1 && resu2 != r2 {
		t.Fail()
	}
}
