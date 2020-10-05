package solution

import (
	"testing"
)

func TestCountPrimes(t *testing.T) {
	n1 := 10
	r1 := 4
	n2 := 0
	r2 := 0
	n3 := 1
	r3 := 0

	resu1 := countPrimes(n1)
	resu2 := countPrimes(n2)
	resu3 := countPrimes(n3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
