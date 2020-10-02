package solution

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	s1 := "the sky is blue"
	r1 := "blue is sky the"
	s2 := "  hello world  "
	r2 := "world hello"
	s3 := "a good   example"
	r3 := "example good a"
	s4 := "  Bob    Loves  Alice   "
	r4 := "Alice Loves Bob"
	s5 := "Alice does not even like bob"
	r5 := "bob like even not does Alice"

	resu1 := reverseWords(s1)
	resu2 := reverseWords(s2)
	resu3 := reverseWords(s3)
	resu4 := reverseWords(s4)
	resu5 := reverseWords(s5)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 || resu5 != r5 {
		t.Fail()
	}
}
