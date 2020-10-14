package solution

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tokens1 := []string{"2", "1", "+", "3", "*"}
	r1 := 9
	tokens2 := []string{"4", "13", "5", "/", "+"}
	r2 := 6
	tokens3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	r3 := 22

	resu1 := evalRPN(tokens1)
	resu2 := evalRPN(tokens2)
	resu3 := evalRPN(tokens3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
