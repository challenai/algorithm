package solution

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	t1 := "nagaram"
	r1 := true

	s2 := "rat"
	t2 := "car"
	r2 := false

	resu1 := isAnagram(s1, t1)
	resu2 := isAnagram(s2, t2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
