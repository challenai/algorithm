package solution

import (
	"testing"
)

func TestAddStrings(t *testing.T) {
	nums11 := "123412341234"
	nums12 := "432143214321"
	r1 := "555555555555"
	nums21 := "684447545854"
	nums22 := "4423344343"
	r2 := "688870890197"

	resu1 := addStrings(nums11, nums12)
	resu2 := addStrings(nums21, nums22)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
