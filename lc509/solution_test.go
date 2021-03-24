package solution

import (
	"testing"
)

func TestFibN(t *testing.T) {
	n1 := 2
	r1 := 1
	n2 := 4
	r2 := 3
	n3 := 7
	r3 := 13

	resu1 := Fib(n1)
	resu2 := Fib(n2)
	resu3 := Fib(n3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
