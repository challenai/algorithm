package solution

import "strings"

func restoreIpAddresses(s string) []string {
	s = strings.Trim(s, " ")
	resu := []string{}
	if len(s) < 4 || len(s) > 12 {
		return resu
	}
	dp := [][]string{}
	for i := 0; i < 4; i++ {
		dp = append(dp, []string{})
	}
	for i := 0; i < len(resu); i++ {
		print(resu[i], " || ")
	}
	return resu
}

func checkValidSegement(s string) bool {
	return len(s) == 1 || (len(s) == 2 && s[0] != '0') || (len(s) == 3 && (s[0] == '1' || (s[0] == '2' && ((s[1] < '5' && s[1] >= '0') || (s[1] == '5' && s[2] <= '5' && s[2] > '0')))))
}
