package solution

import (
	"sort"
	"strconv"
)

func findKthNumber(n, k int) int {
	if k > n {
		k %= n
	}
	strs := []string{}
	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	var rsu, minIdx int
	for i := 1; i <= k; i++ {
		minIdx = 0
		for j := 1; j < len(strs); j++ {
			if strs[j] < strs[minIdx] {
				minIdx = j
			}
		}
		if i == k {
			rsu, _ = strconv.Atoi(strs[minIdx])
		}
		copy(strs[minIdx:], strs[minIdx+1:])
		strs = strs[:len(strs)-1]
	}
	return rsu
}

func findKthNumber2(n, k int) int {
	if k > n {
		k %= n
	}
	strs := []string{}
	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	for i := 0; i < len(strs); i++ {
		print(strs[i], " ")
	}
	// minIdx := 0
	// var rsu int
	// for i := 1; i <= k; i++ {
	// 	for j := 0; j < len(strs); j++ {
	// 		if strs[i] > strs[minIdx] {
	// 			minIdx = i
	// 		}
	// 	}
	// 	rsu, _ = strconv.Atoi(strs[i])
	// 	copy(strs[minIdx:], strs[minIdx+1:])
	// 	strs = strs[:len(strs)-1]
	// }
	// println(rsu)
	rsu, _ := strconv.Atoi(strs[k-1])
	return rsu
}
