package solution

import (
	"testing"
)

func TestNthSuperUglyNumber(t *testing.T) {
	n := 12
	primes := []int{2, 7, 13, 19}
	r := 32

	rsu := nthSuperUglyNumber(n, primes)

	if rsu != r {
		t.Fail()
	}
}
