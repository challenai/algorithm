package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	v11 := "1.01"
	v12 := "1.001"
	r1 := 0
	v21 := "1.0"
	v22 := "1.0.0"
	r2 := 0
	v31 := "0.1"
	v32 := "1.1"
	r3 := -1
	v41 := "1.0.1"
	v42 := "1"
	r4 := 1
	v51 := "7.5.2.4"
	v52 := "7.5.3"
	r5 := -1

	resu1 := compareVersion(v11, v12)
	resu2 := compareVersion(v21, v22)
	resu3 := compareVersion(v31, v32)
	resu4 := compareVersion(v41, v42)
	resu5 := compareVersion(v51, v52)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 || resu5 != r5 {
		t.Fail()
	}
}
