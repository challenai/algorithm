package solution

import (
	"testing"
)

func TestMinDistance(t *testing.T) {

	word11 := "horse"
	word12 := "ros"
	r1 := 3
	word21 := "intention"
	word22 := "execution"
	r2 := 5

	resu1 := minDistance(word11, word12)
	resu2 := minDistance(word21, word22)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
