package solution

import (
	"testing"
)

func TestValidIPAddress(t *testing.T) {
	IP1 := "172.16.254.1"
	r1 := "IPv4"
	IP2 := "2001:0db8:85a3:0:0:8A2E:0370:7334"
	r2 := "IPv6"
	IP3 := "256.256.256.256"
	r3 := "Neither"
	IP4 := "2001:0db8:85a3:0:0:8A2E:0370:733:"
	r4 := "Neither"
	IP5 := "1e1.4.5.6"
	r5 := "Neither"

	rsu1 := validIPAddress(IP1)
	rsu2 := validIPAddress(IP2)
	rsu3 := validIPAddress(IP3)
	rsu4 := validIPAddress(IP4)
	rsu5 := validIPAddress(IP5)
	println(rsu1)
	println(rsu2)
	println(rsu3)
	println(rsu4)
	println(rsu5)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 || rsu5 != r5 {
		t.Fail()
	}
}
