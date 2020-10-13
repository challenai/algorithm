package solution

import (
	"testing"
)

func TestRemoveInvalidParentheses(t *testing.T) {
	s1 := "()())()"
	r1 := []string{"()()()", "(())()"}
	s2 := "(a)())()"
	r2 := []string{"(a)()()", "(a())()"}
	s3 := ")("
	r3 := []string{""}

	resu1 := removeInvalidParentheses(s1)
	resu2 := removeInvalidParentheses(s2)
	resu3 := removeInvalidParentheses(s3)

	if len(r1) != len(resu1) || len(r2) != len(resu2) || len(r3) != len(resu3) {
		t.Fail()
	}
}
