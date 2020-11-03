package solution

func findAnagrams(s, p string) []int {
	rsu := []int{}
	if len(s) < len(p) {
		return rsu
	}
	hash := map[byte]bool{}
	current := map[byte]int{}
	for i := 0; i < len(p); i++ {
		hash[p[i]] = false
		if _, ok := current[p[i]]; ok {
			current[p[i]]++
		} else {
			current[p[i]] = 1
		}
	}
	for i := 0; i < len(p); i++ {
		// ignore if byte is useless
		if _, ok := hash[s[i]]; ok {
			if v, ok2 := current[s[i]]; ok2 {
				if v == 1 {
					delete(current, s[i])
					continue
				}
				current[s[i]]--
			} else {
				current[s[i]] = -1
			}
		}
	}
	if len(current) == 0 {
		rsu = append(rsu, 0)
	}
	for i := len(p); i < len(s); i++ {
		if s[i] == s[i-len(p)] {
			if len(current) == 0 {
				rsu = append(rsu, i-len(p)+1)
			}
			continue
		}
		println(i, i-len(p))
		if v, ok := current[s[i-len(p)]]; ok {
			if v == -1 {
				delete(current, s[i-len(p)])
			} else {
				current[s[i-len(p)]]++
			}
		} else {
			current[s[i-len(p)]] = 1
		}
		if v, ok := current[s[i]]; ok {
			if v == 1 {
				delete(current, s[i])
			} else {
				current[s[i]]--
			}
		} else {
			current[s[i]] = -1
		}
		if len(current) == 0 {
			rsu = append(rsu, i-len(p)+1)
		}
	}
	return rsu
}
