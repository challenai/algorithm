package solution

import (
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	S1 := "5F3Z-2e-9-w"
	K1 := 4
	r1 := "5F3Z-2E9W"
	S2 := "2-5g-3-J"
	K2 := 2
	r2 := "2-5G-3J"

	rsu1 := LicenseKeyFormatting(S1, K1)
	rsu2 := LicenseKeyFormatting(S2, K2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
