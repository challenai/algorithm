package solution

import (
	"testing"
)

func TestFindLongestWord(t *testing.T) {
	s1 := "abpcplea"
	d1 := []string{"ale", "apple", "monkey", "plea"}
	r1 := "apple"
	s2 := "abpcplea"
	d2 := []string{"a", "b", "c"}
	r2 := "a"

	rsu1 := findLongestWord(s1, d1)
	rsu2 := findLongestWord(s2, d2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
