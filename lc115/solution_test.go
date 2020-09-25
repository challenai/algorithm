package solution

import (
	"testing"
)

func TestNumDistinct(t *testing.T) {
	s1 := "rabbbit"
	t1 := "rabbit"
	r1 := 3
	s2 := "babgbag"
	t2 := "bag"
	r2 := 5

	resu1 := numDistinct(s1, t1)
	resu2 := numDistinct(s2, t2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
