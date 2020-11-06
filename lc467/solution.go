package solution

func findSubstringInWraproundString(s string) int {
	if len(s) == 0 {
		return 0
	}
	exist := map[string]bool{}
	rsu := 1
	base := 1
	currentCombo := string(s[0])
	exist[string(s[0])] = false
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 || (s[i] == 'a' && s[i-1] == 'z') {
			for i < len(s) && (s[i] == s[i-1]+1 || (s[i] == 'a' && s[i-1] == 'z')) {
				currentCombo += string(s[i])
				base++
				if _, ok := exist[currentCombo]; !ok {
					rsu += base
				}
				i++
			}
		} else if _, ok := exist[string(s[i])]; !ok {
			rsu += base
			exist[string(s[i])] = false
		}
		if i < len(s) {
			base = 1
			currentCombo = string(s[i])
		}
	}
	return rsu
}
