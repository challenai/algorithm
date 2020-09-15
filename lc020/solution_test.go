package solution

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	str1 := "()"
	r1 := true
	str2 := "()[]{}"
	r2 := true
	str3 := "(]"
	r3 := false
	str4 := "([)]"
	r4 := false
	str5 := "{[]}"
	r5 := true

	resu1 := isValid(str1)
	resu2 := isValid(str2)
	resu3 := isValid(str3)
	resu4 := isValid(str4)
	resu5 := isValid(str5)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 || resu5 != r5 {
		t.Fail()
	}
}
