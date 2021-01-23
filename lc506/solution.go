package solution

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	rsu := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	switch len(nums) {
	case 0:
		return []string{}
	case 1:
		return rsu[:1]
	case 2:
		return rsu[:2]
	case 3:
		return rsu[:3]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 3; i < len(nums); i++ {
		rsu = append(rsu, strconv.Itoa(nums[i]))
	}

	return rsu
}
