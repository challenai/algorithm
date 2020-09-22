package solution

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	str1 := "(()"
	r1 := 2
	str2 := ")()())"
	r2 := 4
	resu1 := longestValidParentheses(str1)
	resu2 := longestValidParentheses(str2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
