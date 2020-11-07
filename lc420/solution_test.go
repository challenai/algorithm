package solution

import (
	"testing"
)

func TestStrongPasswordChecker(t *testing.T) {
	passwd1 := "fjfe.aaafdsh."
	r1 := 1
	passwd2 := "...aa1..a.A"
	r2 := 0
	passwd3 := "IFEESFK342"
	r3 := 0
	passwd4 := "abc"
	r4 := 0

	rsu1 := strongPasswordChecker(passwd1)
	rsu2 := strongPasswordChecker(passwd2)
	rsu3 := strongPasswordChecker(passwd3)
	rsu4 := strongPasswordChecker(passwd4)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
