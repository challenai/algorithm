package solution

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	var left, rsu int
	hash := map[byte]int{}
	left = 0
	rsu = 1
	currentMaxKey := s[0]
	hash[s[0]] = 1
	for i := 1; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			hash[s[i]]++
		} else {
			hash[s[i]] = 1
		}
		for k_, v := range hash {
			if v > hash[currentMaxKey] {
				currentMaxKey = k_
			}
		}
		if i-left+1-hash[currentMaxKey] > k {
			// move left
			if s[left] == currentMaxKey {
				for s[left] == currentMaxKey {
					hash[s[left]]--
					left++
				}
				if hash[currentMaxKey] == 0 {
					delete(hash, currentMaxKey)
				}
				continue
			}
			hash[s[left]]--
			if hash[s[left]] == 0 {
				delete(hash, s[left])
			}
			left++
		} else {
			// update max
			// bug : i - left +++++++1111
			if i-left+1 > rsu {
				rsu = i - left + 1
			}
		}
	}
	return rsu
}
