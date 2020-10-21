package solution

func reverseString(s []byte) []byte {
	var low, high int
	low = 0
	high = len(s) - 1
	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
	return s
}
