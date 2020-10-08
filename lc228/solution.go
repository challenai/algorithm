package solution

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	lastEnd := -1
	resu := []string{strconv.Itoa(nums[0])}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			lastEnd = nums[i]
		} else {
			if lastEnd != -1 {
				resu[len(resu)-1] += "->" + strconv.Itoa(lastEnd)
			}
			lastEnd = -1
			resu = append(resu, strconv.Itoa(nums[i]))
		}
	}
	if lastEnd != -1 {
		resu[len(resu)-1] += "->" + strconv.Itoa(lastEnd)
	}
	return resu
}
