package solution

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	int1 := 121
	r1 := true
	int2 := -121
	r2 := false
	int3 := 10
	r3 := false

	resu1 := isPalindrome(int1)
	resu2 := isPalindrome(int2)
	resu3 := isPalindrome(int3)
	println(resu1)
	println(resu2)
	println(resu3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
