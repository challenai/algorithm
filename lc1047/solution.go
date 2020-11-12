package solution

func removeDuplicates(s string) string {
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if len(st) > 0 && s[i] == st[len(st)-1] {
			st = st[:len(st)-1]
			continue
		}
		st = append(st, s[i])
	}
	rsu := ""
	for i := 0; i < len(st); i++ {
		rsu += string(st[i])
	}
	return rsu
}
