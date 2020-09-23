package solution

import (
	"testing"
)

func TestRestoreIpAddressses(t *testing.T) {
	s1 := "25525511135"
	r1 := []string{"255.255.11.135", "255.255.111.35"}
	s2 := "0000"
	r2 := []string{"0.0.0.0"}
	s3 := "010010"
	r3 := []string{"0.10.0.10", "0.100.1.0"}
	s4 := "101023"
	r4 := []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}

	resu1 := restoreIpAddresses(s1)
	resu2 := restoreIpAddresses(s2)
	resu3 := restoreIpAddresses(s3)
	resu4 := restoreIpAddresses(s4)

	if len(r1) != len(resu1) || len(r2) != len(resu2) || len(r3) != len(resu3) || len(r4) != len(resu4) {
		t.Fail()
	}
}
