package solution

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	r := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	resu := generateParenthesis(n)

	if len(resu) != len(r) {
		t.Fail()
	}
}
