package solution

var rsu string
var canRemoveAll bool

func removeKdigits(num string, k int) string {
	rsu = ""
	canRemoveAll = false
	dfs(num, k)
	if canRemoveAll {
		return "0"
	}
	return rsu
}

func dfs(num string, k int) {
	if canRemoveAll {
		return
	}
	for i := 0; i < len(num); i++ {
		if num[i] != '0' {
			break
		}
		num = num[1:]
	}
	if num == "" {
		canRemoveAll = true
		return
	}
	if k == 0 {
		if rsu == "" || len(rsu) > len(num) || (len(rsu) == len(num) && rsu > num) {
			rsu = num
		}
		return
	}
	for i := 1; i <= k; i++ {
		if num[i] == '0' {
			removeKdigits(num[i+1:], k-i)
			return
		}
	}
	eraseIdx := 0
	for i := 1; i < len(num)-k; i++ {
		if num[i] > num[eraseIdx] {
			eraseIdx = i
		}
	}
	dfs(num[0:eraseIdx]+num[eraseIdx+1:], k-1)
}

// func removeKdigits(num string, k int) string {
// 	countZeroes := 0
// 	for i := 0; i < len(num); i++ {
// 		if num[i] == '0' {
// 			countZeroes++
// 		}
// 	}
// 	if k >= len(num)-countZeroes {
// 		return "0"
// 	}
// 	mark := make([]bool, len(num))

// 	idx := 0
// 	for idx < len(num) {
// 		n := checkExistZero(num, idx, k)
// 		if n > 0 {
// 			for i := idx; i <= n; i++ {
// 				mark[i] = true
// 			}
// 			idx = n
// 			k--
// 			continue
// 		}
// 		currentMaxIdx := idx
// 		for i := idx + 1; i < len(num)-k; i++ {
// 			if num[i] > num[currentMaxIdx] {
// 				currentMaxIdx = i
// 			}
// 		}
// 		idx = currentMaxIdx + 1
// 		mark[currentMaxIdx] = true
// 		k--
// 	}
// 	resu := ""
// 	for i := 0; i < len(mark); i++ {
// 		if !mark[i] {
// 			resu += string(num[i])
// 		}
// 	}
// 	return resu
// }

// func checkExistZero(num string, idx, offset int) int {
// 	resu := -1
// 	for i := idx + 1; i < offset+idx; i++ {
// 		if num[i] == 0 {
// 			resu = i
// 			break
// 		}
// 	}
// 	if resu >= 0 && num[resu+1] == '0' {
// 		for i := resu + 1; i < len(num)-offset; i++ {
// 			if num[i] == '0' {
// 				resu = i
// 			}
// 		}
// 	}
// 	return resu
// }
