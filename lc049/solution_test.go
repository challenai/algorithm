package solution

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	r1 := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	strs2 := []string{""}
	r2 := [][]string{{""}}
	strs3 := []string{"a"}
	r3 := [][]string{{"a"}}

	resu1 := groupAnagrams(strs1)
	resu2 := groupAnagrams(strs2)
	resu3 := groupAnagrams(strs3)

	if len(resu1) != len(r1) || len(resu2) != len(r2) || len(resu3) != len(r3) {
		t.Fail()
	}
}
