package solution

import (
	"regexp"
	"strings"
)

func validIPAddress(IP string) string {
	if strings.Index(IP, ".") > 0 {
		segements := strings.Split(IP, ".")
		if len(segements) == 4 {
			for i := 0; i < len(segements); i++ {
				if !checkValidIPv4Segement(segements[i]) {
					return "Neither"
				}
			}
			return "IPv4"
		}
	}
	if strings.Index(IP, ":") > 0 {
		segements := strings.Split(IP, ":")
		if len(segements) == 8 {
			for i := 0; i < len(segements); i++ {
				if !checkValidIPv6Segement(segements[i]) {
					return "Neither"
				}
			}
			return "IPv6"
		}
	}
	return "Neither"
}

func checkValidIPv4Segement(s string) bool {
	ok, _ := regexp.MatchString("^25[0-5]$|^2[0-4]\\d$|^1\\d\\d$|^[1-9]\\d$|^\\d$", s)
	return ok
}

func checkValidIPv6Segement(s string) bool {
	ok, _ := regexp.MatchString("^[\\d|a-z|A-Z]{1,4}$", s)
	return ok
}
