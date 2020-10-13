package solution

import (
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	s1 := "2-1-1"
	r1 := map[int]bool{
		0: false,
		2: false,
	}
	s2 := "2*3-4*5"
	r2 := map[int]bool{
		-34: false,
		-14: false,
		-10: false,
		10:  false,
	}
	s3 := "23-4"
	r3 := map[int]bool{
		19: false,
	}

	resu1 := diffWaysToCompute(s1)
	resu2 := diffWaysToCompute(s2)
	resu3 := diffWaysToCompute(s3)

	for i := 0; i < len(resu1); i++ {
		if _, ok := r1[resu1[i]]; !ok {
			t.Fail()
		}
	}
	for i := 0; i < len(resu2); i++ {
		if _, ok := r2[resu2[i]]; !ok {
			t.Fail()
		}
	}
	for i := 0; i < len(resu3); i++ {
		if _, ok := r3[resu3[i]]; !ok {
			t.Fail()
		}
	}
}
