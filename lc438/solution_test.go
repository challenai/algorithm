package solution

import (
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s1 := "cbaebabacd"
	p1 := "abc"
	r1 := []int{0, 6}
	s2 := "abab"
	p2 := "ab"
	r2 := []int{0, 1, 2}

	rsu1 := findAnagrams(s1, p1)
	rsu2 := findAnagrams(s2, p2)

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) {
		t.Fail()
	}
}
