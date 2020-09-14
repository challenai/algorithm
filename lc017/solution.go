package solution

func letterCombinations(digits string) []string {
	var k int
	charHash := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	current := []string{"", "0"}
	for i := 0; i < len(digits); i++ {
		if chs, ok := charHash[digits[i]]; ok {
			for j := 0; j < len(chs); j++ {
				k = 0
				for current[k] != "0" {
					current = append(current, current[k]+chs[j])
					k++
				}
			}
		}
		for current[0] != "0" {
			current = current[1:]
		}
		current = current[1:]
		current = append(current, "0")
	}
	current = current[:len(current)-1]
	return current
}
