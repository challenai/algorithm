package solution

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	// idea: iterate the nums array, put all the data into a hash.
	// then use two pointer travel the array, find all the two number combination, and use 0 - sum[left, right]
	// but this problem seems have dulplicate element

	// anther idea: sort the array, then fix on element, at last search the pontienal combination of 3 elements
	resu := [][]int{}
	hash := map[string][]int{}
	var left, right, sum int
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	for i := 0; i < len(nums)-2; i++ {
		// !!!
		if nums[i] > 0 {
			break
		}
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				hash[strconv.Itoa(nums[i])+"_"+strconv.Itoa(nums[left])+"_"+strconv.Itoa(nums[right])] = []int{nums[i], nums[left], nums[right]}
				right--
				// for nums[left] == nums[left+1] {
				// 	left++
				// }
				left++
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	for _, val := range hash {
		resu = append(resu, val)
	}

	return resu
}
