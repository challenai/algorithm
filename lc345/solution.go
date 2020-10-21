package solution

func reverseVowels(s string) string {
	bts := []byte(s)
	hashVowels := map[byte]bool{
		'a': false,
		'e': false,
		'i': false,
		'o': false,
		'u': false,
	}
	var low, high int
	low = 0
	high = len(bts) - 1
	for low < high {
		_, ok := hashVowels[bts[low]]
		for low < high && !ok {
			low++
			_, ok = hashVowels[bts[low]]
		}
		_, ok = hashVowels[bts[high]]
		for low < high && !ok {
			high--
			_, ok = hashVowels[bts[high]]
		}
		if low < high {
			bts[low], bts[high] = bts[high], bts[low]
			low++
			high--
		}
	}
	return string(bts)
}
