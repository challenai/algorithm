package solution

func convert(s string, n int) string {
	resu := ""

	var current, groupSize int

	groupSize = 2*n - 2
	current = 0
	for current < len(s) {
		resu += string(s[current])
		current += groupSize
	}
	for i := 1; i < n-1; i++ {
		current = i
		for current < len(s) {
			resu += string(s[current])
			if groupSize-i+current-i < len(s) {
				resu += string(s[groupSize-i+current-i])
			}
			current += groupSize
		}
	}
	current = n - 1
	for current < len(s) {
		resu += string(s[current])
		current += groupSize
	}
	return resu
	// for current < len(s) {
	// 	// if current <= len(s) {
	// 	// 	resu += string(s[current])
	// 	// 	if current >= len(s) {
	// 	// 		break
	// 	// 	}
	// 	// }
	// 	if current >= len(s) {
	// 		break
	// 	}
	// 	resu += string(s[current])
	// 	for i := 0; i < n-1; i++ {
	// 		if current+i >= len(s) {
	// 			break
	// 		}
	// 		resu += string(s[current+i])
	// 		if current+2*n-2-i < len(s) {
	// 			resu += string(s[current+2*n-2-i])
	// 		}
	// 	}
	// 	if current+n-1 < len(s) {
	// 		resu += string(s[current+n-1])
	// 	}

	// 	current += 2*n - 2
	// }

	// return resu
}
