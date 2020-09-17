package solution

func myAtoi(s string) int {
	// idea: find first valid char, test if it is a - or +, or just a common number
	// then just parse sequentially
	var startPoint, resu, needMinusSymbol int
	resu = 0
	needMinusSymbol = false
	for i := 0; i < len(s); i++ {
		if s[i] == 45 || s[i] == 43 || (s[i] >= 48 && s[i] <= 57) {
			startPoint = i
			break
		}
	}
	// 45 in byte represent '-'
	if s[startPoint] == 45 {
		needMinusSymbol = true
	}

	if s[startPoint] == 45 || s[startPoint] == 47 {
		startPoint++
	}
	// jump to next number position, try to parse string like this: " - da#^*(jdas  31  " as -31, avoid useless chars
	// deny this idea, because ths situation "  - e2e- 32 3" could appears, we just allow "  -23 3"
	// for i := startPoint; i < len(s); i++ {
	// 	if s[i] >= 48 && s[i] <= 57 {
	// 		startPoint = i
	// 		break
	// 	}
	// }

	for i := startPoint; i < len(s); i++ {

	}

	return 0
}
