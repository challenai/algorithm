package solution

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	s1 := "egg"
	t1 := "add"
	r1 := true
	s2 := "foo"
	t2 := "bar"
	r2 := false
	s3 := "paper"
	t3 := "title"
	r3 := true

	rsu1 := isIsomorphic(s1, t1)
	rsu2 := isIsomorphic(s2, t2)
	rsu3 := isIsomorphic(s3, t3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
