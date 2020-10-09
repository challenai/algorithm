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
		// lol , no rule ~~
		str1 := numsStr[i] + numsStr[j]
		str2 := numsStr[j] + numsStr[i]
		if str1 > str2 {
			return true
		}
		// consider 9537 vs 953 and 5537 vs 553,,,
		// if len(numsStr[i]) > len(numsStr[j]) {
		// 	if numsStr[i][len(numsStr[j])] < numsStr[i][0] {
		// 		return false
		// 	} else
		// }
		return false
	})
	for i := 0; i < len(numsStr); i++ {
		resu += numsStr[i]
	}
	return resu
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
