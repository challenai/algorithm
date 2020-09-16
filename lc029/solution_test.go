package solution

import (
	"testing"
)

func TestDevide(t *testing.T) {
    devidend1 := 10
    devisor1 := 3
    devidend2 := 7
    devisor2 := -3
    r1 := 3
    r2 := -2

    resu1 := devide(devidend1, devisor1)
    resu2 := devide(devidend2, devisor2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
