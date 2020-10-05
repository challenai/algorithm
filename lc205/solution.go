package solution

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	exist := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := exist[s[i]]; ok {
			exist[s[i]] = append(exist[s[i]], i)
		} else {
			exist[s[i]] = []int{i}
		}
	}
	for i := 0; i < len(s); i++ {
		if temp, ok := exist[s[i]]; ok {
			for j := 1; j < len(temp); j++ {
				if t[temp[j]] != t[temp[0]] {
					return false
				}
			}
			delete(exist, s[i])
		}
	}
	return true
}
