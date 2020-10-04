package solution

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	resu := ""
	numsStr := []string{}
	for i := 0; i < len(nums); i++ {
		numsStr = append(numsStr, strconv.Itoa(nums[i]))
	}
	sort.Slice(numsStr, func(i, j int) bool {
		for k := 0; k < min(len(numsStr[i]), len(numsStr[j])); k++ {
			if numsStr[i][k] < numsStr[j][k] {
				return false
			} else if numsStr[i][k] > numsStr[j][k] {
				return true
			}
		}
		// consider 9537 vs 953 and 5537 vs 553,,,
		// if len(numsStr[i]) > len(numsStr[j]) {
		// 	if numsStr[i][len(numsStr[j])] < numsStr[i][0] {
		// 		return false
		// 	} else
		// }
		return true
	})
	for i := 0; i < len(numsStr); i++ {
		resu += numsStr[i]
	}
	println(resu)
	return resu
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
