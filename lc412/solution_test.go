package solution

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	n := 15
	r := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}
	resu := fizzBuzz(n)

	if len(resu) != len(r) {
		t.Fail()
	}
	for i := 0; i < len(resu); i++ {
		if resu[i] != r[i] {
			t.Fail()
		}
	}
}
