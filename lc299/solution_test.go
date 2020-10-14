package solution

import (
	"testing"
)

func TestGetHint(t *testing.T) {
	secret1 := "1807"
	guess1 := "7810"
	r1 := "1A3B"
	secret2 := "1123"
	r2 := "1A1B"
	guess2 := "0111"
	secret3 := "1"
	r3 := "0A0B"
	guess3 := "0"
	secret4 := "1"
	guess4 := "1"
	r4 := "1A0B"

	rsu1 := getHint(secret1, guess1)
	rsu2 := getHint(secret2, guess2)
	rsu3 := getHint(secret3, guess3)
	rsu4 := getHint(secret4, guess4)

	if r1 != rsu1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
