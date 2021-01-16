package solution

func LicenseKeyFormatting(s string, k int) string {
	if k < 1 || len(s) == 0 {
		return s
	}
	st := []byte{}
	rsu := ""
	counter := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '-' {
			if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'A' && s[i] <= 'Z') {
				st = append(st, s[i])
				continue
			}
			if s[i] >= 'a' && s[i] <= 'z' {
				st = append(st, s[i]-32)
			}
		}
	}
	for len(st) > 0 {
		if counter == k {
			rsu = "-" + rsu
			counter = 0
		}
		rsu = string(st[len(st)-1]) + rsu
		counter++
		st = st[:len(st)-1]
	}
	return rsu
}
