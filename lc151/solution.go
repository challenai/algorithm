package solution

func reverseWords(s string) string {
	current := ""
	resu := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(current) > 0 {
				resu = current + " " + resu
				current = ""
			}
			continue
		}
		current += string(s[i])
	}
	if len(current) > 0 {
		resu = current + " " + resu
	}
	if len(resu) == 0 {
		return ""
	}
	return resu[:len(resu)-1]
}
